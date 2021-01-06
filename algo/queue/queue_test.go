package queue

import (
	"testing"
)

func TestQueue(t *testing.T) {
	queue := NewQueue()
	queue.EnQueue(1)
	queue.EnQueue(2)
	queue.EnQueue(3)
	queue.EnQueue(4)
	queue.Print()
	queue.DeQueue()
	queue.DeQueue()
	queue.DeQueue()
	queue.Print()
}
