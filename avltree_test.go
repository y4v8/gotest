package gotest

import (
	"math/rand"
	"sort"
	"strconv"
	"strings"
	"testing"
)

func sampleItems(n int) []Item {
	users := make([]Item, n)
	for i := range users {
		users[i].ID = i + 1
		users[i].Name = "name" + strconv.Itoa(i+1)
		users[i].UpdateID = i + 1
	}

	rand.Shuffle(len(users), func(i, j int) {
		users[i].UpdateID, users[j].UpdateID = users[j].UpdateID, users[i].UpdateID
	})

	return users
}

func BenchmarkTree(b *testing.B) {
	sample := sampleItems(2000)
	u := &Item{
		ID:       88,
		UpdateID: 1,
	}
	_ = u

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tree := NewAVLTree(func(u *Item) int {
			return u.UpdateID
		})
		for u := range sample {
			tree.Insert(&sample[u])
		}
		node := tree.Get(u)
		if node == nil || node.Item.UpdateID != u.UpdateID {
			b.Fatalf("error: %v", tree)
		}
	}
}

func BenchmarkArray(b *testing.B) {
	sample := sampleItems(2000)
	u := &Item{
		ID:       88,
		UpdateID: 88,
	}
	_ = u

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		users := make([]User, 0, len(sample))
		for u := range sample {
			users = append(users, (User)(sample[u]))
			sort.Sort(ByUpdateID(users)) // TODO:
		}
		n := sort.Search(len(users), func(i int) bool { return users[i].UpdateID >= u.UpdateID })
		if n >= len(users) || users[n].UpdateID != u.UpdateID {
			b.Fatalf("error: %v", n)
		}
	}
}

func BenchmarkSortedArray(b *testing.B) {
	sample := sampleItems(2000)
	u := &Item{
		ID:       88,
		UpdateID: 88,
	}
	_ = u

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		users := make([]Item, 0, len(sample))
		for k := range sample {
			n := sort.Search(k, func(i int) bool { return users[i].UpdateID >= sample[k].UpdateID })
			users = append(users, sample[k])
			copy(users[n+1:], users[n:])
			users[n] = sample[k]
		}
		n := sort.Search(len(users), func(i int) bool { return users[i].UpdateID >= u.UpdateID })
		if n >= len(users) || users[n].UpdateID != u.UpdateID {
			b.Fatalf("error: %v", n)
		}
	}
}

func testAVLTreeGet(t *testing.T, tree *AVLTree, slen int, getIndex func(*Item) int) {
	tlen := tree.Root.Length()
	if slen != tlen {
		t.Errorf("length is %v, expect %v", tlen, slen)
	}

	u := &Item{ID: 88, UpdateID: 88}
	node := tree.Get(u)
	if node == nil || getIndex(node.Item) != getIndex(u) {
		t.Errorf("item with index %v is not found", getIndex(u))
	}

	u = &Item{ID: 8888, UpdateID: 8888}
	node = tree.Get(u)
	if node != nil {
		t.Errorf("item with index %v must do not be found", getIndex(u))
	}

	min := getIndex(u) - 1
	items := tree.GetItems(u)
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
	users := sampleItems(100)
	getIndex := func(u *Item) int { return u.UpdateID }

	tree := NewAVLTree(getIndex)
	for i := range users {
		tree.Insert(&users[i])
	}
	testAVLTreeGet(t, tree, len(users), getIndex)
}

func TestDelete(t *testing.T) {
	users := sampleItems(100)
	getIndex := func(u *Item) int { return u.UpdateID }

	tree := NewAVLTree(getIndex)
	for i := range users {
		tree.Insert(&users[i])
	}

	u := &Item{ID: 11, UpdateID: 56}
	tree.Delete(u)

	testAVLTreeGet(t, tree, len(users)-1, getIndex)
}

func TestGetItems(t *testing.T) {
	users := sampleItems(100)
	getIndex := func(u *Item) int { return u.UpdateID }

	tree := NewAVLTree(getIndex)
	for i := range users {
		tree.Insert(&users[i])
	}

	u := &Item{ID: 11, UpdateID: 93}
	tree.Delete(u)
	u = &Item{ID: 11, UpdateID: 92}
	tree.Delete(u)
	u = &Item{ID: 11, UpdateID: 94}
	tree.Delete(u)

	u = &Item{ID: 11, UpdateID: 90}

	items := tree.GetItems(u)
	indices := make([]string, len(items))
	for i := range items {
		indices[i] = strconv.Itoa(getIndex(items[i]))
	}

	result := strings.Join(indices, ",")
	expect := "90,91,95,96,97,98,99,100"
	if result != expect {
		t.Errorf("indices are '%v', expect '%v'", result, expect)
	}
}
