package link_list

import (
	"fmt"
	"testing"
)

func TestList(t *testing.T) {
	nodeList := NewNodeList()
	nodeList.InsertAfterNode(nodeList.Node, 1)
	nodeList.InsertAfterNode(nodeList.Node.Next, 2)
	nodeList.InsertAfterNode(nodeList.Node.Next.Next, 3)
	nodeList.InsertBeforeNode(nodeList.Node.Next.Next, 4)
	nodeList.InsertHeadNode(5)
	nodeList.InsertTailNode(6)
	nodeList.DelNode(nodeList.Node.Next.Next)
	nodeList.Print()
	fmt.Println(nodeList.Length)

}
