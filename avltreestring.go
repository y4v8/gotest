package gotest

import (
	"fmt"
	"strconv"
	"strings"
)

func (t *AVLTree) String() string {
	if t.Root == nil {
		return "[]"
	}

	max := 0
	for last := t.Root; last != nil; last = last.Right {
		max = t.getIndex(last.Item)
	}
	repeat := 0
	for ; max > 0; max /= 10 {
		repeat++
	}
	vs := "|  " + strings.Repeat(" ", repeat)
	es := "   " + strings.Repeat(" ", repeat)
	format := "%v[%0" + strconv.Itoa(repeat+2) + "v]\n"

	height := t.Root.getHeight()
	rowLength := int(height)*4 - 1
	rowCount := 2<<(height-1) - 1

	var sb strings.Builder
	sb.Grow(rowLength * rowCount)

	fmt.Fprintln(&sb, "")

	writeBuf := make([]byte, 0, 32)
	writeNode := func(n *Node, prefix string) {
		writeBuf = writeBuf[:0]
		writeBuf = strconv.AppendInt(writeBuf, int64(t.getIndex(n.Item)), 10)

		balance := int8(n.GetBalance())
		if balance > 0 {
			writeBuf = append(writeBuf, '+', byte(0x30+balance))
		} else if balance < 0 {
			writeBuf = append(writeBuf, '-', byte(0x30-balance))
		} else {
			writeBuf = append(writeBuf, '.', '0')
		}

		fmt.Fprintf(&sb, format, prefix, string(writeBuf))
	}

	var iterateNode func(string, int, *Node)
	iterateNode = func(prefix string, dir int, n *Node) {
		if n == nil {
			return
		}
		if n.Right != nil {
			if dir < 0 {
				iterateNode(prefix+vs, 1, n.Right)
			} else {
				iterateNode(prefix+es, 1, n.Right)
			}
		}

		writeNode(n, prefix)

		if n.Left != nil {
			if dir > 0 {
				iterateNode(prefix+vs, -1, n.Left)
			} else {
				iterateNode(prefix+es, -1, n.Left)
			}
		}
	}
	iterateNode("", 0, t.Root)

	return sb.String()
}
