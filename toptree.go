package gotest

import (
	"fmt"
	"io"
	"log"
	"strings"
)

type TopTree struct {
	Root *Node
	Last *Node
	path []*Node
}

func NewTopTree() *TopTree {
	const max int = 45
	return &TopTree{
		path: make([]*Node, 0, max),
	}
}

func (t *TopTree) String() string {
	var sb strings.Builder
	sb.Grow(1024 * 1024)

	fmt.Fprintln(&sb, "")

	var write func(io.Writer, int, *Node)
	write = func(prefix string) func(b io.Writer, level int, n *Node) {
		return func(b io.Writer, level int, n *Node) {
			if n == nil {
				return
			}
			if n.Right != nil {
				write(b, level+1, n.Right)
			}
			fmt.Fprintf(b, "%v[%02v]\n", strings.Repeat(prefix, level), n.GetID())
			if n.Left != nil {
				write(b, level+1, n.Left)
			}
		}
	}("   ")
	write(&sb, 1, t.Root)

	return sb.String()
}

func (t *TopTree) searchPath(id int) (*Node, int) {
	t.path = t.path[:]

	if t.Root == nil {
		return nil, 0
	}

	root := t.Root
	for {
		t.path = append(t.path, root)
		fmt.Println("append.path:", root.Item.ID, "search:", id)

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
	if t.Last == nil {
		fmt.Println("append,nil,", u.ID)
		t.Last = NewNode(u, nil, nil)
		t.Root = t.Last
		return
	}
	fmt.Println("append", u.ID)

	lastNode, dir := t.searchPath(t.Last.Item.ID)
	if lastNode != t.Last {
		log.Fatal("The found node ", lastNode.Item.ID, " is not equal to the last node ", u.ID)
	}
	if dir < 0 {
		log.Fatal("A new node must be insert to the right:", u.ID, dir)
	}
	t.Last.Right = NewNode(u, nil, nil)
	t.Last = t.Last.Right

	t.rotate()
}

func (t *TopTree) rotate() {
	for i := len(t.path) - 2; i > 0; i-- {
		b := t.path[i].GetBalanceFactor()
		if b > 1 {
			if t.path[i].Right.b >= 0 {
				if t.path[i-1].Left == t.path[i] {
					t.path[i-1].Left = t.path[i].Right
				} else {
					t.path[i-1].Right = t.path[i].Right
				}
				t.path[i].Right, t.path[i].Right.Left = t.path[i].Right.Left, t.path[i]
			} else {
				if t.path[i-1].Left == t.path[i] {
					t.path[i-1].Left = t.path[i].Right.Left
				} else {
					t.path[i-1].Right = t.path[i].Right.Left
				}
				t.path[i].Right, t.path[i].Right.Left, t.path[i].Right.Left.Left, t.path[i].Right.Left.Right =
					t.path[i].Right.Left.Left, t.path[i].Right.Left.Right, t.path[i], t.path[i].Right
			}
		} else if b < -1 {
			if t.path[i].Left.b >= 0 {
				if t.path[i-1].Left == t.path[i] {
					t.path[i-1].Left = t.path[i].Left
				} else {
					t.path[i-1].Right = t.path[i].Left
				}
				t.path[i].Left, t.path[i].Left.Right = t.path[i].Left.Right, t.path[i]
			} else {
				if t.path[i-1].Left == t.path[i] {
					t.path[i-1].Left = t.path[i].Left.Right
				} else {
					t.path[i-1].Right = t.path[i].Left.Right
				}
				t.path[i].Left.Right, t.path[i].Left, t.path[i].Left.Right.Left, t.path[i].Left.Right.Right =
					t.path[i].Left.Right.Left, t.path[i].Left.Right.Right, t.path[i].Left, t.path[i]
			}
		}

		b2 := t.path[i].GetBalanceFactor()
		if b == b2 {
			return
		}
		t.path[i].b = b2
	}
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

	t.rotate()
}
