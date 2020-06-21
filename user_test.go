package gotest

import (
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	"testing"
)

func SampleUsers3(n int) []User {
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

func BenchmarkTree(b *testing.B) {
	sample := SampleUsers(100)
	u := &User{
		ID:       88,
		UpdateID: 1,
	}
	_ = u

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tree := NewAVLTree(func(u *User) int {
			return u.UpdateID
		})
		for u := range sample {
			tree.Insert(&sample[u])
		}
		node := tree.Get(u)
		if node == nil || node.Item.UpdateID != u.UpdateID {
			b.Fatalf("error: %v", tree)
		}

		min := u.UpdateID
		items := tree.Root.GetItems(u, tree.getIndex)
		for _, item := range items {
			if item.UpdateID > min {
				min = item.UpdateID
			} else {
				b.Fatal(min, ">=", item.UpdateID)
			}
			fmt.Println(item)
		}
		b.Fatal(1)
	}
}

func BenchmarkArray(b *testing.B) {
	sample := SampleUsers(10000)
	u := &User{
		ID:       88,
		UpdateID: 88,
	}
	_ = u

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		users := make([]User, 0, len(sample))
		for u := range sample {
			users = append(users, sample[u])
			sort.Sort(ByUpdateID(users))
		}
		n := sort.Search(len(users), func(i int) bool { return users[i].UpdateID >= u.UpdateID })
		if n >= len(users) || users[n].UpdateID != u.UpdateID {
			b.Fatalf("error: %v", n)
		}
	}
}
