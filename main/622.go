package main

type MyCircularQueue struct {
	arr   []int
	front int
	rear  int
	size  int
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		make([]int, k),
		0,
		0,
		0,
	}
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.arr == nil || (this.size == len(this.arr) && (this.rear)%len(this.arr) == this.front) {
		return false
	}
	this.arr[this.rear%len(this.arr)] = value
	this.rear = this.rear%len(this.arr) + 1
	this.size++
	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.arr == nil || this.size == 0 {
		return false
	} else {
		this.front = this.front%len(this.arr) + 1
		this.size--
		return true
	}
}

func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}
	return this.arr[this.front%len(this.arr)]
}

func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}
	last := (this.rear - 1 + len(this.arr)) % len(this.arr)
	return this.arr[last]
}

func (this *MyCircularQueue) IsEmpty() bool {
	if this.size == 0 {
		return true
	}
	return false
}

func (this *MyCircularQueue) IsFull() bool {
	if this.size == len(this.arr) {
		return true
	}
	return false
}
