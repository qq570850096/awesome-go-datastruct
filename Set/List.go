package Set

import "algo/Linked"
// 集合，基于链表实现
type ListSet struct {
	list *Linked.List
}
func InitListSet () *ListSet{
	return &ListSet{
		list:Linked.InitList(),
	}
}
func (this *ListSet) Add(e int) {
	if !this.Contains(e) {
		this.list.AddFirst(e)
	}
}

func (this *ListSet) IsEmpty() bool {
	return this.list.IsEmpty()
}

func (this *ListSet) Contains(e int) bool {
	return this.list.Contains(e)
}

func (this *ListSet) GetSize() int {
	return this.list.Size()
}

func (this *ListSet) Remove(e int)  {
	this.list.RemoveElement(e)
}
