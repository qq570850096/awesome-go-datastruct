package Linked

import (
	"testing"
)

func TestList(t *testing.T) {
	list := InitList()
	for i:=0;i<7 ; i++  {
		list.AddLast(i+1)
	}
	loop := list.FindLastK(1)
	loop.Next = list.Head().Next.Next
	meet := list.FindLoop()
	if meet == nil {
		t.Log("此链表无环")
	} else {
		t.Log("此链表有环")
		entry := list.FindLoopEntryNode(meet)
		t.Log(*entry)
	}
	for i:=0;i<=7 ; i++  {
		t.Log(*loop)
		loop = loop.Next
	}
	//for i:=9;i>0;i--{
	//	if list.FindLastK(i) == nil {
	//		t.Error("超过链表可表示长度")
	//	} else {
	//		t.Log(list.FindLastK(i).E)
	//	}
	//}

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
