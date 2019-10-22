package DoubleLinked

import (
	"fmt"
	"strings"
)

type Node struct {
	key interface{}
	value interface{}
	prev,next *Node
}


func (this Node) String() string {
	builder := strings.Builder{}
	fmt.Fprintf(&builder,"{%v : %v}",this.key,this.value)
	return builder.String()
}

func InitNode(key,value interface{}) *Node {
	return &Node{
		key:key,
		value:value,
	}
}

type List struct {
	capacity int
	head *Node
	tail *Node
	size int
}

func InitList(capcity int) *List {
	return &List{
		capacity:capcity,
		size:0,
	}
}

func (this *List) addHead (node *Node) *Node {
	if this.head == nil {
		this.head = node
		this.tail = node
		this.head.prev = nil
		this.tail.next = nil
	} else {
		node.next = this.head
		this.head.prev = node
		this.head = node
		this.head.prev = nil
	}
	this.size++
	return node
}

func (this *List)addTail(node *Node) *Node {
	if this.tail == nil {
		this.tail = node
		this.head = node
		this.head.prev = nil
		this.tail.next = nil
	} else {
		this.tail.next = node
		node.prev = this.tail
		this.tail = node
		this.tail.next = nil
	}
	this.size++
	return node
}

func (this *List) removeTail() *Node {
	if this.tail == nil {
		return nil
	}
	node := this.tail
	if node.prev != nil {
		this.tail = node.prev
		this.tail.next = nil
	} else {
		this.tail = nil
		this.head = nil
	}
	this.size--
	return node
}

func (this *List) removeHead() *Node {
	if this.head == nil {
		return nil
	}
	node := this.head
	if node.next != nil {
		this.head = node.next
		this.head.prev = nil
	} else {
		this.tail = nil
		this.head = nil
	}
	this.size--
	return node
}

func (this *List)remove(node *Node) *Node {
	// 如果node==nil,默认删除尾节点
	if node == nil {
		node = this.tail
	}
	if node == this.tail {
		this.removeTail()
	} else if node == this.head {
		this.removeHead()
	} else {
		node.next.prev = node.prev
		node.prev.next = node.next
		this.size--
	}
	return node
}
// 弹出头结点
func (this *List)Pop() *Node {
	return this.removeHead()
}

// 添加节点,默认添加到尾部
func (this *List)Append(node *Node) *Node {
	return this.addTail(node)
}
func (this *List)AppendToHead(node *Node) *Node {
	return this.addHead(node)
}

func (this *List)Remove(node *Node) *Node {
	return this.remove(node)
}

func (this *List)String() string {
	p := this.head
	builder := strings.Builder{}
	for p != nil {
		fmt.Fprintf(&builder,"%s",p)
		p = p.next
		if p != nil {
			fmt.Fprintf(&builder,"=>")
		}
	}
	return builder.String()
}