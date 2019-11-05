package Linked

import (
	"testing"
)

func TestList(t *testing.T) {
	list := InitList()
	list2 := InitList()
	for i:=0;i<5 ; i=i+2  {
		list.AddLast(i)
	}
	for i:=0;i<6 ; i=i+2  {
		list2.AddLast(i+1)
	}
	t.Log(list)
	t.Log(list2)
	list2.Head().Next.Next = list.Head().Next.Next
	t.Log(list2)
	t.Log(*list2.CheckIntersect(list2.Head(),list.Head()))
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
