package Linked

func MergeTwoList(head1, head2 *Node) *Node {
	if head2 == nil || head2.Next == nil {
		return head1
	}
	if head1 == nil || head1.Next == nil {
		return head2
	}

	var start *Node
	var end *Node
	start1 := head1.Next
	start2 := head2.Next

	if head1.Next.E > head2.Next.E {
		start = head2
		end = start2
		start2 = start2.Next
	} else {
		start = head1
		end = start1
		start1 = start1.Next
	}
	for start1 != nil && start2 != nil {
		if start1.E < start2.E {
			end.Next = start1
			end = start1
			start1 = start1.Next
		} else {
			end.Next = start2
			end = start2
			start2 = start2.Next
		}
	}

	if start1 == nil {
		end.Next = start2
	}
	if start2 == nil {
		end.Next = start1
	}
	return start
}
