package Linked

func (this *List) FindLastK(k int) *Node {
	if this.Head() == nil || this.Head().Next == nil {
		return nil
	}
	fast, slow := this.Head().Next, this.Head().Next
	var i int
	for i = 0; i < k && fast != nil; i++ {

		fast = fast.Next
	}

	if i < k {
		return nil
	}

	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	return slow
}
