package gotest

import (
	"fmt"
	"io"
	"strconv"
	"strings"
)

type AVLTree struct {
	Root     *Node
	path     []*Node
	GetIndex func(u *User) int
}

func NewAVLTree(getIndex func(*User) int) *AVLTree {
	const max int = 45
	return &AVLTree{
		path:     make([]*Node, 0, max),
		GetIndex: getIndex,
	}
}

func (t *AVLTree) String() string {
	if t.Root == nil {
		return "[]"
	}

	max := 0
	for last := t.Root; last != nil; last = last.Right {
		max = t.GetIndex(last.Item)
	}
	repeat := 0
	for ; max > 0; max /= 10 {
		repeat++
	}
	prefix := "   " + strings.Repeat(" ", repeat)
	format := "%v[%0" + strconv.Itoa(repeat) + "v%v]\n"

	height := t.Root.getHeight()
	rowLength := height*4 - 1
	rowCount := 2<<(height-1) - 1

	var sb strings.Builder
	sb.Grow(rowLength * rowCount)

	fmt.Fprintln(&sb, "")

	var write func(io.Writer, int, *Node)
	write = func(b io.Writer, level int, n *Node) {
		if n == nil {
			return
		}
		if n.Right != nil {
			write(b, level+1, n.Right)
		}
		nb := ".0"
		if n.b > 0 {
			nb = "+" + strconv.FormatInt(int64(n.b), 10)
		} else if n.b < 0 {
			nb = strconv.FormatInt(int64(n.b), 10)
		}
		fmt.Fprintf(b, format, strings.Repeat(prefix, level), t.GetIndex(n.Item), nb)
		if n.Left != nil {
			write(b, level+1, n.Left)
		}
	}
	write(&sb, 1, t.Root)

	return sb.String()
}

func (t *AVLTree) Length() int {
	var sum func(*Node) int
	sum = func(n *Node) int {
		k := 1
		if n.Left != nil {
			k += sum(n.Left)
		}
		if n.Right != nil {
			k += sum(n.Right)
		}
		return k
	}
	return sum(t.Root)
}

func (t *AVLTree) Insert(u *User) {
	node := NewNode(u, nil, nil)
	if t.Root == nil {
		t.Root = node
		return
	}

	last, dir := t.searchNode(u)
	if dir > 0 {
		last.Right = node
	} else if dir < 0 {
		last.Left = node
	} else {
		panic("Index '" + strconv.Itoa(t.GetIndex(u)) + "' is already exist")
	}

	t.rotate()
}

func (t *AVLTree) Delete(u *User) {
	last, dir := t.searchNode(u)
	if last == nil || dir != 0 {
		return
	}

	var parent *Node
	i := len(t.path) - 2
	if i >= 0 {
		parent = t.path[i]
	}

	var p *Node
	lastLeft, lastRight := last.Left, last.Right
	if lastRight == nil {
		p = lastLeft
	} else {
		p = t.searchMin(lastRight)
		m := len(t.path) - 2
		if m >= 0 {
			if t.path[m].Left == p {
				t.path[m].Left = p.Right
			} else {
				t.path[m].Right = p.Right
			}
		}
		p.Left = lastLeft
		p.Right = lastRight
	}

	if parent == nil {
		t.Root = p
	} else if parent.Left == last {
		parent.Left = p
	} else {
		parent.Right = p
	}

	t.rotate()
}

func (t *AVLTree) Get(u *User) *Node {
	id := t.GetIndex(u)

	last := t.Root
	for last != nil {
		lastID := t.GetIndex(last.Item)
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

func (t *AVLTree) searchNode(u *User) (*Node, int) {
	t.path = t.path[0:0]

	if t.Root == nil {
		return nil, 0
	}

	id := t.GetIndex(u)
	last := t.Root
	for {
		t.path = append(t.path, last)

		lastID := t.GetIndex(last.Item)
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
	last := n
	for last.Left != nil {
		last = last.Left
		t.path = append(t.path, last)
	}
	return last
}

func (t *AVLTree) rotateLeft(n *Node) *Node {
	p := n.Right
	n.Right = p.Left
	p.Left = n
	n.b = n.GetBalanceFactor()
	p.b = p.GetBalanceFactor()
	return p
}

func (t *AVLTree) rotateRight(n *Node) *Node {
	p := n.Left
	n.Left = p.Right
	p.Right = n
	n.b = n.GetBalanceFactor()
	p.b = p.GetBalanceFactor()
	return p
}

func (t *AVLTree) rotate() {
	var b int
	var p *Node
	for i := len(t.path) - 1; i >= 0; i-- {
		b = t.path[i].GetBalanceFactor()
		if b > 1 {
			if t.path[i].Right.b < 0 {
				t.path[i].Right = t.rotateRight(t.path[i].Right)
			}
			p = t.rotateLeft(t.path[i])
		} else if b < -1 {
			if t.path[i].Left.b > 0 {
				t.path[i].Left = t.rotateLeft(t.path[i].Left)
			}
			p = t.rotateRight(t.path[i])
		} else {
			t.path[i].b = b
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
