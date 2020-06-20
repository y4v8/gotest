package gotest

import (
	"math/rand"
	"strconv"
)

type User struct {
	ID       int
	Name     string
	UpdateID int
}

type ByID []User

func (a ByID) Len() int           { return len(a) }
func (a ByID) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByID) Less(i, j int) bool { return a[i].ID < a[j].ID }

type ByName []User

func (a ByName) Len() int           { return len(a) }
func (a ByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByName) Less(i, j int) bool { return a[i].Name < a[j].Name }

type ByUpdateID []User

func (a ByUpdateID) Len() int           { return len(a) }
func (a ByUpdateID) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByUpdateID) Less(i, j int) bool { return a[i].UpdateID < a[j].UpdateID }

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
