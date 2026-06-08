package Linked

func (this *List) FindLoop() *Node {

	if this.Head() == nil || this.Head().Next == nil {
		return nil
	}

	fast, slow := this.Head().Next, this.Head().Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			return fast
		}
	}
	return nil
}

func (this *List) FindLoopEntryNode(meet *Node) *Node {
	entry := this.Head().Next

	for entry != meet {
		entry = entry.Next
		meet = meet.Next
	}
	return entry
}
