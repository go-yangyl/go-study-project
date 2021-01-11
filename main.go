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
