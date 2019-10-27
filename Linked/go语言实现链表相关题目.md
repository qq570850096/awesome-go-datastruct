## go语言实现链表相关题目

[TOC]

### 1.1 链表的逆序

#### 题目描述:
给定一一个带头结点的单链表,请将其逆序。即如果单链表原来为head->1->2->3->4->5->6->7.则逆序后变为head->7->6->5->4->3->2->1。
#### 分析与解答:
由于单链表与数组不同，单链表中每个结点的地址都存储在其前驱结点
的指针域中，因此,对单链表中任何一个结点的访问只能从链表的头指针开始进行遍历。在对链表的操作过程中，需要特别注意在修改结点指针域的时候,记录下后继结点的地址,否则会秩后继结点。
#### 方法一:就地逆序
主要思路为:在遍历链表的时候，修改当前结点指针域的指向,让其指
向它的前驱结点。为此需要用一个指针变量来保存前驱结点的地址。此外，为了在调整当前结点指针域的指向后还能找到后继结点，还需要另外-一个指针变量来保存后继结点的地址,在所有的结点都被保存好以后就可以直接完成指针的逆序了。除此之外,还需要特别注意对链表首尾结点的特殊处理。

具体实现方式如下图所示。

![](http://exia.gz01.bdysite.com/uploads/big/2f6d38bdaa3f719d61492070d836e107.png)

在上图中，假设当前已经遍历到cur结点，由于它所有的前驱结点都已
经完成了逆序操作,因此,只需要使cur.next= pre即可完成逆序操作,在此
之前为了能够记录当前结点的后继结点的地址,需要用一个额外的指针next
来保存后继结点的信息，通过上图(1) ~ (4)四步把实线的指针调整为虚线
的指针就可以完成当前结点的逆序;当前结点完成逆序后，通过向后移动指
针来对后续的结点用同样的方法进行逆序操作。实现代码如下:

```go
// 文档中的数据结构均使用我自己实现的golang数据结构
func (this *List)Reverse()  {
	var pre *Node
	var cur *Node
	next := this.Head().Next

	for next != nil {
		cur = next.Next
		next.Next = pre
		pre = next
		next = cur
	}
	this.Head().Next = pre
}

// 测试代码
package Linked

import (
	"testing"
)

func TestList(t *testing.T) {
	list := InitList()
	for i:=0;i<5 ; i++  {
		list.AddFirst(i)
	}
	t.Log(list.Contains(4))
	t.Log(list.Get(2))
	list.AddIndex(2,666)
	t.Log(list)
	list.Reverse()
	t.Log(list)
}

```

测试结果:

```
--- PASS: TestList (0.00s)
    List_test.go:12: true
    List_test.go:13: 2
    List_test.go:15: 4 -> 3 -> 666 -> 2 -> 1 -> 0 -> NULL
    List_test.go:17: 0 -> 1 -> 2 -> 666 -> 3 -> 4 -> NULL
PASS
```
---

##### 算法性能分析:

以上这种方法只需要对链表进行一一次遍历，因此,时间复杂度为0
(n),其中，n为链表的长度。但是需要常数个额外的变量来保存当前结点的
前驱结点与后继结点，因此，空间复杂度为0(1)。

#### 方法二:递归法

假定原链表为1->2->3->4->5->6->7.递归法的主要思路为:先逆序除第一个结点以外的子链表(将1->2->3->4->5->6->7 变为1->7->6->5->4->3->2),接着把结点1添加到逆序的子链表的后面(1->2->3->4->5->6->7变为7->6->5->4->3->2->1)。同理,在逆序链表2->3->4->5->6->7时，也是先逆序子链表3->4->5->6->7 (逆序为2->7->6->5->4->3),接着实现链表的整体逆序(2->7->6->5->4->3 转换为7->6->5->4->3->2)。实现代码如下:

```go
func recursiveReverse(node *Node) *Node {
	if node == nil || node.Next == nil {
		return node
	}
	newHead := recursiveReverse(node.Next)
	node.Next.Next = node
	node.Next = nil
	return newHead
}

func (this *List) RecursiveReverse() {
	firstNode := this.Head().Next
	newHead := recursiveReverse(firstNode)
	this.Head().Next = newHead
}
```

***

##### 算法性能分析:
由于递归法也只需要对链表进行一次遍历，因此，算法的时间复杂度也
为0 (n),其中, n为链表的长度。递归法的主要优点是:思路比较直观,容
易理解，而且也不需要保存前驱结点的地址;缺点是:算法实现的难度较
大，此外，于递归法需要不断地调用自己，需要额外的压栈与弹栈操作,因此，与方法一相比性能会有所下降。

#### 方法三:插入法
插入法的主要思路为:从链表的第2 - 2个结点开始,把遍历到的结点插入
到头结点的后面，直到遍历结束。假定原链表为node->1->2->3->4->5->6->7,在遍历到2的时候,将其插入到头结点后，链表变为node->2->1->3->4->5->6->7.同理将后序遍历到的所有结点都插入到头结点head后,就可以实现链表的逆序。实现代码如下:

```go
func (this *List) InsertReverse ()  {
	if this.Head() == nil || this.Head().Next == nil{
		return
	}
	var cur *Node //当前节点
	var next *Node //后继节点
	cur = this.Head().Next.Next
	// 设置链表第一个节点为尾节点
	this.Head().Next.Next = nil
	// 把遍历到的节点插入到头结点的后面
	for cur != nil {
		next = cur.Next
		cur.Next = this.Head().Next
		this.Head().Next = cur
		cur = next
	}
}
```

##### 算法性能分析:

以上这种方法也只需要对单链表进行一次遍历,因此，时间复杂度为0
(n),其中，n为链表的长度。与方法一相比， 这种方法不需要保存前驱结点的地址，与方法二相比，这种方法不需要递归地调用，效率更高。

#### 引申

##### 方法一:就地逆序+顺序输出

首先对链表进行逆序,然后顺序输出逆序后的链表。这种方法的缺点是
改变了链表原来的结构。

##### 方法二:逆序+顺序输出

申请新的存储空间，对链表进行逆序,然后顺序输出逆序后的链表。逆
序的主要思路为:每当遍历到-一个结点的时候，申请一块新的存储空间来存
储这个结点的数据域，同时把新结点插入到新链表的头结点后。这种方法的
缺点是需要申请额外的存储空间。

##### 方法三:递归输出

递归输出的主要思路为:先输出除当前结点外的后继子链表,然后输出
当前结点，假如链表为: 1->2->3->4->5->6->7.那么先输出2->3->4->5->6->7,再输出1。同理，对于链表2->3->4->5->6->7.也是先输出3->4->5->6->7,接着输出2,直到遍历到链表的最后-个结点7的时候会
输出结点7,然后递归地输出6, 5, ... 1。实现代码如下: 

```go
func (this List)ReversePrint(node *Node)  {
	if node == nil {
		return
	}
	this.ReversePrint(node.Next)
	fmt.Printf("%d ",node.E)
}
```

##### 算法性能分析

这种方法显然只需要对链表进行一次遍历，因此，时间复杂度为O(n)，n为链表的长度。


### 1.2 从无序链表中移除重复项

难度系数:★★★
被考查系数:★★★★

#### 题目描述:

给定一个没有排序的链表,去掉其重复项,并保留原顺序,例如链表1->3->1->5->5->7,去掉重复项后变为1->3->5->7.
分析与解答:

#### 方法一:顺序删除

主要思路为:通过双重循环直接在链表上进行删除操作。外层循环用一个指针从第-一个结点开始遍历整个链表,然后内层循环用另外一个指针遍历其余结点,将与外层循环遍历到的指针所指结点的数据域相同的结点删除。
如下图所示:

![](http://www.liuanqihappybirthday.top/uploads/big/7c3f40b9821ac4498f52e403f95bd1e1.png)

假设外层循环从outerCur开始遍历，当内层循环指针innerCur遍历到上图实线所示的位置(outerCur.data= = innerCur.data)时，此时需要把
innerCur指向的结点删除。具体步骤如下:

1. 用tmp记录待删除的结点的地址。

2. 为了能够在删除tmp结点后继续遍历链表中其余的结点，使
   innerCur指向它的后继结点: innerCur= innerCur.next.

3. 从链表中删除tmp结点。


实现代码如下:

```go
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
```

测试代码和结果

```go
func TestList(t *testing.T) {
	list := InitList()
	for i:=0;i<5 ; i++  {
		list.AddFirst(i)
		list.AddFirst(i)
	}
	t.Log(list)
	list.RemoveDup()
	t.Log(list)	
}

// 测试结果

=== RUN   TestList
--- PASS: TestList (0.00s)
    List_test.go:24: 4 -> 4 -> 3 -> 3 -> 2 -> 2 -> 1 -> 1 -> 0 -> 0 -> NULL
    List_test.go:26: 4 -> 3 -> 2 -> 1 -> 0 -> NULL
PASS
```

##### 性能分析:

由于这种方法采用双重循环对链表进行遍历，因此,时间复杂度为O(n^2^),其中，N为链表的长度,在遍历链表的过程中,使用了常量个额外的
指针变量来保存当前遍历的结点、前驱结点和被删除的结点，因此，空间复杂度为0(1)。

#### 方法二:递归法

主要思路为:对于结点cur,首先递归地删除以cur.next为首的子链表
中重复的结点，接着从以cur.next为首的子链表中找出与cur有着相同数据域的结点并删除，实现代码如下:

```go
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
```

##### 算法性能分析:

这种方法与方法一类似, 从本质上而言,由于这种方法需要对链表进行
双重遍历，因此，时间复杂度为O(n2),其中, N为链表的长度。由于递归法会增加许多额外的函数调用,因此,从理论上讲,该方法效率比方法一低。

#### 方法三:空间换时间

通常情况下，为了降低时间复杂度,往往在条件允许的情况下，通过使
用辅助空间实现。具体而言,主要思路为:

1. 建立一个HashSet, HashSet中的内容为已经遍历过的结点内容,
   并将其初始化为空。
2. 从头开始遍历链表中的所以结点,存在以下两种可能性:
   1. 如果结点内容已经在HashSet中，则删除此结点,继续向后遍历。
   2. 如果结点内容不在HashSet中，则保留此结点,将词结点内容添加
      到HashSet中,继续向后遍历。

代码实现如下:

```go
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
```

##### 算法性能分析:

在最坏的情况下，链表没有重复的元素，我们就要申请n个空间，也就是说空间复杂度为O(n),但是链表我们只需要扫描一次即可，时间复杂度是O(n)，比n^2^的性能要好。

##### 引申:如何从有序链表中移除重复项。

分析与解答:

上述介绍的方法也适用于链表有序的情况，但是由于以上方法没有充分
利用到链表有序这个条件，因此，算法的性能肯定不是最优的。本题中，由于链表具有有序性,因此,不需要对链表进行两次遍历。所以,有如下思路:用cur指向链表第一个结点，此时需要分为以下两种情况讨论:

* 如果cur.data= =cur.next.data,那么删除cur.next 结点。
* 如果cur.data! =cur.next.data,那么cur=cur.next,继续遍历其
  余结点。
  
### 1.3 计算两个单链表所代表的的代数之和



### 1.4 对链表进行重新排序



### 1.5 找出单链表中倒数第k个元素



### 1.6 检测一个较大的单链表是否有环



### 1.7 把链表相邻元素翻转



### 1.8 把链表以k个节点为一组进行翻转



### 1.9 合并两个有序链表



### 1.10 在只给定单链表中某个节点指针的情况下删除该节点



### 1.11 判断两个单链表（无环）是否交叉



### 1.12 展开链表列表 