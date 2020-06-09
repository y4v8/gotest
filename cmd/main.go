package main

import (
	"fmt"
	"sort"

	"github.com/y4v8/gotest"
)

func main() {
	users := gotest.SampleUsers(100000)

	sort.Sort(gotest.ByUpdateID(users))

	n := sort.Search(len(users), func(i int) bool { return users[i].UpdateID >= 2222 })

	slice := users[n : n+10]
	for i := range slice {
		fmt.Printf("%#v\n", slice[i])
	}
}
