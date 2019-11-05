package Linked

// 给定一个链表中一个节点的指针，让你删除这个节点
func (this *List)RemoveNode(node *Node) bool {
	// 分析一下，我们如果想要删除一个节点，那么是需要找到这个节点的前驱的
	// 那么在这一题中，如果给的是尾节点，那么我们是没有办法删除的，但是如果不是尾节点的话
	// 我们可以采用这种取巧的方式删除这个节点
	// 1. 将该节点的后继节点的值赋值给它
	// 2. 将该节点的后继节点删除

	// 这里判断如果是空链表或者传入的是尾节点，那么我们就直接返回false
	if this.Head() == nil || this.Head().Next == nil||node.Next == nil{
		return false
	}
	E := node.Next.E
	node.E = E
	node.Next = node.Next.Next
	return true
}
