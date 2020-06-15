package gotest

import (
	"testing"
)

func TestTopTree(t *testing.T) {
	tree := NewTopTree()

	u1 := &User{ID: 1, Name: "usr1", UpdateID: 10}
	u2 := &User{ID: 2, Name: "usr2", UpdateID: 20}
	u3 := &User{ID: 3, Name: "usr3", UpdateID: 30}
	u4 := &User{ID: 4, Name: "usr44", UpdateID: 40}

	// tree.Append(u1)
	// tree.Append(u2)
	// tree.Append(u3)
	// tree.Append(u4)

	n1 := NewNode(u1, nil, nil)
	n2 := NewNode(u2, nil, nil)
	n3 := NewNode(u3, nil, nil)
	n4 := NewNode(u4, nil, nil)
	n1.Left = n2
	n1.Right = n3
	n3.Right = n4

	t.Error(tree)

}
