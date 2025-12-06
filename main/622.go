package main

type MyCircularQueue struct {
	arr   []int
	front int
	rear  int
	size  int
}

/** 初始化循环队列，设置容量为 k。 */
func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		make([]int, k),
		0,
		0,
		0,
	}
}

/** 向循环队列插入元素，成功返回 true。 */
func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.arr == nil || (this.size == len(this.arr) && (this.rear)%len(this.arr) == this.front) {
		return false
	}
	this.arr[this.rear%len(this.arr)] = value
	this.rear = this.rear%len(this.arr) + 1
	this.size++
	return true
}

/** 从循环队列删除一个元素，成功返回 true。 */
func (this *MyCircularQueue) DeQueue() bool {
	if this.arr == nil || this.size == 0 {
		return false
	} else {
		this.front = this.front%len(this.arr) + 1
		this.size--
		return true
	}
}

/** 获取队首元素。 */
func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}
	return this.arr[this.front%len(this.arr)]
}

/** 获取队尾元素。 */
func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}
	last := (this.rear - 1 + len(this.arr)) % len(this.arr)
	return this.arr[last]
}

/** 判断循环队列是否为空。 */
func (this *MyCircularQueue) IsEmpty() bool {
	if this.size == 0 {
		return true
	}
	return false
}

/** 判断循环队列是否已满。 */
func (this *MyCircularQueue) IsFull() bool {
	if this.size == len(this.arr) {
		return true
	}
	return false
}

/**
 * 使用示例：
 * obj := Constructor(k); // 创建队列
 * param_1 := obj.EnQueue(value); // 入队
 * param_2 := obj.DeQueue(); // 出队
 * param_3 := obj.Front(); // 查看队首
 * param_4 := obj.Rear(); // 查看队尾
 * param_5 := obj.IsEmpty(); // 判空
 * param_6 := obj.IsFull(); // 判满
 */
