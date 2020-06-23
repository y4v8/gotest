package gotest

import "strconv"

type Item User

type Node struct {
	Item   *Item
	Left   *Node
	Right  *Node
	height int
}

func NewNode(item *Item) *Node {
	return &Node{Item: item, height: 1}
}

func (n *Node) GetBalanceFactor() int {
	return n.Right.getHeight() - n.Left.getHeight()
}

func (n *Node) String(getIndex func(*Item) int) string {
	nb := ".0"
	b := n.GetBalanceFactor()
	if b > 0 {
		nb = "+" + strconv.Itoa(b)
	} else if b < 0 {
		nb = strconv.Itoa(b)
	}
	return strconv.Itoa(getIndex(n.Item)) + nb
}

func (n *Node) getHeight() int {
	if n == nil {
		return 0
	}
	return n.height
}

func (n *Node) FixHeight() {
	hl := n.Left.getHeight()
	hr := n.Right.getHeight()
	if hl > hr {
		n.height = hl + 1
	} else {
		n.height = hr + 1
	}
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
