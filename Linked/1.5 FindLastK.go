package Linked

func (this *List) FindLastK(k int) *Node {
	if this.Head() == nil || this.Head().Next == nil {
		return nil
	}
	fast,slow := this.Head().Next,this.Head().Next
	var i int
	for i=0;i<k && fast != nil;i++ {
		// 这里让快指针先走
		fast = fast.Next
	}
	// 说明k比链表长度长了，直接返回空节点
	if i < k {
		return nil
	}
	// 让快慢指针一起移动，当快指针到链表尾部的时候，慢指针就在链表的LastK位置
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	return slow
}
