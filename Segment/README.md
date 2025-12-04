## 线段树（Segment Tree）

### 定义

**线段树（segment tree）**，顾名思义， 是用来存放给定区间（segment, or interval）内对应信息的一种数据结构。与树状数组（binary indexed tree）相似，线段树也用来处理数组相应的**区间查询（range query）**和**元素更新（update）**操作。与树状数组不同的是，线段树不止可以适用于区间求和的查询，也可以进行**区间最大值，区间最小值（Range Minimum/Maximum Query problem）或者区间异或值**的查询。

对应于树状数组，线段树进行更新（update）的操作为O(logn)，进行区间查询（range query）的操作也为O(logn)。

注意：

1. 线段树不是完全二叉树
2. 线段树是平衡二叉树
3. 堆（完全二叉树）也是平衡二叉树
4. 二分搜索树不一定是平衡二叉树

### 为什么要使用线段树？

对于一类问题，我们关心的是一个线段（区间）比如：区间染色问题

> 有一面墙，长度为n，每次选择一段墙进行染色，m次操作后，我们可以在[i,j]区间内看到多少颜色？

另一类经典问题：区间查询

> 查询一个区间[i,j]的最大值，最小值，或者区间数字和等等。

实质：基于区间的统计查询，更具体的例子，查询在2019年注册用户中查消费最高的用户？消费最少的用户？

### 线段树的数组表示

线段树是平衡二叉树，依然可以用数组表示，把线段树看做一个满二叉树即可。

> 问题：如果区间有n个元素，数组表示需要有多少个节点？

对于满二叉树来说，第h层一共有2^h^ -1个节点(大约是2^h^),最后一层有2^(h-1)^ 个节点，最后一层的节点数大致等于前面所有层节点之和。

再回头看问题，可以得到如下结论：

如果n = 2^k^需要2n的空间

最坏的情况下如果n = 2^k^+1，那么我们需要4n的空间

那么我们需要4n的空间才能表示出一颗满二叉树

在这里我们不考虑添加元素，即区间固定，使用4n的空间即可。

#### 线段树的创建

```go
// 线段树的数据结构
type Tree struct {
	data []int
	tree []int
	// 融合器
	merger func(l,r int)int
}

func (this *Tree) GetSize () int {
	return len(this.data)
}
func (this *Tree) Get(index int) (res int,err error) {
	if index < 0 || index >= len(this.data) {
		err = errors.New("index out of range")
	}
	res = this.data[index]
	return
}

func (this *Tree) leftChild (k int) int {
	return 2*k+1
}

func (this *Tree) rightChild (k int) int {
	return 2*k+2
}
```

线段树的初始化方法

```go
func (this *Tree) Init(arr []int) {
	this.data = make([]int,len(arr))
	for i,v := range arr{
		this.data[i] = v
	}
	this.tree = make([]int,4*len(arr))
}
```

在建立线段树的时候，有一个问题，那就是我们拿到这个区间之后到底该干什么其实我们是不确定的，比如我们想给对应的区间求和，我们想取出区间最大值等等，那么这里利用了go语言中函数是一等公民的特性。通过传入一个函数来对数据做处理就可。

完善后的初始化方法:

```go
func (this *Tree) Init(arr []int,merger func(l,r int)int) {
	this.data = make([]int,len(arr))
	for i,v := range arr{
		this.data[i] = v
	}
	this.tree = make([]int,4*len(arr))
	this.merger = merger
	// 如果传入的数组大于0才创建二叉树，空数组没必要创建。
	if len(this.data) > 0{
		this.buildSegmentTree(0,0,len(this.data)-1)
	}
}
// 在treeIndex的位置创建表示区间[l……r]的线段树
func (this *Tree) buildSegmentTree(treeIndex,l,r int) {
	if l == r {
		this.tree[treeIndex] = this.data[l]
		return
	}
	leftChild,rightChild := this.leftChild(treeIndex),this.rightChild(treeIndex)
	mid := l + (r-l)/2

	// 递归调用时创建从[l……mid]和[mid+1……r]两个区间的线段树
	this.buildSegmentTree(leftChild,l,mid)
	this.buildSegmentTree(rightChild,mid+1,r)

	// 在这里tree[index]的数据确定应该由具体的业务逻辑决定的
	this.tree[treeIndex] = this.merger(this.tree[leftChild],this.tree[rightChild])
}
```

