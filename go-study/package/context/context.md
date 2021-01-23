## context

上下文 context.Context Go 语言中用来设置截止日期、同步信号，传递请求相关值的结构体。上下文与 Goroutine 有比较密切的关系，是 Go 语言中独特的设计，在其他编程语言中我们很少见到类似的概念。

context.Context 是 Go 语言在 1.7 版本中引入标准库的接口1，该接口定义了四个需要实现的方法，其中包括：

Deadline — 返回 context.Context 被取消的时间，也就是完成工作的截止日期；
Done — 返回一个 Channel，这个 Channel 会在当前工作完成或者上下文被取消后关闭，多次调用 Done 方法会返回同一个 Channel；
Err — 返回 context.Context 结束的原因，它只会在 Done 方法对应的 Channel 关闭时返回非空的值；
如果 context.Context 被取消，会返回 Canceled 错误；
如果 context.Context 超时，会返回 DeadlineExceeded 错误；
Value — 从 context.Context 中获取键对应的值，对于同一个上下文来说，多次调用 Value 并传入相同的 Key 会返回相同的结果，该方法可以用来传递请求特定的数据；

```go
type Context interface {
	Deadline() (deadline time.Time, ok bool)
	Done() <-chan struct{}
	Err() error
	Value(key interface{}) interface{}
}
```

```go
func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	go handle(ctx, 500*time.Millisecond)
	select {
	case <-ctx.Done():
		fmt.Println("main", ctx.Err())
	}
}

func handle(ctx context.Context, duration time.Duration) {
	select {
	case <-ctx.Done():
		fmt.Println("handle", ctx.Err())
	case <-time.After(duration):
		fmt.Println("process request with", duration)
	}
}
```
因为过期时间大于处理时间，所以我们有足够的时间处理该请求，运行上述代码会打印出下面的内容：
```shell script
$ go run context.go
process request with 500ms
main context deadline exceeded

```

handle 函数没有进入超时的 select 分支，但是 main 函数的 select 却会等待 context.Context 超时并打印出 main context deadline exceeded。

如果我们将处理请求时间增加至 1500ms，整个程序都会因为上下文的过期而被中止，

```shell script
$ go run context.go
main context deadline exceeded
handle context deadline exceeded
```
