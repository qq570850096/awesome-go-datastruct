package BinarySearch

import (
	"fmt"
	"strings"
)

type Node struct {
	E int
	Left *Node
	Right *Node
}

// 约定在此样例代码中我们的二分搜索树中没有重复元素
// 如果想包涵重复元素的话，只需要以下定义：
// 左子树小于等于此节点，或右子树大于等于节点
type Tree struct {
	root *Node
	size int
}

func (t Tree) Size() int {
	return t.size
}

func (t Tree) Root() *Node {
	return t.root
}
// 判断二叉树是否为空
func (this *Tree) IsEmpty() bool {
	if this.size == 0 {
		return true
	}
	return false
}

func InitNode(E int) *Node {
	return &Node{
		E:E,
		Left:nil,
		Right:nil,
	}
}

func (this *Tree) AddE(e int) {
	this.root = this.add(this.root,e)
}
// 向以node为根的二分搜索树中插入元素E，递归算法
func (this *Tree) add(node *Node,e int) *Node{

	// 不管是递归还是回溯，首先我们都应该先写出递归的结束条件是什么
	if node == nil {
		this.size++
		return InitNode(e)
	}

	if e > node.E {
		node.Right = this.add(node.Right, e)
	} else if e < node.E {
		node.Left = this.add(node.Left, e)
	}
	return node
}

// 查找二分搜索中是否含有元素E
func (this *Tree)Contains(e int) bool {
	return this.contains(this.root, e)
}
// 递归的方式查找元素是否存在
func (this *Tree)contains (node *Node, e int) bool {
	if this.root == nil {
		return false
	}
	if e == node.E {
		return true
	} else if e > node.E {
		this.contains(node.Right,e)
	} else {
		this.contains(node.Left, e)
	}
	panic("运行出错！")
}

// 遍历算法
// 1.前序遍历
func(this *Tree)PreOrder(){
	PreOrder(this.root)
	fmt.Println()
}
func PreOrder(node *Node)  {
	if node == nil {
		return
	}
	fmt.Printf("%d ",node.E)
	PreOrder(node.Left)
	PreOrder(node.Right)
}
// 非递归的前序遍历
func (this *Tree) PreOrderNR() {
	stack := make([]*Node, 0)
	stack = append(stack, this.root)
	for len(stack) > 0 {
		curNode := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		fmt.Printf("%d ",curNode.E)
		if curNode.Right != nil {
			stack = append(stack, curNode.Right)
		}
		if curNode.Left != nil {
			stack = append(stack, curNode.Left)
		}
	}
	fmt.Println()
}
// 2.中序遍历
func(this *Tree)MidOrder(){
	MidOrder(this.root)
}
func MidOrder(node *Node)  {
	if node == nil {
		return
	}

	MidOrder(node.Left)
	fmt.Printf("%d ",node.E)
	MidOrder(node.Right)
}
// 3.后序遍历
func (this *Tree) BackOrder(){
	BackOrder(this.root)
}
func BackOrder(node *Node)  {
	if node == nil {
		return
	}
	BackOrder(node.Left)
	BackOrder(node.Right)
	fmt.Printf("%d ",node.E)
}
// 二分搜索树的层序遍历
func (this *Tree)LevelOrder(){
	queue := make([]*Node,0)
	queue = append(queue, this.root)
	for len(queue) > 0 {
		curNode := queue[0]
		queue = queue[1:]
		fmt.Printf("%d ",curNode.E)
		if curNode.Left != nil {
			queue = append(queue,curNode.Left)
		}
		if curNode.Right != nil {
			queue = append(queue,curNode.Right)
		}
	}
}
// 二分搜索树中搜索最小值
func (this *Tree) FindMin() int{
	if this.IsEmpty() {
		panic("二叉树为空，无法删除任何节点")
	}
	return minimum(this.root).E
}
func minimum(node *Node) *Node {
	if node.Left == nil {
		return node
	}
	return minimum(node.Left)
}
// 二分搜索树中搜索最大值
func (this *Tree) FindMax() int{
	if this.IsEmpty() {
		panic("二叉树为空，无法删除任何节点")
	}
	return maximum(this.root).E
}
func maximum(node *Node) *Node {
	if node.Right == nil {
		return node
	}
	return maximum(node.Right)
}
// 从二分搜索树种删除最小值
func (this *Tree) DelMin() int {
	var ret int = this.FindMin()
	this.root = this.rmMin(this.root)
	return ret
}
// 删除掉以node为根的二分搜索树的最小节点
// 返回删除节点后新的二分搜索树的根
func (this *Tree) rmMin(node *Node) *Node {
	if node.Left == nil {
		nodeRight := node.Right
		node.Right = nil
		this.size--
		return nodeRight
	}
	node.Left = this.rmMin(node.Left)
	return node
}
// 从二分搜索树种删除最大值
func (this *Tree) DelMax() int {
	var ret int = this.FindMax()
	this.root = this.rmMax(this.root)
	return ret
}
// 删除掉以node为根的二分搜索树的最小节点
// 返回删除节点后新的二分搜索树的根
func (this *Tree) rmMax(node *Node) *Node {
	if node.Right == nil {
		nodeLeft := node.Left
		node.Left = nil
		this.size--
		return nodeLeft
	}
	node.Right = this.rmMax(node.Right)
	return node
}
// 在二分搜索树中删除值为e的方法
func (this *Tree) Remove (e int){
	this.root = this.remove(this.root,e)
}
func (this *Tree) remove(node *Node,e int) *Node {
	if node == nil {
		return nil
	}
	if e > node.E {
		node.Right = this.remove(node.Right,e)
		return node
	} else if e < node.E {
		node.Left = this.remove(node.Left,e)
		return node
	} else {
		// 如果左子树为空的时候
		if node.Left == nil {
			nodeRight := node.Right
			node.Right = nil
			this.size--
			return nodeRight
		}
		// 如果右子树为空
		if node.Right == nil {
			nodeLeft := node.Left
			node.Left = nil
			this.size--
			return nodeLeft
		}
		// 如果左右子树都不为空，那么我们需要找到node的后继
		// 就是所有比node值大的节点中值最小的那个，显然它在node的右子树中
		nodeNext := minimum(node.Right)
		nodeNext.Right = this.rmMin(node.Right)
		nodeNext.Left = node.Left
		node.Left = nil
		node.Right = nil
		return nodeNext
	}
}
// 重构二叉树的打印方法
func (this *Tree) String() string{
	var (
		builder strings.Builder
	)
	generateBSTString(this.root,0,&builder)
	return builder.String()
}
func generateBSTString(node *Node,depth int,builder *strings.Builder){
	if node == nil {
		fmt.Fprintln(builder,generateDepthString(depth) + "null")
		return
	}
	fmt.Fprintln(builder,generateDepthString(depth),node.E)
	generateBSTString(node.Left,depth+1,builder)
	generateBSTString(node.Right,depth+1,builder)
}
func generateDepthString(depth int) string{
	var builder strings.Builder
	for i:=0;i<depth;i++ {
		fmt.Fprintf(&builder,"--")
	}
	return builder.String()
}
