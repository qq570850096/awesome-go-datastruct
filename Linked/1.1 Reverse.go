package Linked

import "fmt"

func (this *List) Reverse() {
	var pre *Node
	var cur *Node
	next := this.Head().Next

	for next != nil {
		cur = next.Next
		next.Next = pre
		pre = next
		next = cur
	}
	this.Head().Next = pre
}

func recursiveReverse(node *Node) *Node {
	if node == nil || node.Next == nil {
		return node
	}
	newHead := recursiveReverse(node.Next)
	node.Next.Next = node
	node.Next = nil
	return newHead
}

func (this *List) RecursiveReverse() {
	firstNode := this.Head().Next
	newHead := recursiveReverse(firstNode)
	this.Head().Next = newHead
}

func (this *List) InsertReverse() {
	if this.Head() == nil || this.Head().Next == nil {
		return
	}
	var cur *Node
	var next *Node
	cur = this.Head().Next.Next

	this.Head().Next.Next = nil

	for cur != nil {
		next = cur.Next
		cur.Next = this.Head().Next
		this.Head().Next = cur
		cur = next
	}
}

func (this List) ReversePrint(node *Node) {
	if node == nil {
		return
	}
	this.ReversePrint(node.Next)
	fmt.Printf("%d ", node.E)
}
