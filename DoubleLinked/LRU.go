package DoubleLinked

import "fmt"

type LRUCache struct {
	capacity int
	find map[interface{}]*Node
	list *List
	k int
	count int
}

func InitLRU(capacity int) *LRUCache {
	return &LRUCache{
		capacity:capacity,
		list:InitList(capacity),
		find:make(map[interface{}]*Node),
	}
}

func (this *LRUCache)Get(key interface{}) interface{} {
	if value,ok := this.find[key];ok{
		node := value
		this.list.Remove(node)
		this.list.AppendToHead(node)
		return node.value
	} else {
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
	}
}

func (this *LRUCache)Put(key,value interface{})  {
	if v,ok := this.find[key];ok {
		node := v
		this.list.Remove(node)
		node.value = value
		this.list.AppendToHead(node)
	} else {
		node := InitNode(key,value)
		// 缓存已经满了
		if this.list.size >= this.list.capacity {
			oldNode := this.list.Remove(nil)
			delete(this.find,oldNode.value)
		}
		this.list.AppendToHead(node)
		this.find[key] = node
	}
}

func (this *LRUCache)String() string {
	return this.list.String()
}