package main

import (
	"fmt"
	"sync"
)

func main() {
	var mu sync.Mutex

	mu.Lock()

	go func() {
		fmt.Println(111)
		mu.Unlock()
	}()

	mu.Lock()

}
