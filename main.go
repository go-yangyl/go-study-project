package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	var a int32
	fmt.Println(atomic.AddInt32(&a, 1))
}
