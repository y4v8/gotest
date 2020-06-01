package gotest

type User struct {
	ID   int
	Name string
}

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
	parents := make([]*Node, 0, 16)
	parents = append(parents, t.Root)
	for i := 0; i < len(parents); i++ {
		if u.ID > parents[i].User.ID {
			//
		}
	}

}

func (t *Tree) Delete(u *User) {

}

func (n *Node) search(u *User) *Node {
	parents := make([]*Node, 0, 16)
	parents = append(parents, n)
	for i := 0; i < len(parents); i++ {
		if u.ID > n.User.ID {
			//
		}
	}
	return nil
}

func (n *Node) rotateLeft() {

}

func (n *Node) rotateRight() {

}
