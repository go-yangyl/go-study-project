## sync



#### waitgroup

问题：如何确保所有的子协程在主goroutine之前都要执行

1.这种方式可以实现
```go
var num int
	for i := 0; i < 10; i++ {
		go func() {
			num++
			fmt.Println(num)

		}()
	}

	for num != 10 {

	}

    fmt.Println("确保所有的goroutine都执行啦")
```

