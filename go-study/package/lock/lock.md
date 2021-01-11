## 并发安全和锁

多个请求或者goroutine操作同一块资源的时候，会发送数据的不一致

### 互斥锁与读写锁

1. **原子性**：把一个互斥量锁定为一个原子操作，这意味着操作系统（或pthread函数库）保证了如果一个线程锁定了一个互斥量，没有其他线程在同一时间可以成功锁定这个互斥量；
2. **唯一性**：如果一个线程锁定了一个互斥量，在它解除锁定之前，没有其他线程可以锁定这个互斥量
3. **非繁忙等待** : 如果一个线程已经锁定了一个互斥量，第二个线程又试图去锁定这个互斥量，则第二个线程将被挂起（不占用任何cpu资源），直到第一个线程解除对这个互斥量的锁定为止，第二个线程则被唤醒并继续执行，同时锁定这个互斥量。



```go
  var mu sync.Mutex
	mu.Lock()
	defer mu.Unlock()

  var rw sync.Mutex
  rw.Lock()
	defer rw.Unlock()
```

----

#### 死锁代码
```go
func TradeUnionPointGame_GetTaskDetail(unionId, taskId int) (bool, *TradeUnionTask) {
	AllTradeUnionInfoList.syncTradeInfoLock.RLock()
	defer AllTradeUnionInfoList.syncTradeInfoLock.RUnlock()
    // 如果不使用defer，加锁后return会直接返回，并没有释放锁，下次请求再次尝试加锁
	for _, v := range AllTradeUnionInfoList.TradeUnionInfo {
		for _, val := range v.TradeUnionTask {
			if v.UnionId == unionId && val.TaskId == taskId {
				return true, val
			}
		}
	}


	// 数据同步到内存
	err, tradeUnionTask := TradeUnionPointGame_GetTradeUnionInfo(unionId)
	if err != nil {
		return false, nil
	}
	for _, v := range tradeUnionTask.TradeUnionTask {
		if v.TaskId == taskId {
			return true, v
		}
	}

	return false, nil
}
```

```go
package main

import "sync"

var mu sync.Mutex

func main() {

	mu.Lock()
	A()
	mu.Unlock()
}

func A() {
	mu.Lock()

	mu.Unlock()
}

```

#### 源码解读

```go
type Mutex struct {
	state int32    // 信号状态
	sema  uint32   // 信号量
}

const (
	mutexLocked = 1 << iota // mutex is locked
	mutexWoken
	mutexStarving
	mutexWaiterShift = iota
  ）
```

其中state是记录用来记录加锁状态的，将一个整型按位划分来表示不同的含义，从低到高分别为第1位到第32位，

- 第1位表示是否被锁住，即0表示没有锁住，mutexLocked也就是1表示已经被锁住。
- 第2表示是否被唤醒，1表示被唤醒，mutexWoken=2表示被唤醒。
- 第3位到第32位表示等待在mutex上协程数量，mutexWaiterShift=3表示在获取等待协程数量，需要将state右移位3位。



信号量是进程间通信处理同步互斥的机制，通过一个计数器来控制对共享资源的访问次数限制。例如一个办公室有两台打印机，有几十台电脑连上，这是同时只能允许两个电脑进行打印，而其他电脑必须排队等待完成后才能打印。

sema就是信号量，是一个非负数的全局变量，该变量有两个操作P和V，PV操作都是不可中断的。



**P(S):**
（1）执行S=S-1；
（2）进行以下判断：

- 如果S < 0，进入阻塞队列，直到满足S>=0
- 如果S >= 0, 直接返回
  因此P操作执行一次意味着分配一个资源，如上打印机意味着是资源，当S小于0意味着没有可用资源了，只能一直等待，直到资源空闲出来时才能继续。



**V(S):**
（1）执行S=S+1；
（2）进行以下判断：

- 如果S > 0，直接返回
- 如果S <= 0， 释放阻塞队列中的第一个等待进程
  因此V操作执行一次意味着释放一个资源，当S小于等于0时，意味着还有进程在请求资源，此时释放了一个资源，就需要从等待队列中拿出一个进程来使用此刻释放的资源。

#### golang中信号量操作

**runtime_Semacquire**
func runtime_Semacquire(s *uint32)，P操作，等待*s大于等于0，源码在runtime/sema.go中

**runtime_Semrelease**
func runtime_Semrelease, V操作，阻塞等待被唤醒，目前版本在runtime/sema.go中(定义稍有不同了)。



简单的实现方式

```go
type Mutex struct {
    sema uint32
}

func NewMutex() *Mutex {
    var mu Mutex
    mu.sema = 1
    return &mu
}

func (m *Mutex) Lock() {
        runtime_Semacquire(&m.sema)
}

func (m *Mutex2) Unlock() {
    runtime_Semrelease(&m.sema)
}
```



当然，这个实现有点不符合要求。如果有个家伙不那么靠谱，加锁了一次，但是解锁了两次。第二次解锁的时候，应该报出一个错误，而不是让错误隐藏。于是乎，我们想到用一个变量表示加锁的次数。这样就可以判断有没有多次解锁。



改进代码如下：

```go
type Mutex struct {
        key  int32
        sema uint32
}

func (m *Mutex) Lock() {
        if atomic.AddInt32(&m.key, 1) == 1 {
                // changed from 0 to 1; we hold lock
                return
        }
        runtime_Semacquire(&m.sema)
}

func (m *Mutex) Unlock() {
        switch v := atomic.AddInt32(&m.key, -1); {
        case v == 0:
                // changed from 1 to 0; no contention
                return
        case v == -1:
                // changed from 0 to -1: wasn't locked
                // (or there are 4 billion goroutines waiting)
                panic("sync: unlock of unlocked mutex")
        }
        runtime_Semrelease(&m.sema)
}
```

这个解决方案除了解决了我们前面说的重复加锁的问题外，还对我们初始化工作做了简化，不需要构造函数了。执行过程中值变化如下：

- 初始：key=0, sema = 0
- Lock第一次：key+1=1返回，sema=0，即第一次不进行P操作，直接将key加1表示获取了锁。
- Lock第二次：key=2，进行P操作，发现sema-1 =-1<0，阻塞等待获取锁。

当执行了一次Lock后，key=1，sema=0，执行以下操作时：

- Unlock第一次：key-1=0返回，sema=0，第一次解锁不执行V操作，直接key减1表示释放锁。
- Unlock第二次：key-1=-1，表示解锁过了，返回异常。

当执行了两次Lock后，key=2，sema=-1，执行以下操作时：

- Unlock第一次：key-1=1，执行V操作runtime_Semrelease，发现sema+1=0，会阻塞直到唤醒了其他协程，然后返回。

简单来说，增加一个key变量后，sema=0表示有一个资源，跟只用信号量时sema=1含义一样，在golang mutex也是基于此实现的。

