package link_list

import (
	"testing"
)

func TestDoubleLinkList(t *testing.T) {
	doubleLinkList := NewDoubleLinkList()

	doubleLinkList.InsertAfterNode(doubleLinkList.DoubleLinkNode, 1)

	doubleLinkList.InsertAfterNode(doubleLinkList.DoubleLinkNode.Next, 2)

	doubleLinkList.InsertAfterNode(doubleLinkList.DoubleLinkNode.Next.Next, 3)

	doubleLinkList.InsertBeforeNode(doubleLinkList.DoubleLinkNode.Next.Next, 4)

	doubleLinkList.Print()
}
