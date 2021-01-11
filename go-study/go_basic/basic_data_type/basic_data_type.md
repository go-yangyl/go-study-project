### 值类型



```text
int 
int8               包含一个基为 10 的数字，用 1 个存储字节表示。uint8 的值从 0 到 255 
int32(rune)        包含一个基为 10 的数字，用 4 个存储字节表示。int32 的值从 -2147483648 到 2147483647
int64              包含一个基为 10 的数字，用 8 个存储字节表示
uint32 
uint8(byte) 
uint32 
uint64

bool

String

array

float32  float64
```



### 引用类型



```
slice 切片
map 
chan
```



### 内置函数



```
append   -- 追加元素到数组 slice当中，返回修改后的数组或者切片，不会改变原来的数组，会开辟一块新的内存空间，只有修改才是引用传递
close    -- 关闭channel 
delete   -- 用于删除map中的key对应的value
make     -- 用于内存分配，返回Type类型本身 可用于 slice map channel
new      -- 用于内存分配，返回Type类型的指针，可用于 struct
len      -- 返回slice array map string channel的长度
cap      -- 返回最大容量 可用于 slice channel
panic    -- 停止常规的goroutine 
recover  -- 允许程序定义goroutine的panic动作
```



### 位运算

```
>>  除以2
<<  乘以2
&   两个都是都为1才是1
|   两个里面只要有1个为1就是1
^   当两对应的二进位相异时，结果为1
```

