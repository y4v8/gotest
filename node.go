package gotest

import "strconv"

type Item User

type Node struct {
	Item   *Item
	Left   *Node
	Right  *Node
	height int8
}

func NewNode(item *Item) *Node {
	return &Node{Item: item, height: 1}
}

func (n *Node) GetBalance() int8 {
	return n.Right.getHeight() - n.Left.getHeight()
}

func (n *Node) String(getIndex func(*Item) int) string {
	nb := ".0"
	b := n.GetBalance()
	if b > 0 {
		nb = "+" + strconv.Itoa(int(b))
	} else if b < 0 {
		nb = strconv.Itoa(int(b))
	}
	return strconv.Itoa(getIndex(n.Item)) + nb
}

func (n *Node) getHeight() int8 {
	if n == nil {
		return 0
	}
	return n.height
}

func (n *Node) FixHeight() bool {
	hn := n.height
	hl := n.Left.getHeight()
	hr := n.Right.getHeight()
	if hl > hr {
		n.height = hl + 1
	} else {
		n.height = hr + 1
	}
	return hn != n.height
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
