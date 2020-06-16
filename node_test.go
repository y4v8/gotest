package gotest

import (
	"testing"
)

func TestNodeStringer(t *testing.T) {
	u1 := &User{ID: 0, Name: "usr1", UpdateID: 10}
	_ = u1
	n1 := NewNode(nil, nil, nil)
	//var n1 *Node
	r := "0, left:0, right:0"
	s := n1.String()
	if s != r {
		t.Errorf("Instead '%v' expect '%v'", r, s)
	}
}
