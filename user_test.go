package gotest

import (
	"sort"
	"testing"
)

func BenchmarkTree(b *testing.B) {
	sample := SampleUsers(10000)
	u := &User{
		ID:       88,
		UpdateID: 88,
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
	}
	b.StopTimer()
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
	b.StopTimer()
}
