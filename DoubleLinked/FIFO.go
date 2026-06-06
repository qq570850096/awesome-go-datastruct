package DoubleLinked

import "fmt"

type FIFOCache struct {
	capacity int
	size     int
	list     *List
	find     map[interface{}]*Node
	k        int
	// count records cache/page misses in the FIFO examples.
	count int
}

func InitFIFO(capacity int) *FIFOCache {
	return &FIFOCache{
		capacity: capacity,
		find:     map[interface{}]*Node{},
		list:     InitList(capacity),
	}
}

func (this *FIFOCache) GetCount() int {
	return this.count
}

func (this *FIFOCache) Get(key interface{}) interface{} {
	if node, ok := this.find[key]; ok {
		return node.value
	}
	if this.capacity == 0 {
		return -1
	}

	fmt.Println("发生了一次缺页中断")
	this.count++
	node := InitNode(key, -1)
	if this.size == this.capacity {
		oldNode := this.list.Pop()
		if oldNode != nil {
			delete(this.find, oldNode.key)
			this.size--
		}
	}
	this.list.Append(node)
	this.find[key] = node
	this.size++
	return -1
}

func (this *FIFOCache) Put(key, value interface{}) {
	if this.capacity == 0 {
		return
	}
	if node, ok := this.find[key]; ok {
		node.value = value
		return
	}

	if this.size == this.capacity {
		oldNode := this.list.Pop()
		if oldNode != nil {
			delete(this.find, oldNode.key)
			this.size--
		}
	}
	node := InitNode(key, value)
	this.list.Append(node)
	this.find[key] = node
	this.size++
}

func (this *FIFOCache) String() string {
	return this.list.String()
}
