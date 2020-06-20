package gotest

type Node struct {
	Item  *User
	Left  *Node
	Right *Node
	b     int
}

func NewNode(item *User, left, right *Node) *Node {
	return &Node{Item: item, Left: left, Right: right}
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
