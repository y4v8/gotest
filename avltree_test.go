package gotest

import (
	"math/rand"
	"strconv"
	"testing"
)

func SampleUsers(n int) []User {
	users := make([]User, n)
	for i := range users {
		users[i].ID = i + 1
		users[i].Name = "name" + strconv.FormatInt(int64(i+1), 10)
		users[i].UpdateID = i + 1
	}

	rand.Shuffle(len(users), func(i, j int) {
		users[i].UpdateID, users[j].UpdateID = users[j].UpdateID, users[i].UpdateID
	})

	return users
}

func testAVLTreeGet(t *testing.T, tree *AVLTree, slen int, getIndex func(*User) int) {
	tlen := tree.Root.Length()
	if slen != tlen {
		t.Errorf("length is %v, expect %v", tlen, slen)
	}

	u := &User{ID: 88, UpdateID: 88}
	node := tree.Get(u)
	if node == nil || getIndex(node.Item) != getIndex(u) {
		t.Errorf("item with index %v is not found", getIndex(u))
	}

	min := getIndex(u) - 1
	items := tree.Root.GetItems(u, tree.getIndex)
	for _, item := range items {
		index := getIndex(item)
		if index > min {
			min = index
		} else {
			t.Errorf("items are not sorted by index - %v,%v", min, index)
			break
		}
	}
}

func TestGet(t *testing.T) {
	users := SampleUsers(100)
	getIndex := func(u *User) int { return u.UpdateID }

	tree := NewAVLTree(getIndex)
	for i := range users {
		tree.Insert(&users[i])
	}
	// t.Fatal(tree)
	testAVLTreeGet(t, tree, len(users), getIndex)
}

func TestDelete(t *testing.T) {
	users := SampleUsers(100)
	getIndex := func(u *User) int { return u.UpdateID }

	tree := NewAVLTree(getIndex)
	for i := range users {
		tree.Insert(&users[i])
	}
	// t.Fatal(tree)

	u := &User{ID: 46, UpdateID: 56}
	_ = u
	tree.Delete(u)

	// testAVLTreeGet(t, tree, len(users)-1, getIndex)
}

func TestDelete2(t *testing.T) {
	users := SampleUsers(100)
	getIndex := func(u *User) int { return u.UpdateID }

	tree := NewAVLTree(getIndex)
	for i := range users {
		tree.Insert(&users[i])
	}
	// t.Fatal(tree)

	u := &User{ID: 11, UpdateID: 3}
	_ = u
	tree.Delete(u)

	testAVLTreeGet(t, tree, len(users)-1, getIndex)
}
