package Linked

import (
	"fmt"
	"strings"
)

type Node struct {
	E int
	Next *Node
}

type List struct {
	dummyHead *Node
	size int
}

func (l List) Head() *Node {
	return l.dummyHead
}


func (l List) Size() int {
	return l.size
}
func initNode (e int) *Node {
	return &Node{
		E:e,
		Next:nil,
	}
}
func InitList () *List {
	return &List{
		dummyHead:initNode(0),
		size:0,
	}
}
func (this *List) IsEmpty() bool {
	return this.size == 0
}
// 在链表的第index索引个元素后插入元素,索引从0开始
func (this *List) AddIndex (index,e int) {
	if index > this.size || index < 0 {
		panic("索引越界，不能插入了")
	}
	prev := this.dummyHead
	node := initNode(e)

	for i:=0;i<index;i++ {
		prev = prev.Next
	}
	node.Next = prev.Next
	prev.Next = node
	this.size++

}
// 在链表头添加元素
func (this *List) AddFirst (e int) {
	this.AddIndex(0,e)
}
// 在链表尾部添加节点
func (this *List) AddLast (e int) {
	this.AddIndex(this.size,e)
}
// 在链表中查询第index个元素
func (this *List) Get (index int) int {
	if index > this.size || index < 0 {
		panic("索引越界，不能查询")
	}
	cur := this.dummyHead.Next
	for i:=0;i<index;i++ {
		cur = cur.Next
	}
	return cur.E
}
func (this *List) GetFirst (index int) int{
	return this.Get(0)
}
func (this *List) GetLast (index int) int{
	return this.Get(this.size-1)
}
// 在链表index个位置中放入元素e
func (this *List) Set (index,e int) {
	if index > this.size || index < 0 {
		panic("索引越界，不能置入")
	}
	cur := this.dummyHead.Next
	for i:=0;i<index;i++ {
		cur = cur.Next
	}
	cur.E = e
}
// 在链表中查询是否包括元素e
func (this *List) Contains (e int) bool {
	cur := this.dummyHead.Next
	for cur!=nil {
		if cur.E == e{
			return true
		}
		cur = cur.Next
	}
	return false
}
// 在链表中删除元素
func (this *List) Remove (index int) int {
	if index > this.size || index < 0 {
		panic("索引越界，不能删除")
	}
	prev := this.dummyHead
	for i:=0;i<index;i++ {
		prev = prev.Next
	}
	retNode := prev.Next
	prev.Next = retNode.Next
	this.size--
	return retNode.E
}
func (this *List) RemoveFirst (index int) int{
	return this.Remove(0)
}
func (this *List) RemoveLast (index int) int{
	return this.Remove(this.size-1)
}
// 删除元素E
func (this *List) RemoveElement (e int) {
	prev := this.dummyHead
	for prev.Next != nil {
		if prev.E == e {
			break
		}
		prev = prev.Next
	}
	if prev.Next != nil {
		DelNode := prev.Next
		prev.Next = DelNode.Next
		DelNode = nil
	}
}
func (this *List) String () string {
	var builder strings.Builder
	cur := this.dummyHead.Next
	for cur != nil {
		fmt.Fprintf(&builder,"%d -> ",cur.E)
		cur = cur.Next
	}
	fmt.Fprintf(&builder,"NULL")
	return builder.String()
}


