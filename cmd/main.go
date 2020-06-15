package main

import (
	"fmt"
	"sort"

	"github.com/y4v8/gotest"
	. "github.com/y4v8/gotest"
)

func main() {
	tree := NewTopTree()

	u1 := &User{ID: 1, Name: "usr1", UpdateID: 10}
	u2 := &User{ID: 2, Name: "usr2", UpdateID: 20}
	u3 := &User{ID: 3, Name: "usr3", UpdateID: 30}
	u4 := &User{ID: 4, Name: "usr44", UpdateID: 40}
	tree.Append(u1)
	tree.Append(u2)
	tree.Append(u3)

	fmt.Println(tree.Root.GetID())
	fmt.Println(tree)

	tree.Append(u4)
}

func main2() {
	users := gotest.SampleUsers(100000)

	sort.Sort(gotest.ByUpdateID(users))

	n := sort.Search(len(users), func(i int) bool { return users[i].UpdateID >= 2222 })

	slice := users[n : n+10]
	for i := range slice {
		fmt.Printf("%#v\n", slice[i])
	}
}
