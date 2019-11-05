package Linked

// 对单链表两两节点翻转的算法

func (this *List) ReverseTwo()  {
	if this.Head() == nil || this.Head().Next==nil {
		return
	}

	cur := this.Head().Next
	pre := this.Head()
	var next *Node
	for cur!=nil && cur.Next != nil {
		// 记录逆序后的下一个节点
		next = cur.Next.Next
		// pre的下一个节点指向之后第二个节点
		pre.Next = cur.Next
		// cur的下一个节点指向cur
		cur.Next.Next = cur
		// 这时逆序已经完成了，我们把cur的下个节点指向next
		cur.Next = next
		// pre和cur都后移一位
		pre = cur
		cur = next
	}
}
