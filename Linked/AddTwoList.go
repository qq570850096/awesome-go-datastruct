package Linked

func AddTwoList(head1,head2 *Node) *List {
	newHead := InitList()
	p1 := head1.Next
	p2 := head2.Next
	carry := 0
	for p1 != nil || p2 != nil {
		if p1 == nil {
			newHead.AddLast(p2.E+ + carry)
			p2 = p2.Next
		}
		if p2 == nil {
			newHead.AddLast(p1.E + carry)
			p1 = p1.Next
		}
		temp := (p1.E + p2.E + carry )%10
		newHead.AddLast(temp)
		carry = (p1.E + p2.E + carry )/10
		p1 = p1.Next
		p2 = p2.Next
	}
	return newHead
}
