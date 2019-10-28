package Linked

import (
	"testing"
)

func TestList(t *testing.T) {
	list := InitList()
	list2 := InitList()
	for i:=0;i<7 ; i++  {
		list.AddLast(i+1)
	}
	for i:=0;i<6 ; i++  {
		list2.AddFirst(i+1)
	}
	t.Log(list)
	list.Reorder()
	t.Log(list)
	//t.Log(list.Contains(4))
	//t.Log(list.Get(2))
	//list.AddIndex(2,666)
	//t.Log(list)
	//list.Reverse()
	//t.Log(list)
	//list.RecursiveReverse()
	//t.Log(list)
	//list.InsertReverse()
	//t.Log(list)
	//list.ReversePrint(list.Head().Next)

}
