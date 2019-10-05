package Segment

import (
	"errors"
	"fmt"
	"strings"
)

type Tree struct {
	data []int
	tree []int
	// 融合器
	merger func(l,r int)int
}

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
// 返回区间[l……R]的值
func (this *Tree) QueryLR (QueryL,QueryR int) (res int,err error) {
	if QueryL<0 || QueryL>=len(this.data) || QueryR<0 || QueryR >=len(this.data) || QueryL>QueryR{
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
	} else if queryR <= mid {
		return this.query(leftChild,l,mid,queryL,queryR)
	}
	LeftRes := this.query(leftChild,l,mid,queryL,mid)
	RightRes := this.query(rightChild,mid+1,r,mid+1,queryR)
	return this.merger(LeftRes,RightRes)
}
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
	this.tree[treeIndex] = this.merger(this.tree[leftChild],this.tree[rightChild])
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

func (this *Tree) String() string {
	var builder strings.Builder
	builder.WriteString("[")
	for i:=0; i<len(this.tree); i++ {
		fmt.Fprintf(&builder,"%d,",this.tree[i])
	}
	builder.WriteString("]")
	return builder.String()
}
