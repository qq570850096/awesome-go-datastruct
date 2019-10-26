package Linked

import (
	"testing"
)

func TestList(t *testing.T) {
	list := InitList()
	for i:=0;i<5 ; i++  {
		list.AddFirst(i)
	}
	t.Log(list.Contains(4))
	t.Log(list.Get(2))
	list.AddIndex(2,666)
	t.Log(list)
	list.Reverse()
	t.Log(list)
	list.RecursiveReverse()
	t.Log(list)
	list.InsertReverse()
	t.Log(list)
	list.ReversePrint(list.Head().Next)
}
