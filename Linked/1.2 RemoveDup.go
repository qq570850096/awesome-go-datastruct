package Linked
// 顺序式删除重复节点
func (this *List) RemoveDup () {

	if this.Head() == nil || this.Head().Next == nil {
		return
	}
	// 外层循环，指向链表的第一个节点
	outerCur := this.Head().Next
	// 内层循环innerPre 和 innerCur
	var innerPre,innerCur *Node

	for ;outerCur != nil ; outerCur = outerCur.Next {
		for innerPre,innerCur = outerCur,outerCur.Next; innerCur != nil ; {
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

func (this *List) RemoveDupRecursion (){
	if this.Head() == nil {
		return
	}
	this.Head().Next = removeDupRecursionChild(this.Head().Next)
}
// 递归式删除重复节点
func removeDupRecursionChild (node *Node) *Node {
	if node == nil || node.Next == nil {
		return node
	}
	var pointer *Node
	cur := node
	// 对以node.Next为首的子链表删除重复的节点
	node.Next = removeDupRecursionChild(node.Next)
	// 找出以node.Next为首的子链表中与node结点相同的结点并删除
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

// 用空间换时间
func (this *List) RemoveDupWithMap () {
	if this.Head() == nil || this.Head().Next == nil {
		return
	}
	searchMap := make(map[int]*Node)
	pre := this.Head()
	cur := this.Head().Next
	for cur != nil {
		// 如果在哈希表中找到了这个数值，那就删除掉cur
		if _,ok := searchMap[cur.E]; ok {
			pre.Next = cur.Next
			cur = cur.Next
		} else {
			searchMap[cur.E] = cur
			cur = cur.Next
			pre = pre.Next
		}
	}
}
