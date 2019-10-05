package Segment

import (
	"fmt"
	"testing"
)

func TestTree(t *testing.T) {
	tree := &Tree{}
	arr  := []int{-2,0,3,-5,2,-1}
	tree.Init(arr,add)
	fmt.Println(tree)
	fmt.Println(tree.QueryLR(0,2))
	fmt.Println(tree.QueryLR(2,5))
}
func add (a,b int) int {
	return a+b
}
