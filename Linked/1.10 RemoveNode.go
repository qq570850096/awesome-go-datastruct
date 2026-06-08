package Linked

func (this *List) RemoveNode(node *Node) bool {

	if this.Head() == nil || this.Head().Next == nil || node.Next == nil {
		return false
	}
	E := node.Next.E
	node.E = E
	node.Next = node.Next.Next
	return true
}
