## select

select 是操作系统中的系统调用，我们经常会使用 select、poll 和 epoll 等函数构建 I/O 多路复用模型提升程序的性能。Go 语言的 select 与操作系统中的 select 比较相似，本节会介绍 Go 语言 select 关键字常见的现象、数据结构以及实现原理。

C 语言的 select 系统调用可以同时监听多个文件描述符的可读或者可写的状态，Go 语言中的 select 也能够让 Goroutine 同时等待多个 Channel 可读或者可写，在多个文件或者 Channel状态改变之前，select 会一直阻塞当前线程或者 Goroutine。

当我们在 Go 语言中使用 select 控制结构时，会遇到两个有趣的现象：

select 能在 Channel 上进行非阻塞的收发操作；
select 在遇到多个 Channel 同时响应时，会随机执行一种情况；
这两个现象是学习 select 时经常会遇到的，我们来深入了解具体场景并分析这两个现象背后的设计原理。

非阻塞的收发
在通常情况下，select 语句会阻塞当前 Goroutine 并等待多个 Channel 中的一个达到可以收发的状态。但是如果 select 控制结构中包含 default 语句，那么这个 select 语句在执行时会遇到以下两种情况：

当存在可以收发的 Channel 时，直接处理该 Channel 对应的 case；
当不存在可以收发的 Channel 时，执行 default 中的语句；
当我们运行下面的代码时就不会阻塞当前的 Goroutine，它会直接执行 default 中的代码。