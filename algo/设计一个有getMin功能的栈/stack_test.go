package stask

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
	stack := NewStack()

	stack.Push(2)
	stack.Push(3)
	stack.Push(1)
	fmt.Println(stack.GetMIn())
	stack.Pop()
	fmt.Println(stack.GetMIn())

}
