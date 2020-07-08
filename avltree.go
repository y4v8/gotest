package gotest

type AVLTree struct {
	Root     *Node
	path     []*Node
	getIndex func(*Item) int
	tmpNode  *Node
}

func NewAVLTree(getIndex func(*Item) int) *AVLTree {
	const maxHeight int = 45
	return &AVLTree{
		path:     make([]*Node, 0, maxHeight),
		getIndex: getIndex,
	}
}

func (t *AVLTree) Insert(item *Item) {
	node := t.tmpNode
	if node == nil {
		node = NewNode(item)
	} else {
		node.Item = item
		t.tmpNode = nil
	}

	if t.Root == nil {
		t.Root = node
		return
	}

	last, dir := t.searchNode(item)
	if dir > 0 {
		last.Right = node
	} else if dir < 0 {
		last.Left = node
	} else {
		panic("index is already exist")
	}

	t.rotate()
}

func (t *AVLTree) Delete(item *Item) {
	last, dir := t.searchNode(item)
	if last == nil || dir != 0 {
		return
	}

	var parent *Node
	i := len(t.path) - 2
	if i >= 0 {
		parent = t.path[i]
	}

	var p *Node
	if last.Right == nil {
		p = last.Left
	} else {
		p = t.searchMin(last.Right)
		p.Left = last.Left
		if p != last.Right {
			i = len(t.path) - 1
			t.path[i].Left = p.Right
			p.Right = last.Right
		}
	}

	if parent == nil {
		t.Root = p
	} else if parent.Left == last {
		parent.Left = p
	} else {
		parent.Right = p
	}

	last.Item, last.Left, last.Right, last.height = nil, nil, nil, 1
	t.tmpNode = last

	t.rotate()
}

func (t *AVLTree) Get(item *Item) *Node {
	id := t.getIndex(item)

	last := t.Root
	for last != nil {
		lastID := t.getIndex(last.Item)
		if id < lastID {
			last = last.Left
		} else if id > lastID {
			last = last.Right
		} else {
			return last
		}
	}
	return nil
}

func (t *AVLTree) GetItems(item *Item) []*Item {
	nodes := make([]*Item, 0, t.Root.Length())
	if t.Root == nil {
		return nodes
	}

	id := t.getIndex(item)
	var appendItem func(*Node)
	appendItem = func(n *Node) {
		if n.Left != nil && t.getIndex(n.Item) > id {
			appendItem(n.Left)
		}
		if t.getIndex(n.Item) >= id {
			nodes = append(nodes, n.Item)
		}
		if n.Right != nil {
			appendItem(n.Right)
		}
	}
	appendItem(t.Root)

	return nodes
}

func (t *AVLTree) searchNode(item *Item) (*Node, int) {
	t.path = t.path[0:0]

	if t.Root == nil {
		return nil, 0
	}

	id := t.getIndex(item)
	last := t.Root
	for {
		t.path = append(t.path, last)

		lastID := t.getIndex(last.Item)
		if id < lastID {
			if last.Left == nil {
				return last, -1
			}
			last = last.Left
		} else if id > lastID {
			if last.Right == nil {
				return last, 1
			}
			last = last.Right
		} else {
			return last, 0
		}
	}
}

func (t *AVLTree) searchMin(n *Node) *Node {
	i := len(t.path) - 1
	last := n
	for last.Left != nil {
		t.path = append(t.path, last)
		last = last.Left
	}
	t.path[i] = last
	return last
}

func (t *AVLTree) rotateLeft(n *Node) *Node {
	p := n.Right
	p.Left, n.Right = n, p.Left
	n.FixHeight()
	p.FixHeight()
	return p
}

func (t *AVLTree) rotateRight(n *Node) *Node {
	p := n.Left
	p.Right, n.Left = n, p.Right
	n.FixHeight()
	p.FixHeight()
	return p
}

func (t *AVLTree) rotate() {
	var b int8
	var p *Node
	for i := len(t.path) - 1; i >= 0; i-- {
		b = t.path[i].GetBalance()
		if b > 1 {
			if t.path[i].Right.GetBalance() < 0 {
				t.path[i].Right = t.rotateRight(t.path[i].Right)
			}
			p = t.rotateLeft(t.path[i])
		} else if b < -1 {
			if t.path[i].Left.GetBalance() > 0 {
				t.path[i].Left = t.rotateLeft(t.path[i].Left)
			}
			p = t.rotateRight(t.path[i])
		} else {
			t.path[i].FixHeight()
			continue
		}

		if i == 0 {
			t.Root = p
		} else if t.path[i-1].Left == t.path[i] {
			t.path[i-1].Left = p
		} else {
			t.path[i-1].Right = p
		}
	}
}
