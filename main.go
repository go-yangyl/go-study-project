package main

import (
	"fmt"
)

type Duck interface {
	Quack()
}
type Cat struct{}

func (c Cat) Quack() {
	fmt.Println("meow")
}

func main() {
	var a = []int{1, 2, 3}
	b := arr(a)
	b[0] = 10
	fmt.Println(a, b)

}

func arr(a []int) []int {
	b := make([]int, len(a))
	copy(b, a)
	return b
}
