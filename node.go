package gotest

import "strconv"

type Item User

type Node struct {
	Item  *Item
	Left  *Node
	Right *Node
	b     int
}

func NewNode(item *Item) *Node {
	return &Node{Item: item}
}

func (n *Node) GetBalanceFactor() int {
	return n.Right.getHeight() - n.Left.getHeight()
}

func (n *Node) String(getIndex func(*Item) int) string {
	nb := ".0"
	if n.b > 0 {
		nb = "+" + strconv.Itoa(n.b)
	} else if n.b < 0 {
		nb = strconv.Itoa(n.b)
	}
	return strconv.Itoa(getIndex(n.Item)) + nb
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

func (n *Node) Length() int {
	if n == nil {
		return 0
	}

	k := 1
	if n.Left != nil {
		k += n.Left.Length()
	}
	if n.Right != nil {
		k += n.Right.Length()
	}
	return k
}
