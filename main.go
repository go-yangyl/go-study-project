package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	var num int32
	atomic.AddInt32(&num, 1)

	atomic.StoreInt32(&num, 100)
	fmt.Println(atomic.LoadInt32(&num))
}
