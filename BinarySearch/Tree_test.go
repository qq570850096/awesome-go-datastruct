package BinarySearch

import (
	"fmt"
	"testing"
)

func TestTree(t *testing.T) {
	tree := &Tree{}
	fmt.Println(tree.IsEmpty())
	var arr [6]int = [6]int{5,3,6,8,4,2}
	for _,v := range arr{
		tree.AddE(v)
	}
	tree.PreOrder()
	tree.PreOrderNR()
	fmt.Println(tree.FindMax())
	fmt.Println(tree.FindMin())
	fmt.Println(tree.DelMax())
	fmt.Println(tree.DelMin())
	fmt.Println(tree.size)
	fmt.Println(tree)
	tree.Remove(5)
	fmt.Println(tree)
}
