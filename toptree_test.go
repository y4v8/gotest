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

func TestTopTree1(t *testing.T) {
	tree := NewTopTree()
	tree.Root = NewNode(
		&User{ID: 1, Name: "usr1", UpdateID: 10},
		nil, nil,
	)
	tree.Root.Right = NewNode(
		&User{ID: 2, Name: "usr2", UpdateID: 20},
		nil, nil,
	)
	tree.Root.Right.Right = NewNode(
		&User{ID: 3, Name: "usr3", UpdateID: 30},
		nil, nil,
	)
	tree.Root.Right.Right.Left = NewNode(
		&User{ID: 4, Name: "usr4", UpdateID: 40},
		nil, nil,
	)

	node, n := tree.searchNode(111)
	t.Fatalf("node:%v, n:%v", node.GetID(), n)
	t.Fatal(tree.String())
}

func TestCmp(t *testing.T) {
	var cmp1, cmp2 Comparer
	cmp1 = &Ucmp{ID: 2}
	cmp2 = &Ucmp{ID: 2}

	r := cmp1.Compare(cmp2)
	t.Error(r)
}

func TestTopTree2(t *testing.T) {

	cmp := func(u1, u2 *User) int {
		//
		return 0
	}
	_ = cmp
	cmi := func(i1 int) func(int) int {
		return func(i2 int) int {
			return i1 - i2
		}
	}
	_ = cmi

	tree := NewTopTree()

	u1 := &User{ID: 1, Name: "usr1", UpdateID: 10}
	u2 := &User{ID: 2, Name: "usr2", UpdateID: 20}
	u3 := &User{ID: 3, Name: "usr3", UpdateID: 30}
	u4 := &User{ID: 4, Name: "usr4", UpdateID: 40}
	u5 := &User{ID: 5, Name: "usr5", UpdateID: 50}
	u6 := &User{ID: 6, Name: "usr6", UpdateID: 60}
	u7 := &User{ID: 7, Name: "usr7", UpdateID: 70}
	u8 := &User{ID: 8, Name: "usr8", UpdateID: 80}

	tree.Insert(u1)
	tree.Insert(u2)
	tree.Insert(u3)
	tree.Insert(u4)
	tree.Insert(u5)
	tree.Insert(u6)
	tree.Insert(u7)
	tree.Insert(u8)

	t.Error(tree.String())

}
