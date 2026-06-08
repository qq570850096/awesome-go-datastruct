package Linked

func (this *List) CheckIntersect(head1, head2 *Node) *Node {
	if head2 == nil || head2.Next == nil || head1 == nil || head1.Next == nil || head1 == head2 {
		return nil
	}
	findTail := func(node *Node) *Node {
		cur := node
		for cur.Next != nil {
			cur = cur.Next
		}
		return cur
	}
	tail := findTail(head1)
	tail.Next = head2.Next
	meet := this.FindLoop()
	entry := this.FindLoopEntryNode(meet)
	return entry
}
