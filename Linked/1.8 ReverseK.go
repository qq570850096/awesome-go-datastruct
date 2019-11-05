package Linked

// 对单向链表的前k个节点进行翻转
// 对没有虚拟头结点的链表进行翻转
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

func (this *List) ReverseK(k int)  {
	if this.Head() == nil || this.Head().Next == nil {
		return
	}
	pre := this.Head()
	begin := this.Head().Next
	var end *Node
	var pNext *Node
	for begin != nil {
		end = begin

		for i:=1;i<k;i++ {
			if end.Next != nil {
				end = end.Next
			}else {
				return
			}
		}

		pNext = end.Next
		end.Next = nil
		// 将翻转后的链表接到头结点上
		pre.Next = reverse(this.Head())
		begin.Next = pNext
		pre = begin
		begin = pNext
	}
}
