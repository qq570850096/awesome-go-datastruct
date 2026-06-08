package BinarySearch

import (
	"fmt"
	"strings"
)

type Node struct {
	E     int
	Left  *Node
	Right *Node
}

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

func (this *Tree) IsEmpty() bool {
	if this.size == 0 {
		return true
	}
	return false
}

func InitNode(E int) *Node {
	return &Node{
		E:     E,
		Left:  nil,
		Right: nil,
	}
}

func (this *Tree) AddE(e int) {
	this.root = this.add(this.root, e)
}

func (this *Tree) add(node *Node, e int) *Node {

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

func (this *Tree) Contains(e int) bool {
	return this.contains(this.root, e)
}

func (this *Tree) contains(node *Node, e int) bool {
	if node == nil {
		return false
	}
	if e == node.E {
		return true
	} else if e > node.E {
		return this.contains(node.Right, e)
	} else {
		return this.contains(node.Left, e)
	}
}

func (this *Tree) PreOrder() {
	PreOrder(this.root)
	fmt.Println()
}
func PreOrder(node *Node) {
	if node == nil {
		return
	}
	fmt.Printf("%d ", node.E)
	PreOrder(node.Left)
	PreOrder(node.Right)
}

func (this *Tree) PreOrderNR() {
	stack := make([]*Node, 0)
	stack = append(stack, this.root)
	for len(stack) > 0 {
		curNode := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		fmt.Printf("%d ", curNode.E)
		if curNode.Right != nil {
			stack = append(stack, curNode.Right)
		}
		if curNode.Left != nil {
			stack = append(stack, curNode.Left)
		}
	}
	fmt.Println()
}

func (this *Tree) MidOrder() {
	MidOrder(this.root)
}
func MidOrder(node *Node) {
	if node == nil {
		return
	}

	MidOrder(node.Left)
	fmt.Printf("%d ", node.E)
	MidOrder(node.Right)
}

func (this *Tree) BackOrder() {
	BackOrder(this.root)
}
func BackOrder(node *Node) {
	if node == nil {
		return
	}
	BackOrder(node.Left)
	BackOrder(node.Right)
	fmt.Printf("%d ", node.E)
}

func (this *Tree) LevelOrder() {
	queue := make([]*Node, 0)
	queue = append(queue, this.root)
	for len(queue) > 0 {
		curNode := queue[0]
		queue = queue[1:]
		fmt.Printf("%d ", curNode.E)
		if curNode.Left != nil {
			queue = append(queue, curNode.Left)
		}
		if curNode.Right != nil {
			queue = append(queue, curNode.Right)
		}
	}
}

func (this *Tree) FindMin() int {
	if this.IsEmpty() {
		panic("empty tree: cannot remove nodes")
	}
	return minimum(this.root).E
}
func minimum(node *Node) *Node {
	if node.Left == nil {
		return node
	}
	return minimum(node.Left)
}

func (this *Tree) FindMax() int {
	if this.IsEmpty() {
		panic("empty tree: cannot remove nodes")
	}
	return maximum(this.root).E
}
func maximum(node *Node) *Node {
	if node.Right == nil {
		return node
	}
	return maximum(node.Right)
}

func (this *Tree) DelMin() int {
	var ret int = this.FindMin()
	this.root = this.rmMin(this.root)
	return ret
}

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

func (this *Tree) DelMax() int {
	var ret int = this.FindMax()
	this.root = this.rmMax(this.root)
	return ret
}

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

func (this *Tree) Remove(e int) {
	this.root = this.remove(this.root, e)
}
func (this *Tree) remove(node *Node, e int) *Node {
	if node == nil {
		return nil
	}
	if e > node.E {
		node.Right = this.remove(node.Right, e)
		return node
	} else if e < node.E {
		node.Left = this.remove(node.Left, e)
		return node
	} else {

		if node.Left == nil {
			nodeRight := node.Right
			node.Right = nil
			this.size--
			return nodeRight
		}

		if node.Right == nil {
			nodeLeft := node.Left
			node.Left = nil
			this.size--
			return nodeLeft
		}

		nodeNext := minimum(node.Right)
		nodeNext.Right = this.rmMin(node.Right)
		nodeNext.Left = node.Left
		node.Left = nil
		node.Right = nil
		return nodeNext
	}
}

func (this *Tree) String() string {
	var (
		builder strings.Builder
	)
	generateBSTString(this.root, 0, &builder)
	return builder.String()
}
func generateBSTString(node *Node, depth int, builder *strings.Builder) {
	if node == nil {
		fmt.Fprintln(builder, generateDepthString(depth)+"null")
		return
	}
	fmt.Fprintln(builder, generateDepthString(depth), node.E)
	generateBSTString(node.Left, depth+1, builder)
	generateBSTString(node.Right, depth+1, builder)
}
func generateDepthString(depth int) string {
	var builder strings.Builder
	for i := 0; i < depth; i++ {
		fmt.Fprintf(&builder, "--")
	}
	return builder.String()
}
