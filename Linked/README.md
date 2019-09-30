# 单向链表

[TOC]



## 单向链表

单向链表属于链表的一种，也叫单链表，单向即是说它的链接方向是单向的，它由若干个节点组成，每个节点都包含下一个节点的指针。

数据结构：

```go
type Node struct {
	E int
	Next *Node
}

type List struct {
	dummyHead *Node
	size int
}

func (l List) Head() *Node {
	return l.dummyHead
}

func (l List) Size() int {
	return l.size
}
```



## 单链表特点

- 创建单链表时无需指定链表的长度，这个比起数组结构更加有优势，而数组纵使实现成动态数组也是需要指定一个更大的数组长度，而且要把原来的数组元素一个个复制到新数组中。
- 单链表中的节点删除操作很方便，它可以直接改变指针指向来实现删除操作，而某些场景下数组的删除会导致移动剩下的元素。
- 单链表中的元素访问需要通过顺序访问，即要通过遍历的方式来寻找元素，而数组则可以使用随机访问，这点算是单链表的缺点。

## 单链表创建

创建一个空链表，这里用到了虚拟头结点的方式

```go
func initNode (e int) *Node {
	return &Node{
		E:e,
		Next:nil,
	}
}
func InitList () *List {
	return &List{
		dummyHead:initNode(0),
		size:0,
	}
}
```

## 链表的插入

```go
// 在链表的第index索引个元素后插入元素,索引从0开始
func (this *List) AddIndex (index,e int) {
	if index > this.size || index < 0 {
		panic("索引越界，不能插入了")
	}
	prev := this.dummyHead
	node := initNode(e)

	for i:=0;i<index;i++ {
		prev = prev.Next
	}
	node.Next = prev.Next
	prev.Next = node
	this.size++

}
// 在链表头添加元素
func (this *List) AddFirst (e int) {
	this.AddIndex(0,e)
}
// 在链表尾部添加节点
func (this *List) AddLast (e int) {
	this.AddIndex(this.size,e)
}
```

## 删除节点

```go
// 在链表中删除元素
func (this *List) Remove (index int) int {
	if index > this.size || index < 0 {
		panic("索引越界，不能删除")
	}
	prev := this.dummyHead
	for i:=0;i<index;i++ {
		prev = prev.Next
	}
	retNode := prev.Next
	prev.Next = retNode.Next
	this.size--
	return retNode.E
}
func (this *List) RemoveFirst (index int) int{
	return this.Remove(0)
}
func (this *List) RemoveLast (index int) int{
	return this.Remove(this.size-1)
}
func (this *List) RemoveElement (e int) {
	prev := this.dummyHead
	for prev.Next != nil {
		if prev.E == e {
			break
		}
		prev = prev.Next
	}
	if prev.Next != nil {
		DelNode := prev.Next
		prev.Next = DelNode.Next
		DelNode = nil
	}
}
```

## 查询结点

```go
// 在链表中查询第index个元素
func (this *List) Get (index int) int {
	if index > this.size || index < 0 {
		panic("索引越界，不能查询")
	}
	cur := this.dummyHead.Next
	for i:=0;i<index;i++ {
		cur = cur.Next
	}
	return cur.E
}
func (this *List) GetFirst (index int) int{
	return this.Get(0)
}
func (this *List) GetLast (index int) int{
	return this.Get(this.size-1)
}
```

## 修改结点

```go
// 在链表index个位置中放入元素e
func (this *List) Set (index,e int) {
	if index > this.size || index < 0 {
		panic("索引越界，不能置入")
	}
	cur := this.dummyHead.Next
	for i:=0;i<index;i++ {
		cur = cur.Next
	}
	cur.E = e
}
// 在链表中查询是否包括元素e
func (this *List) Contains (e int) bool {
	cur := this.dummyHead.Next
	for cur!=nil {
		if cur.E == e{
			return true
		}
		cur = cur.Next
	}
	return false
}
```

