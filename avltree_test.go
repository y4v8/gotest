package gotest

import (
	"testing"
)

func TestMath(t *testing.T) {
	p := 2<<(10-1) - 1
	if p != 1023 {
		t.Fatalf("p=%v expect 4", p)
	}
}

func TestAVLTree(t *testing.T) {
	tree := NewAVLTree(func(u *User) int {
		return u.ID
	})

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

	tree.Delete(u6)

	t.Error(tree.String())

}
