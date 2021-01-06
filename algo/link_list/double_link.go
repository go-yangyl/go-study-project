package link_list

import "fmt"

type DoubleLinkNode struct {
	Value interface{}
	Prev  *DoubleLinkNode
	Next  *DoubleLinkNode
}

type DoubleLinkList struct {
	DoubleLinkNode *DoubleLinkNode
	Length         int
}

func NewDoubleLinkNode() *DoubleLinkNode { return &DoubleLinkNode{Value: 0, Prev: nil, Next: nil} }

func (d *DoubleLinkNode) GetValue() interface{} { return d.Value }

func (d *DoubleLinkNode) GetPrevNode() *DoubleLinkNode { return d.Prev }

func (d *DoubleLinkNode) GetNextNode() *DoubleLinkNode { return d.Next }

func NewDoubleLinkList() *DoubleLinkList {
	return &DoubleLinkList{DoubleLinkNode: NewDoubleLinkNode(), Length: 0}
}

// 双链表插入到后面
func (d *DoubleLinkList) InsertAfterNode(node *DoubleLinkNode, value interface{}) bool {
	if node == nil {
		return false
	}

	newNode := NewDoubleLinkNode()
	newNode.Value = value
	newNode.Prev = node
	node.Next = newNode
	return true
}

// 双链表插入到前面
func (d *DoubleLinkList) InsertBeforeNode(node *DoubleLinkNode, value interface{}) bool {
	if node == nil {
		return false
	}

	newNode := NewDoubleLinkNode()
	newNode.Value = value
	newNode.Prev = node.Prev
	newNode.Next = node
	node.Prev.Next = newNode
	return true
}

func (d *DoubleLinkList) Print() {
	if d.DoubleLinkNode.Next == nil {
		return
	}

	cur := d.DoubleLinkNode.Next
	format := ""
	for nil != cur {
		format += fmt.Sprintf("%+v->", cur.Value)
		cur = cur.Next
	}

	fmt.Println(format)
}
