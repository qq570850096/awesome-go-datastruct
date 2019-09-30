package Linked

import (
	"fmt"
	"testing"
)

func TestList(t *testing.T) {
	list := InitList()
	for i:=0;i<5 ; i++  {
		list.AddFirst(i)
	}
	fmt.Println(list.Contains(4))
	fmt.Println(list.Get(2))
	list.AddIndex(2,666)
	fmt.Println(list)
	list.Remove(2)
	fmt.Println(list)
}
