package main

import (
	"fmt"

	"github.com/y4v8/gotest"
)

func main() {
	u := &gotest.User{
		ID:   1,
		Name: "user",
	}

	fmt.Println(u)
}
