package link_list

import "fmt"

type LinkNode struct {
	Value interface{}
	Next  *LinkNode
}

type NodeList struct {
	Node   *LinkNode
	Length int
}

func NewLinkNode() *LinkNode {
	return &LinkNode{Value: 0, Next: nil}
}

// 获取链表的值
func (l *LinkNode) GetValue() interface{} {
	return l.Value
}

// 获取链表的下一个节点
func (l *LinkNode) GetNode() *LinkNode {
	return l.Next
}

func NewNodeList() *NodeList {
	return &NodeList{Node: NewLinkNode(), Length: 0}
}

// 在某个节点后面插入节点
func (n *NodeList) InsertAfterNode(node *LinkNode, value interface{}) bool {
	if node == nil {
		return false
	}

	newNode := NewLinkNode()
	newNode.Value = value
	newNode.Next = node.Next
	node.Next = newNode
	n.Length++

	return true
}

// 在某个节点前面插入节点
func (n *NodeList) InsertBeforeNode(node *LinkNode, value interface{}) bool {
	if node == nil || node == n.Node {
		return false
	}
	cur := n.Node.Next
	pre := n.Node

	for cur != nil {
		if cur == node {
			break
		}

		pre = cur
		cur = cur.Next
	}

	if cur == nil {
		return false
	}
	// 此时cur节点就是p节点
	newNode := NewLinkNode()
	newNode.Next = cur
	newNode.Value = value
	pre.Next = newNode
	n.Length++
	return true

}

// 在链表的头部插入节点
func (n *NodeList) InsertHeadNode(v interface{}) {
	n.InsertAfterNode(n.Node, v)
}

// 在链表尾部插入节点
func (n *NodeList) InsertTailNode(v interface{}) {
	cur := n.Node

	for cur.Next != nil {
		cur = cur.Next
	}

	n.InsertAfterNode(cur, v)

}

// 删除传入的节点
func (n *NodeList) DelNode(node *LinkNode) bool {
	if node == nil {
		return false
	}

	cur := n.Node.Next
	pre := n.Node

	for nil != cur {
		if cur == node {
			break
		}

		pre = cur
		cur = cur.Next
	}

	if cur == nil {
		return false
	}
	pre.Next = cur.Next
	n.Length--
	return true
}

// 遍历所有的节点
func (n *NodeList) Print() {
	cur := n.Node.Next

	var format = ""
	for cur != nil {
		format += fmt.Sprintf("%+v", cur.GetValue())
		cur = cur.Next
		if nil != cur {
			format += "->"
		}
	}
	fmt.Println(format)
}
