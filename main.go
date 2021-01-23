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

}
