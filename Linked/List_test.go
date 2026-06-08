package Linked

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestList(t *testing.T) {
	list := InitList()
	list2 := InitList()
	for i := 0; i < 5; i++ {
		list.AddFirst(i)
		list2.AddFirst(rand.Int() % 17)
	}
	list2.Sort()
	list.Sort()
	fmt.Println(list, "\n", list2)
	fmt.Println(MergeTwoList(list2.Head(), list.Head()))
	fmt.Println(list)

}
