package Linked
// 第一种方法，首尾相接法，把两个链表首尾相接，如果有交叉，必成一个环，只需要用1.6做的寻找入口节点的方式就能找到结果
// 时间复杂度O(n1+n2)，空间复杂度O(n1)
func (this *List)CheckIntersect(head1,head2 *Node) *Node {
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
