package Heap

type Queue interface {
	Enqueue(e int)
	Dequeue() int
	GetFront() int
	Size() int
	IsEmpty() bool
}

type PriorityQueue struct {
	heap *MaxHeap
}

func (this *PriorityQueue) Enqueue(e int) {
	this.heap.Add(e)
}

func (this *PriorityQueue) Dequeue() int {
	return this.heap.RemoveMax()
}

func (this *PriorityQueue) GetFront() int {
	return this.heap.FindMax()
}
func (this *PriorityQueue) Size() int {
	return this.heap.Size()
}

func (this *PriorityQueue) IsEmpty() bool {
	return this.heap.IsEmpty()
}



