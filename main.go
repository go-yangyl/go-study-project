package main

import "fmt"

func main() {

	fmt.Printf()
}

type W struct {
}

func (w *W) Write(p []byte) (n int, err error) {
	return 1, nil
}
