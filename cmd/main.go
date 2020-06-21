package main

import (
	"fmt"

	. "github.com/y4v8/gotest"
)

func main() {
	tree := NewAVLTree(func(u *User) int {
		return u.UpdateID
	})

	u1 := &User{ID: 1, Name: "usr1", UpdateID: 10}
	u2 := &User{ID: 2, Name: "usr2", UpdateID: 20}
	u3 := &User{ID: 3, Name: "usr3", UpdateID: 30}
	u4 := &User{ID: 4, Name: "usr4", UpdateID: 40}
	tree.Insert(u1)
	tree.Insert(u2)
	tree.Insert(u3)

	fmt.Println(tree.Root.Item.ID)
	fmt.Println(tree)

	tree.Insert(u4)
}
