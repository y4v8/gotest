package gotest

import (
	"fmt"
	"io"
	"strconv"
	"strings"
)

type Node struct {
	Item  *User
	Left  *Node
	Right *Node
	b     int
}

func NewNode(item *User, left, right *Node) *Node {
	return &Node{Item: item, Left: left, Right: right}
}

func (n *Node) Sprint(getIndex func(*User) int) string {
	if n == nil {
		return "[]"
	}

	max := 0
	for last := n; last != nil; last = last.Right {
		max = getIndex(last.Item)
	}
	repeat := 0
	for ; max > 0; max /= 10 {
		repeat++
	}
	vs := "|  " + strings.Repeat(" ", repeat)
	es := "   " + strings.Repeat(" ", repeat)
	format := "%v[%0" + strconv.Itoa(repeat) + "v%v]\n"

	height := n.getHeight()
	rowLength := height*4 - 1
	rowCount := 2<<(height-1) - 1

	var sb strings.Builder
	sb.Grow(rowLength * rowCount)

	fmt.Fprintln(&sb, "")

	var write func(io.Writer, string, int, *Node)
	write = func(b io.Writer, prefix string, dir int, n *Node) {
		if n == nil {
			return
		}
		if n.Right != nil {
			if dir < 0 {
				write(b, prefix+vs, 1, n.Right)
			} else {
				write(b, prefix+es, 1, n.Right)
			}
		}
		nb := ".0"
		if n.b > 0 {
			nb = "+" + strconv.FormatInt(int64(n.b), 10)
		} else if n.b < 0 {
			nb = strconv.FormatInt(int64(n.b), 10)
		}
		fmt.Fprintf(b, format, prefix, getIndex(n.Item), nb)
		if n.Left != nil {
			if dir > 0 {
				write(b, prefix+vs, -1, n.Left)
			} else {
				write(b, prefix+es, -1, n.Left)
			}
		}
	}
	write(&sb, "", 0, n)

	return sb.String()
}

func (n *Node) GetItems(u *User, getIndex func(*User) int) []*User {
	nodes := make([]*User, 0, n.Length())
	if n == nil {
		return nodes
	}

	id := getIndex(u)
	var appendItem func(*Node)
	appendItem = func(n *Node) {
		if n.Left != nil {
			appendItem(n.Left)
		}
		if getIndex(n.Item) >= id {
			nodes = append(nodes, n.Item)
		}
		if n.Right != nil {
			appendItem(n.Right)
		}
	}
	appendItem(n)

	return nodes
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
