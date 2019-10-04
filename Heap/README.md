## 堆和优先队列

### 优先队列

* 普通队列：先进先出；后进后出

* 优先队列：出队顺序和入队顺序无关；和优先级有关

**队列**是一种**FIFO**（First-In-First-Out）先进先出的数据结构，对应于生活中的排队的场景，排在前面的人总是先通过，**依次进行**

**优先队列**是特殊的队列，从“优先”一词，可看出**有“插队现象”**。比如在火车站排队进站时，就会有些比较急的人来插队，他们就在前面先通过验票。优先队列**至少含有两种操作**的数据结构：**insert（插入）**，即将元素插入到优先队列中（入队）；以及**deleteMin（删除最小者）**，它的作用是找出、删除优先队列中的最小的元素（出队）。

### 优先队列的特性

- 优先队列的实现常选用**二叉堆**，**在数据结构中，优先队列一般也是指堆**。
- **堆的两个性质：**

1. **结构性**：**堆是一颗除底层外被完全填满的二叉树，底层的节点从左到右填入，**这样的树叫做**完全二叉树。**
2. **堆序性：**由于我们想很快找出最小元，则最小元应该在根上，**任意节点都小于它的后裔**，这就是**小顶堆（Min-Heap）**；如果是查找最大元，则最大元应该在根上，**任意节点都要大于它的后裔**，这就是**大顶堆(Max-heap)。**

