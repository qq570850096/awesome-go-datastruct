package Heap

type MaxHeap struct {
	arr []int
}

func (heap *MaxHeap) InitHeapWithArray(arr []int) *MaxHeap {
	heap.arr = make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		heap.arr[i] = arr[i]
	}
	for i := heap.parent(len(arr) - 1); i >= 0; i-- {
		heap.siftDown(i)
	}
	return heap
}

// 返回堆中元素个数
func (heap *MaxHeap) Size() int {
	return len(heap.arr)
}

// 返回堆是否为空
func (heap *MaxHeap) IsEmpty() bool {
	if len(heap.arr) == 0 {
		return true
	}
	return false
}

// 计算父节点
func (heap *MaxHeap) parent(index int) int {
	if index == 0 {
		panic("0节点没有父亲节点")
	}
	return (index - 1) / 2
}

// 计算左孩子索引
func (heap *MaxHeap) leftChild(index int) int {
	return index*2 + 1
}

// 计算右孩子索引
func (heap *MaxHeap) rightChild(index int) int {
	return index*2 + 2
}

// 增加元素
func (heap *MaxHeap) Add(e int) {
	heap.arr = append(heap.arr, e)
	heap.siftUp(heap.Size() - 1)
}

// 数据上浮
func (heap *MaxHeap) siftUp(size int) {

	for size > 0 && heap.arr[heap.parent(size)] < heap.arr[size] {
		heap.arr[heap.parent(size)], heap.arr[size] = heap.arr[size], heap.arr[heap.parent(size)]
		size = heap.parent(size)
	}
}

// 取出堆中最大的元素，并且替换成元素e
func (heap *MaxHeap) Replace(e int) int {
	ret := heap.FindMax()
	heap.arr[0] = e
	heap.siftDown(0)
	return ret
}

// 删除最大元素
func (heap *MaxHeap) RemoveMax() int {
	if heap.Size() == 0 {
		panic("堆是空的，不能继续删除")
	}
	ret := heap.arr[0]
	heap.arr[0], heap.arr[heap.Size()-1] = heap.arr[heap.Size()-1], heap.arr[0]
	heap.arr = heap.arr[:heap.Size()-1]
	if heap.Size() > 0 {
		heap.siftDown(0)
	}
	return ret
}

// 数据下沉
func (heap *MaxHeap) siftDown(index int) {
	var (
		j int
	)
	for heap.leftChild(index) < heap.Size() {
		j = heap.leftChild(index)
		// 先判断右孩子的索引是否越界，如果不越界继续判断左右孩子哪个大
		if j+1 < heap.Size() && heap.arr[j+1] > heap.arr[j] {
			// 此时arr[j]将是左孩子和右孩子中最大的那个值
			j = heap.rightChild(index)
		}
		if heap.arr[index] >= heap.arr[j] {
			break
		} else {
			heap.arr[index], heap.arr[j] = heap.arr[j], heap.arr[index]
			index = j
		}
	}
}

func (heap *MaxHeap) FindMax() int {
	if heap.Size() == 0 {
		panic("堆是空的，不能查询了")
	}
	return heap.arr[0]
}
