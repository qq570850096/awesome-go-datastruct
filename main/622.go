package main

type MyCircularQueue struct {
	arr   []int
	front int
	rear  int
	size  int
}

/** Initialize your data structure here. Set the size of the queue to be k. */
func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		make([]int, k),
		0,
		0,
		0,
	}
}

/** Insert an element into the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.arr == nil || (this.size == len(this.arr) && (this.rear)%len(this.arr) == this.front) {
		return false
	}
	this.arr[this.rear%len(this.arr)] = value
	this.rear = this.rear%len(this.arr) + 1
	this.size++
	return true
}

/** Delete an element from the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) DeQueue() bool {
	if this.arr == nil || this.size == 0 {
		return false
	} else {
		this.front = this.front%len(this.arr) + 1
		this.size--
		return true
	}
}

/** Get the front item from the queue. */
func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}
	return this.arr[this.front%len(this.arr)]
}

/** Get the last item from the queue. */
func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}
	last := (this.rear - 1 + len(this.arr)) % len(this.arr)
	return this.arr[last]
}

/** Checks whether the circular queue is empty or not. */
func (this *MyCircularQueue) IsEmpty() bool {
	if this.size == 0 {
		return true
	}
	return false
}

/** Checks whether the circular queue is full or not. */
func (this *MyCircularQueue) IsFull() bool {
	if this.size == len(this.arr) {
		return true
	}
	return false
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */
