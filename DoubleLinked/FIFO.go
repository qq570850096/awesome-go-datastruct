package DoubleLinked

import "fmt"

type FIFOCache struct {
	capacity int
	size int
	list *List
	find map[interface{}]*Node
	k int
	// count记录缺页中断的次数
	count int
}

func InitFIFO(capacity int) *FIFOCache {
	return &FIFOCache{
		capacity: capacity,
		find:     map[interface{}]*Node{},
		list:     InitList(capacity),
	}
}

func (this *FIFOCache)GetCount() int {
	return this.count
}

func (this *FIFOCache)Get(key interface{}) interface{} {
	if value,ok := this.find[key];!ok{
		this.k = this.k%this.capacity
		node := this.list.head
		for i:=0;i<this.k;i++ {
			node = node.next
		}
		fmt.Println("发生了一次缺页中断")
		delete(this.find,node.key)
		node.key = key
		this.find[key] = node
		this.k++
		this.count++
		return -1
	} else {
		node := value
		return node.value
	}
}

func (this *FIFOCache)Put(key,value interface{})  {
	if this.capacity == 0 {
		return
	}
	if v,ok := this.find[key];ok {
		node := v
		this.list.Remove(node)
		node.value = value
		this.list.Append(node)
	} else {

		if this.size == this.capacity {
			node := this.list.Pop()
			delete(this.find,node.key)
			this.size--
		}
		node := InitNode(key,value)
		this.list.Append(node)
		this.find[key] = node
		this.size++
	}
}

func (this *FIFOCache)String() string {
	return this.list.String()
}