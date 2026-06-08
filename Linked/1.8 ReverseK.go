package Linked

func reverse(head *Node) *Node {
	if head == nil || head.Next == nil {
		return nil
	}
	var pre *Node
	var next *Node
	for head != nil {
		next = head.Next
		head.Next = pre
		pre = head
		head = next
	}
	return pre
}

func (this *List) ReverseK(k int) {
	if this.Head() == nil || this.Head().Next == nil {
		return
	}
	pre := this.Head()
	begin := this.Head().Next
	var end *Node
	var pNext *Node
	for begin != nil {
		end = begin

		for i := 1; i < k; i++ {
			if end.Next != nil {
				end = end.Next
			} else {
				return
			}
		}

		pNext = end.Next
		end.Next = nil

		pre.Next = reverse(this.Head())
		begin.Next = pNext
		pre = begin
		begin = pNext
	}
}
