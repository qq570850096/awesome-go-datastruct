package Linked

func (this *List) ReverseTwo() {
	if this.Head() == nil || this.Head().Next == nil {
		return
	}

	cur := this.Head().Next
	pre := this.Head()
	var next *Node
	for cur != nil && cur.Next != nil {

		next = cur.Next.Next

		pre.Next = cur.Next

		cur.Next.Next = cur

		cur.Next = next

		pre = cur
		cur = next
	}
}
