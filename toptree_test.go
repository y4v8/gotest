package gotest

import (
	"testing"
)

func TestTopTree(t *testing.T) {
	tree := NewTopTree()

	u1 := &User{ID: 1, Name: "usr1", UpdateID: 10}
	u2 := &User{ID: 2, Name: "usr2", UpdateID: 20}
	u3 := &User{ID: 3, Name: "usr3", UpdateID: 30}
	u4 := &User{ID: 4, Name: "usr4", UpdateID: 40}
	u5 := &User{ID: 5, Name: "usr5", UpdateID: 50}
	u6 := &User{ID: 6, Name: "usr6", UpdateID: 60}
	u7 := &User{ID: 7, Name: "usr7", UpdateID: 70}
	u8 := &User{ID: 8, Name: "usr8", UpdateID: 80}

	// tree.Append(u1)
	// tree.Append(u2)
	// tree.Append(u3)
	// tree.Append(u4)

	n1 := NewNode(u1, nil, nil)
	n2 := NewNode(u2, nil, nil)
	n3 := NewNode(u3, nil, nil)
	n4 := NewNode(u4, nil, nil)
	n5 := NewNode(u5, nil, nil)
	n6 := NewNode(u6, nil, nil)
	n7 := NewNode(u7, nil, nil)
	n8 := NewNode(u8, nil, nil)

	tree.Root = n5
	n5.Left = n3
	n5.Right = n7
	n3.Left = n2
	n3.Right = n4
	n2.Left = n1
	n7.Right = n8
	n7.Left = n6

	t.Error(tree.String())

}
