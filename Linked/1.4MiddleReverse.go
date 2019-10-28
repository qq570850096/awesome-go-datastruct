package Linked

func (this *List)findMidNode() *Node {
	if this.Head() == nil || this.Head().Next == nil {
		return nil
	}
	slow := this.Head()
	slowPre := this.Head()
	fast := this.Head()

	for fast != nil && fast.Next != nil {
		slowPre = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	slowPre.Next = nil
	return slow
}

func MidReverse(head *Node) *Node {
	if head == nil || head.Next == nil {
		return nil
	}

	var next *Node
	var pre *Node
	for head != nil {
		next = head.Next
		head.Next = pre
		pre = head
		head = next
	}
	return pre
}

func (this *List) Reorder ()  {
	if this.Head() == nil || this.Head() == nil {
		return
	}
	cur1 := this.Head().Next
	mid := this.findMidNode()
	cur2 := MidReverse(mid)
	var temp *Node
	// 合并链表
	for cur1.Next != nil {
		temp = cur1.Next
		cur1.Next = cur2
		cur1 = temp
		temp = cur2.Next
		cur2.Next = cur1
		cur2 = temp
	}
	cur1.Next = cur2
}

