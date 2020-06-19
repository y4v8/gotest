package gotest

import (
	"fmt"
	"io"
	"log"
	"strconv"
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
	const PREFIX string = "     "

	if t.Root == nil {
		return "[    ]"
	}

	height := t.Root.getHeight()
	rowLength := height*4 - 1
	rowCount := 2 ^ height - 1

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
		fmt.Fprintf(b, "%v[%02v%v]\n", strings.Repeat(PREFIX, level), n.GetID(), nb)
		if n.Left != nil {
			write(b, level+1, n.Left)
		}
	}
	write(&sb, 1, t.Root)

	return sb.String()
}

func (t *TopTree) searchNode(id int) (*Node, int) {
	t.path = t.path[0:0]

	if t.Root == nil {
		return nil, 0
	}

	last := t.Root
	for {
		t.path = append(t.path, last)

		if id < last.Item.ID {
			if last.Left == nil {
				return last, -1
			}
			last = last.Left
		} else if id > last.Item.ID {
			if last.Right == nil {
				return last, 1
			}
			last = last.Right
		} else {
			return last, 0
		}
	}
}

type Comparer interface {
	Compare(Comparer) int
}
type Ucmp struct {
	ID int
}

func (u *Ucmp) Compare(u2 Comparer) int {
	return u.ID - u2.(*Ucmp).ID
}

func (t *TopTree) Insert(u *User) {

	node := NewNode(u, nil, nil)
	if t.Root == nil {
		t.Root = node
		return
	}

	last, dir := t.searchNode(u.ID)
	if dir > 0 {
		last.Right = node
	} else if dir < 0 {
		last.Left = node
	} else {
		log.Fatalf("ID '%v' is already exist", u.ID)
	}

	t.rotate()
}

func (t *TopTree) Append2(u *User) {
	t.Last = NewNode(u, nil, nil)

	if t.Root == nil {
		t.Root = t.Last
		return
	}

	lastNode, dir := t.searchNode(t.Last.Item.ID)
	if lastNode != t.Last {
		log.Fatal("The found node ", lastNode.Item.ID, " is not equal to the last node ", u.ID)
	}
	if dir < 0 {
		log.Fatal("A new node must be insert to the right:", u.ID, dir)
	}
	lastNode.Right = t.Last

	t.rotate()
}

func (t *TopTree) rotate() {
	for i := len(t.path) - 1; i >= 0; i-- {
		b := t.path[i].GetBalanceFactor()
		if b > 1 {
			if t.path[i].Right.b >= 0 {
				if i == 0 {
					t.Root = t.path[i].Right
				} else if t.path[i-1].Left == t.path[i] {
					t.path[i-1].Left = t.path[i].Right
				} else {
					t.path[i-1].Right = t.path[i].Right
				}
				t.path[i].Right, t.path[i].Right.Left = t.path[i].Right.Left, t.path[i]
			} else {
				if i == 0 {
					t.Root = t.path[i].Right.Left
				} else if t.path[i-1].Left == t.path[i] {
					t.path[i-1].Left = t.path[i].Right.Left
				} else {
					t.path[i-1].Right = t.path[i].Right.Left
				}
				t.path[i].Right, t.path[i].Right.Left, t.path[i].Right.Left.Left, t.path[i].Right.Left.Right =
					t.path[i].Right.Left.Left, t.path[i].Right.Left.Right, t.path[i], t.path[i].Right
			}
		} else if b < -1 {
			if t.path[i].Left.b >= 0 {
				if i == 0 {
					t.Root = t.path[i].Left
				} else if t.path[i-1].Left == t.path[i] {
					t.path[i-1].Left = t.path[i].Left
				} else {
					t.path[i-1].Right = t.path[i].Left
				}
				t.path[i].Left, t.path[i].Left.Right = t.path[i].Left.Right, t.path[i]
			} else {
				if i == 0 {
					t.Root = t.path[i].Left.Right
				} else if t.path[i-1].Left == t.path[i] {
					t.path[i-1].Left = t.path[i].Left.Right
				} else {
					t.path[i-1].Right = t.path[i].Left.Right
				}
				t.path[i].Left.Right, t.path[i].Left, t.path[i].Left.Right.Left, t.path[i].Left.Right.Right =
					t.path[i].Left.Right.Left, t.path[i].Left.Right.Right, t.path[i].Left, t.path[i]
			}
		}

		t.path[i].b = t.path[i].GetBalanceFactor()
	}
}

func (t *TopTree) Delete(u *User) {
	last, dir := t.searchNode(u.ID)
	if last == nil || dir != 0 {
		return
	}
	if last == t.Root {
		t.Root = nil
		return
	}

	parent := t.path[len(t.path)-2]
	if last.Left == nil && last.Right == nil {
		if parent.Left == last {
			parent.Left = nil
		} else {
			parent.Right = nil
		}
		return
	}

	leftID, rightID := last.Item.ID, last.Item.ID
	if last.Left != nil {
		leftID = last.Item.ID - last.Left.Item.ID
	}
	if last.Right != nil {
		rightID = last.Right.Item.ID - last.Item.ID
	}
	if leftID < rightID {
		if dir == -1 {
			parent.Left, last.Left = last.Left, parent.Left
		} else {
			parent.Right, last.Left = last.Left, parent.Right
		}
	} else {
		if dir == -1 {
			parent.Left, last.Right = last.Right, parent.Left
		} else {
			parent.Right, last.Right = last.Right, parent.Right
		}
	}

	t.rotate()
}
