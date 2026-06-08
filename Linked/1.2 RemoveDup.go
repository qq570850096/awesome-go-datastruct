package Linked

func (this *List) RemoveDup() {

	if this.Head() == nil || this.Head().Next == nil {
		return
	}

	outerCur := this.Head().Next

	var innerPre, innerCur *Node

	for ; outerCur != nil; outerCur = outerCur.Next {
		for innerPre, innerCur = outerCur, outerCur.Next; innerCur != nil; {
			if innerPre.E == innerCur.E {
				innerPre.Next = innerCur.Next
				innerCur = innerCur.Next
			} else {
				innerPre = innerCur
				innerCur = innerCur.Next
			}
		}
	}
}

func (this *List) RemoveDupRecursion() {
	if this.Head() == nil {
		return
	}
	this.Head().Next = removeDupRecursionChild(this.Head().Next)
}

func removeDupRecursionChild(node *Node) *Node {
	if node == nil || node.Next == nil {
		return node
	}
	var pointer *Node
	cur := node

	node.Next = removeDupRecursionChild(node.Next)

	pointer = node.Next
	for pointer != nil {
		if node.E == pointer.E {
			cur.Next = pointer.Next
			pointer = pointer.Next
		} else {
			pointer = pointer.Next
			cur = cur.Next
		}
	}
	return node
}

func (this *List) RemoveDupWithMap() {
	if this.Head() == nil || this.Head().Next == nil {
		return
	}
	searchMap := make(map[int]*Node)
	pre := this.Head()
	cur := this.Head().Next
	for cur != nil {

		if _, ok := searchMap[cur.E]; ok {
			pre.Next = cur.Next
			cur = cur.Next
		} else {
			searchMap[cur.E] = cur
			cur = cur.Next
			pre = pre.Next
		}
	}
}