测试用例:

```go
func TestTree(t *testing.T) {
	tree := &Tree{}
	arr  := []int{-2,0,3,-5,2,-1}
	tree.Init(arr,add)
	fmt.Println(tree)
}
// 我们通过传入add函数求区间的和
func add (a,b int) int {
	return a+b
}
```

测试结果：

```go
=== RUN   TestTree
[-3,1,-4,-2,3,-3,-1,-2,0,0,0,-5,2,0,0,0,0,0,0,0,0,0,0,0,]
--- PASS: TestTree (0.00s)
PASS
```

#### 线段树的区间查询

```go
// 返回区间[l……R]的值
func (this *Tree) QueryLR (QueryL,QueryR int) (res int,err error) {
	if QueryL<0 || QueryL>=len(this.data) || QueryR<0 || QueryR >=len(this.data){
		err = errors.New("index out of range!")
	}
	res = this.query(0,0,len(this.data)-1,QueryL,QueryR)
	return
}

func (this *Tree) query(treeIndex,l,r,queryL,queryR int)int {
	if l == queryL && r == queryR {
		return this.tree[treeIndex]
	}
	leftChild,rightChild := this.leftChild(treeIndex),this.rightChild(treeIndex)
	mid := l + (r - l)/2
	// 如果用户关心的区间跟左孩子完全没有关系，那么只需要查找右子树
	if queryL >= mid+1 {
		return this.query(rightChild,mid+1,r,queryL,queryR)
        // 如果用户关心的区间跟右孩子完全没有关系，那么只需要查找左子树
	} else if queryR <= mid {
		return this.query(leftChild,l,mid,queryL,queryR)
	}
	LeftRes := this.query(leftChild,l,mid,queryL,mid)
	RightRes := this.query(rightChild,mid+1,r,mid+1,queryR)
	return this.merger(LeftRes,RightRes)
}
```

#### 线段树的数值更新

```go
func (this *Tree) Update (index,val int) {
	if index < 0 || index >= len(this.data) {
		panic("index out of range!")
	}
	this.data[index] = val
	this.update(0,0,len(this.data)-1,index,val)
}
// 递归更新区间，时间复杂度为O(logn)
func (this *Tree) update (treeIndex,l,r,index,e int) {
	if l == r {
		this.tree[treeIndex] = e
		return
	}
	leftChild,rightChild := this.leftChild(treeIndex),this.rightChild(treeIndex)
	mid := l + (r-l)/2
	if index >= mid+1 {
		this.update(rightChild,mid+1,r,index,e)
	} else {
		this.update(leftChild,l,mid,index,e)
	}
    // 重新用融合器融合一次
	this.tree[treeIndex] = this.merger(this.tree[leftChild],this.tree[rightChild])
}
```

#### [使用线段树解决LeetCode307号问题](https://leetcode-cn.com/problems/range-sum-query-mutable/submissions/)

```go
// 这里我们先把之前实现的线段树复制过来

type NumArray struct {
    tree *Tree
}

func add(a,b int) int {
    return a+b
}

func Constructor(nums []int) NumArray {
    nA := &NumArray{
        tree:&Tree{},
    }
    nA.tree.Init(nums,add)
    return *nA
}
func (this *NumArray) Update(i int, val int)  {
    this.tree.Update(i,val)
}

func (this *NumArray) SumRange(i int, j int) int {
    res,_ := this.tree.QueryLR(i,j)
    return res
}
```

结果：

```
执行用时 :52 ms, 在所有 Go 提交中击败了97.37%的用户
内存消耗 :9.2 MB, 在所有 Go 提交中击败了71.43%的用户
```
