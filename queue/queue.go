package main

import (
	"fmt"
)

type Queue struct {
	container   []int
	front int
	rear  int
	size  int
}


func NewQueue (k int) *Queue {
	return &Queue{
		make([]int, k),
		0,
		0,
		0,
	}
}


func (this *Queue) EnQueue(value int) bool {
	if this.container == nil || this.IsFull() {
		return false
	}
	this.container[this.rear%len(this.container)] = value
	this.rear = this.rear%len(this.container) + 1
	this.size++
	return true
}


func (this *Queue) DeQueue() (bool,int) {
	if this.container == nil || this.IsEmpty() {
		return false,0
	} else {
		ret := this.container[this.front%len(this.container)]
		this.front = this.front%len(this.container) + 1
		this.size--
		return true,ret
	}
}


func (this *Queue) IsEmpty() bool {
	if this.size == 0 {
		return true
	}
	return false
}

func (this *Queue) IsFull() bool {
	if this.size == len(this.container) {
		return true
	}
	return false
}

func main() {
	queue := NewQueue(5)
	// 循环3次，每次添加5个元素，再出队三个元素
	for i:=0; i<3; i++{
		for j:=0;j<5;j++ {
			queue.EnQueue(j)
		}
		for k:=0;k<3;k++ {
			fmt.Println(queue.DeQueue())
		}
	}
}