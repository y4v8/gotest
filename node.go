package gotest

import "fmt"

type Item interface {
	GetID() int
}

type Node struct {
	Item  *User
	Left  *Node
	Right *Node
	b     int
}

func (n *Node) GetID() int {
	if n == nil || n.Item == nil {
		return 0
	}
	return n.Item.ID
}

func NewNode(item *User, left, right *Node) *Node {
	return &Node{Item: item, Left: left, Right: right}
}

func (n *Node) String() string {
	return fmt.Sprintf("%v-left:%v-right:%v", n.GetID(), n.Left.GetID(), n.Right.GetID())
}

func (n *Node) GetBalanceFactor() int {
	return n.Right.getHeight() - n.Left.getHeight()
}

func (n *Node) getHeight() int {
	height, root := 0, n
	for root != nil {
		if root.b == -1 {
			root = root.Left
		} else {
			root = root.Right
		}
		height++
	}
	return height
}
