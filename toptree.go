package gotest

type TopTree struct {
	Root *Node
	Last *Node
	path []*Node
}

func NewTopTree() *TopTree {
	return &TopTree{path: make([]*Node, 0, 45)}
}

func (t *TopTree) searchPath(id int) (*Node, int) {
	t.path = t.path[:]

	if t.Root == nil {
		return nil, 0
	}

	root := t.Root
	for {
		t.path = append(t.path, root)

		if id < root.Item.ID {
			if root.Left == nil {
				return root, -1
			}
			root = root.Left
		} else if id > root.Item.ID {
			if root.Right == nil {
				return root, 1
			}
			root = root.Right
		} else {
			return root, 0
		}
	}
}

func (t *TopTree) Append(u *User) {
	lastNode, dir := t.searchPath(u.ID)
	if lastNode == nil {
		t.Root = NewNode(u, nil, nil)
	}
	if dir == -1 {
		lastNode.Left = NewNode(u, nil, nil)
	} else if dir == 1 {
		lastNode.Right = NewNode(u, nil, nil)
	} else {
		lastNode.Item = u
	}
	// rotate
}

func (t *TopTree) Delete(n *Node) {
	lastNode, dir := t.searchPath(n.Item.ID)
	if lastNode == nil || dir != 0 {
		return
	}
	if lastNode == t.Root {
		t.Root = nil
		return
	}

	parent := t.path[len(t.path)-2]
	if n.Left == nil && n.Right == nil {
		if dir == -1 {
			parent.Left = nil
		} else {
			parent.Right = nil
		}
		return
	}

	leftID, rightID := n.Item.ID, n.Item.ID
	if n.Left != nil {
		leftID = n.Item.ID - n.Left.Item.ID
	}
	if n.Right != nil {
		rightID = n.Right.Item.ID - n.Item.ID
	}
	if leftID < rightID {
		if dir == -1 {
			parent.Left, n.Left = n.Left, parent.Left
		} else {
			parent.Right, n.Left = n.Left, parent.Right
		}
	} else {
		if dir == -1 {
			parent.Left, n.Right = n.Right, parent.Left
		} else {
			parent.Right, n.Right = n.Right, parent.Right
		}
	}

	// rotate
}

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
	return 0
}

func rotate(left, right *Node) {

}

func rotateLeft(left, right *Node) {

}

func rotateRight(left, right *Node) {

}
