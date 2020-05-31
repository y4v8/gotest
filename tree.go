package main

type Node struct {
	User   *User
	Prev   *Node
	Next   *Node
	Parent *Node // remove
	b      int
}

type Tree struct {
	Root *Node
	Last *Node
}

func NewTree() *Tree {
	return &Tree{}
}

func (t *Tree) Insert(u *User) {

}

func (t *Tree) Delete(u *User) {

}

// func RotateLeft
