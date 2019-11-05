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

//方法二:尾结点法
//主要思路:如果两个链表相交，那么两个链表从相交点到链表结束都是相同的结点，
//必然是Y字形,所以，判断两个链表的最后一个结点是不是相同即可。
//即先遍历-个链表,直到尾部,再遍历另外一个链表,
//如果也可以走到同样的结尾点，则两个链表相交,这时记下两个链表的长度n1、n2,
//再遍历一次，长链表结点先出发前进|n1-n2|步,之后两个链表同时前进，、
//每次一步,相遇的第一点即为两个链表相交的第一个点。


