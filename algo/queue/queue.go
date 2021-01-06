package queue

import "fmt"

var _ IQueue = (*Queue)(nil)

type Item interface{}

type Queue struct {
	Items []Item
}

type IQueue interface {
	EnQueue(v interface{}) *Queue
}

// 初始化一个队列
func NewQueue() *Queue {
	return &Queue{
		Items: []Item{},
	}
}

// 进入队列
func (q *Queue) EnQueue(v interface{}) *Queue {
	q.Items = append(q.Items, v)
	return q
}

// 出队列
func (q *Queue) DeQueue() interface{} {
	var item interface{}

	if len(q.Items) > 0 {
		item = q.Items[0]
		q.Items = q.Items[1:]
	}
	return item
}

// 遍历队列
func (q *Queue) Print() {
	var format string

	for _, v := range q.Items {
		format += fmt.Sprintf("%v->", v)
	}
	fmt.Println(format)
}
