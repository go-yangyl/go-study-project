## channel

#### 设计原理

go 语言中最常见的、也是经常被人提及的设计模式就是：
不要通过共享内存的方式进行通信，而是应该通过通信的方式共享内存。
在很多主流的编程语言中，多个线程传递数据的方式一般都是共享内存，
为了解决线程竞争，我们需要限制同一时间能够读写这些变量的线程数量，
然而这与 Go 语言鼓励的设计并不相同。

虽然我们在 Go 语言中也能使用共享内存加互斥锁进行通信，但是 Go 语言提供了一种不同的并发模型，即通信顺序进程（Communicating sequential processes，CSP）1。
Goroutine 和 Channel 分别对应 CSP 中的实体和传递信息的媒介，Goroutine 之间会通过 Channel 传递数据。


#### 数据结构

```go
type hchan struct {
	qcount   uint
	dataqsiz uint
	buf      unsafe.Pointer
	elemsize uint16
	closed   uint32
	elemtype *_type
	sendx    uint
	recvx    uint
	recvq    waitq
	sendq    waitq

	lock mutex
}
```

qcount — Channel 中的元素个数；
dataqsiz — Channel 中的循环队列的长度；
buf — Channel 的缓冲区数据指针；
sendx — Channel 的发送操作处理到的位置；
recvx — Channel 的接收操作处理到的位置；


