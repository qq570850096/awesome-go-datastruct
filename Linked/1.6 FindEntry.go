package Linked

func (this *List) FindLoop() *Node {
	// 如果链表是空的直接返回就可
	if this.Head() == nil || this.Head().Next == nil {
		return nil
	}
	// 让快慢指针都指向头结点的下一个节点
	fast,slow := this.Head().Next,this.Head().Next
	for fast!=nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			return fast
		}
	}
	return nil
}

func (this *List)FindLoopEntryNode(meet *Node) *Node {
	entry := this.Head().Next

	for entry != meet {
		entry = entry.Next
		meet = meet.Next
	}
	return entry
}