![](http://exia.gz01.bdysite.com/uploads/big/4511f41546749db9fd96d693bd7058fd.jpg)

通过观察发现，**完全二叉树可以直接使用一个数组表示**而不需要使用其他数据结构。所以我们只需要传入一个size就可以构建优先队列的结构。

![完全二叉树的数组实现](http://exia.gz01.bdysite.com/uploads/big/87befae67ff2f1c4972ba6822dadce92.jpg)完全二叉树的数组实现

对于数组中的任意位置 i 的元素，其**左儿子**在位置 **2i** 上，则**右儿子**在 **2i+1** 上，**父节点**在 在 **i/2**（向下取整）上。通常从数组下标1开始存储，这样的好处在于很方便找到左右、及父节点。如果从0开始，左儿子在2i+1,右儿子在2i+2,父节点在(i-1)/2（向下取整）。

#### 堆序性：

我们这建立**最小堆，即对于每一个元素X，X的父亲中的关键字小于（或等于）X中的关键字，根节点除外（它没有父节点）。**

![](http://exia.gz01.bdysite.com/uploads/big/b00fcfe4ac299b3f5d84e614d41642f4.jpg)

### 优先队列的几种实现方案比较

时间复杂度比较:

|              |  入队   | 出队（拿出最大元素） |
| ------------ | :-----: | :------------------: |
| 普通线性结构 |  O(1)   |         O(n)         |
| 顺序线性结构 |  O(n)   |         O(1)         |
| 堆           | O(logn) |       O(logn)        |

#### 用数组实现一个大顶堆

```go
type MaxHeap struct {
	arr []int
}

func(heap *MaxHeap) Size() int {
	return len(heap.arr)
}

func(heap *MaxHeap) IsEmpty () bool {
	if len(heap.arr) == 0 {
		return true
	}
	return false
}
func(heap *MaxHeap) parent (index int) int {
	if index==0 {
		panic("0节点没有父亲节点")
	}
	return (index-1)/2
}

func(heap *MaxHeap) leftChild(index int) int {
	return index*2+1
}

func(heap *MaxHeap) rightChild(index int) int {
	return index*2+2
}
```

#### 向堆中添加元素和Sift UP(堆中元素上浮)

```go
// 添加元素，公有方法
func (heap *MaxHeap)Add(e int)  {
	heap.arr = append(heap.arr,e)
	heap.siftUp(heap.Size()-1)
}
// 元素上浮，私有方法
func (heap *MaxHeap) siftUp(size int) {
	// 如果size>0并且父亲节点比新插入的节点小，那么两个元素交换
	for size > 0 && heap.arr[heap.parent(size)] < heap.arr[size]{
		heap.arr[heap.parent(size)],heap.arr[size]=heap.arr[size],heap.arr[heap.parent(size)]
		size = heap.parent(size)
	}
}
```

#### 堆中元素出队和Sift Down（数据下沉）

```go
func(heap *MaxHeap) RemoveMax() int {
	var ret int = heap.FindMax()
	heap.arr[heap.Size()-1],heap.arr[heap.Size()-1] = heap.arr[heap.Size()-1],heap.arr[heap.Size()-1]
	heap.arr = heap.arr[:heap.Size()-1]
	heap.siftDown(0)
	
	return ret
}

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
			heap.arr[index],heap.arr[j]=heap.arr[j],heap.arr[index]
			index = j
		}
	}
}
```

#### 数组实现堆的测试用例

```go
import (
	"fmt"
	"math"
	"math/rand"
	"testing"
)

func TestMaxHeap(t *testing.T)  {
	n := 1000000
	ArrHeap := MaxHeap{
		arr:make([]int,0),
	}
	for i:=0; i < n; i++ {
		ArrHeap.Add(rand.Intn(math.MaxInt32))
	}

	// 生成一个一百万空间的数组把堆装回去，如果我们的算法无误，那么数组将是有序的
	test_arr := make([]int,1000000)
	for i:=0; i<n; i++ {
		test_arr[i] = ArrHeap.RemoveMax()
	}
	for i:=1;i<n;i++ {
		if test_arr[i-1] < test_arr[i] {
			panic("err!")
		}
	}
	fmt.Println("大顶堆运行成功！")
}
```

测试结果:

```
=== RUN   TestMaxHeap
大顶堆运行成功！
--- PASS: TestMaxHeap (0.08s)
PASS
```

#### Heapify和replace

replace：取出最大元素后，放入一个新元素

实现：可以先RemoveMax，再add，两次O(logn)的操作

实现2：可以直接将堆顶元素替换以后再sfitDown，一次O(logn)的操作。

```go
// 取出堆中最大的元素，并且替换成元素e
func (heap *MaxHeap) Replace(e int) int {
	ret := heap.FindMax()
	heap.arr[0] = e
	heap.siftDown(0)
	return ret
}
```

heapify：将一个任意数组整理成堆的形状

实现1：将n个元素逐个插入到一个空堆中，算法复杂度为O(nlogn)

实现2: 将叶子节点以外的节点依次进行siftDown操作

实现2的代码

```go
func (heap *MaxHeap) InitHeapWithArray (arr []int) *MaxHeap {
	heap.arr = make([]int, len(arr))
	for i:=0; i<len(arr); i++ {
		heap.arr[i] = arr[i]
	}
	for i:=heap.parent(len(arr)-1);i>=0;i-- {
		heap.siftDown(i)
	}
	return heap
}
```

测试用例:

```go
import (
	"fmt"
	"math"
	"math/rand"
	"testing"
	"time"
)

func TestMaxHeap(t *testing.T)  {
	n := 1000000
	ArrHeap := MaxHeap{
		arr:make([]int,0),
	}
	start := time.Now()
	for i:=0; i < n; i++ {
		ArrHeap.Add(rand.Intn(math.MaxInt32))
	}
	end := time.Now()
	// 生成一个一百万空间的数组把堆装回去，如果我们的算法无误，那么数组将是有序的

	test_arr := make([]int,1000000)
	for i:=0; i<n; i++ {
		test_arr[i] = ArrHeap.RemoveMax()
	}
	for i:=1;i<n;i++ {
		if test_arr[i-1] < test_arr[i] {
			panic("err!")
		}
	}

	fmt.Println("大顶堆运行成功！",end.Sub(start))

	HeapArr := &MaxHeap{}
	for i:=0; i<n; i++ {
		test_arr[i] = rand.Intn(math.MaxInt32)
	}
	start = time.Now()
	HeapArr.InitHeapWithArray(test_arr)
	end = time.Now()
	for i:=0; i<n; i++ {
		test_arr[i] = HeapArr.RemoveMax()
	}
	for i:=1;i<n;i++ {
		if test_arr[i-1] < test_arr[i] {
			panic("err!")
		}
	}

	fmt.Println("Heapify 运行成功！",end.Sub(start))
}
```

测试结果:

```
=== RUN   TestMaxHeap
大顶堆运行成功！ 69.0305ms
Heapify 运行成功！ 16.0674ms
--- PASS: TestMaxHeap (0.13s)
PASS
```

可以看到使用heapify创建堆的速度在一百万的数量级上明显是比把数组中的每个元素添加进去要快的多的。