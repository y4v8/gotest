package gotest

import (
	"sort"
	"testing"
)

func Test1(t *testing.T) {
	//
}

func BenchmarkArray(b *testing.B) {
	users := SampleUsers(10000)
	updateID := 222

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sort.Sort(ByUpdateID(users))
		n := sort.Search(len(users), func(i int) bool { return users[i].UpdateID >= updateID })
		if n >= len(users) || users[n].UpdateID != updateID {
			b.Fatalf("error: %v", n)
		}
	}
	b.StopTimer()
}
