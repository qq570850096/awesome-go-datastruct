package Linked

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestList(t *testing.T) {
	list := InitList()
	list2 := InitList()
	for i:=0;i<5 ; i++  {
		list.AddFirst(i)
		list2.AddFirst(rand.Int()%17)
	}
	list2.Sort()
	list.Sort()
	fmt.Println(list,"\n",list2)
	fmt.Println(MergeTwoList(list2.Head(),list.Head()))
	fmt.Println(list)
	//for i:=9;i>0;i--{
	//	//	if list.FindLastK(i) == nil {
	//	//		t.Error("超过链表可表示长度")
	//	//	} else {
	//	//		t.Log(list.FindLastK(i).E)
	//	//	}
	//	//}
	//
	//	//t.Log(list.Contains(4))
	//	//t.Log(list.Get(2))
	//	//list.AddIndex(2,666)
	//	//t.Log(list)
	//	//list.Reverse()
	//	//t.Log(list)
	//	//list.RecursiveReverse()
	//	//t.Log(list)
	//	//list.InsertReverse()
	//	//t.Log(list)
	//	//list.ReversePrint(list.Head().Next)

}
