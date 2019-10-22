package Segment

import (
	"testing"
)

func TestTree(t *testing.T) {
	tree := &Tree{}
	arr  := []int{-2,0,3,-5,2,-1}
	tree.Init(arr,add)
	t.Log(tree)
	t.Log(tree.QueryLR(0,2))
	t.Log(tree.QueryLR(2,5))
}
func add (a,b int) int {
	return a+b
}
