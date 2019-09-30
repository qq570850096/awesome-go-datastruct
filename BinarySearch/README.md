# 二分搜索树

[TOC]



在理解二分搜索树之前，我们先来看看二叉树是什么。

## 1.1 二叉树

二叉树也是一种动态的数据结构。每个节点只有两个叉，也就是两个孩子节点，分别叫做左孩子，右孩子，而没有一个孩子的节点叫做叶子节点。每个节点最多有一个父亲节点，最多有两个孩子节点(也可以没有孩子节点或者只有一个孩子节点)。47左半边的所有节点组合起来形成了47的左子树，47右半边所有节点结合起来形成了47的右子树。如下图所示：



![img](http://exia.gz01.bdysite.com/uploads/big/2827ed392cfebe4fc9041712791fd3ac.png)

1-1

综合一下，涉及到的概念有：

> 根节点：二叉树的起始节点，唯一没有父亲节点的节点；
> 父亲节点：每个节点只有一个父亲节点。如上图47就是35的父亲节点；
> 左右孩子节点：每个节点至多拥有两个孩子节点，分别叫左孩子，右孩子；
> 左子树右子树：每个节点左边或者右边部分所有节点组合成的树结构。

## 1.2 二分搜索树

### 1.2.1 性质

第一，二分搜索树是一颗二叉树，满足二叉树的所有定义；
第二，二分搜索树每个节点的左子树的值都小于该节点的值，每个节点右子树的值都大于该节点的值。
第三，任意节点的每颗子树都满足二分搜索树的定义。



![img](http://exia.gz01.bdysite.com/uploads/big/f3e3cd20b6e8a1c597e83c40fd724147.png)

1-2

### 1.2.2 意义

当我们看到二分搜索树的定义时，是否会去联系这样定义的意义何在呢？其实，二分搜索树是在给数据做整理，因为左右子树的值和根节点存在大小关系，**所以在查找元素时，我们于根节点进行对比后，就能每次近乎一半的去除掉查找范围，这就大大的加快了我们的查询速度，插入元素时也是一样。**

在图1-2中，如果要查找元素55，那么和根节点47对比后，发现55比47大，于是就往47右孩子60中去查询，接着发现55比60小，就往60左孩子中查询，于是就找到了这个元素。想象一下，如果是一个链表，那么将一个一个查询下去，速度可想而知。

其实在生活中，这样的例子也比比皆是，我们去超市买东西，超市也把一二三楼卖的是啥写的很清楚，假如三楼卖的是生鲜果蔬，而我们要买今天的菜，那么我们就直接去三楼，一楼和二楼我们就可以不用去找了，大大加快了我们选购商品的速度。图书馆找书也是这样的例子。所以二分搜索树的意义也就在此，**很多时候数据结构其实来源于生活，于生活中解决实际问题，这就是技术的力量和价值的体现。**

> 但是，为了达到这样的高效性，树结构由此也需要每个节点之间具备可比较性。而链表数据结构就没有这类要求，所以还是那句话，有得必有失。

**注意**：本篇文章中关于二分搜索树做了没有重复元素的假定，如果遇到重复元素则不插入。

### 1.2.3 二分搜索树的数据结构

```go
type Node struct {
	E int
	Left *Node
	Right *Node
}

// 约定在此样例代码中我们的二分搜索树中没有重复元素
// 如果想包涵重复元素的话，只需要以下定义：
// 左子树小于等于此节点，或右子树大于等于节点
type Tree struct {
	root *Node
	size int
}
// 以下是getter方法
func (t Tree) Size() int {
	return t.size
}

func (t Tree) Root() *Node {
	return t.root
}
// 对节点Node做初始化的方法
func InitNode(E int) *Node {
	return &Node{
		E:E,
		Left:nil,
		Right:nil,
	}
}
```

### 1.2.4 插入元素

我们一起来看看在一个树中插入元素的动画演示过程，如下图元素65插入树结构所示，元素在插入时，不断跟当前根节点进行对比，以来选择插入到左子树还是右子树。

![](http://exia.gz01.bdysite.com/uploads/big/2c461a5af08de072bd0918de54de9631.gif)

```go
// 向二分搜索树添加元素的公有方法
func (this *Tree) AddE(e int) {
    // 如果根节点为空，那么直接把这个node作为根节点即可
	if this.root == nil {
		this.root = InitNode(e)
		this.size++
    // 如果根节点不为空，那么进入递归的方式插入node节点
	} else {
		this.add(this.root,e)
	}
}
// 向二分搜索树添加元素的私有方法
// 向以node为根的二分搜索树中插入元素E，递归算法
func (this *Tree) add(node *Node,e int){

	// 不管是递归还是回溯，首先我们都应该先写出递归的结束条件是什么
	if e == node.E {
		return
	} else if e > node.E && node.Right == nil{
		node.Right = InitNode(e)
		this.size++
		return
	} else if  e < node.E && node.Left == nil {
		node.Left = InitNode(e)
		this.size++
		return
	}
	// 如果e大于当前节点的e，那么应该插入到右子树中
	if e > node.E {
		this.add(node.Right, e)
    // 反之则应该插入左子树中
	} else {
		this.add(node.Left, e)
	}
}

```

刚刚我们动画演示了一个插入右子树的过程，接下来的动画演示插入左子树:

![](http://exia.gz01.bdysite.com/uploads/big/d61183686555c4950cb89f6d7735fb4c.gif)

测试代码:

```go
package BinarySearch

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestTree(t *testing.T) {
	tree := &Tree{}
	fmt.Println(tree.IsEmpty())
	for i:=0;i<10;i++ {
		temp := rand.Intn(10000)
		tree.AddE(temp)
	}
	fmt.Println(tree.size)
	fmt.Println(tree.IsEmpty())
}
```

测试结果：

```
=== RUN   TestTree
true
10
false
--- PASS: TestTree (0.00s)
PASS
```

#### 对于插入元素的递归条件的优化

在上文中我们已经可以准确的实现二分搜索树的建立过程了，但是还有一个问题就是：**我们的递归条件其实非常的臃肿**

1. 我们对于e和node.E的值做了两次判断
2. 我们判断了左右子树是否为空，增加了递归条件的复杂性

对于这个递归条件，其实我们可以这么思考一下，**我们其实不需要判断左右子树的情况，只要e>E我们就向右走，反之同理。那么最后我们一定会走到一个值为空的子树上，那么对于二叉树来说，空也是一颗子树。所以我们的递归结束条件可以简化成这样:**

```go
if node == nil {
	return InitNode(e)
}
```

那么这样又引出了新的问题，我们一开始的add函数是没有返回值的，现在要返回一个Node指针，那么由谁来接受呢？

答案其实很简单：**哪边的子树发生了改变哪边来接收这个变化**

```go
func (this *Tree) AddE(e int) {
	this.root = this.add(this.root,e)
}

func (this *Tree) add(node *Node,e int) *Node{

	// 不管是递归还是回溯，首先我们都应该先写出递归的结束条件是什么
	if node == nil {
		return InitNode(e)
	}

	if e > node.E {
		node.Right = this.add(node.Right, e)
	} else if e < node.E {
		node.Left = this.add(node.Left, e)
	}
	return node
}
```

这样优化后，我们的代码就可以更加的简洁。

### 1.2.5二分搜索树的遍历

树的遍历分为前序遍历，中序遍历，后序遍历，层序遍历。我们依次来讲一讲。

前序遍历，就是我们在递归调用前，先做我们的逻辑（这里的逻辑就是打印一下当前元素），前序遍历代码如下：

```go
// 递归式前序遍历
func PreOrder(node *Node)  {
	if node == nil {
		return
	}
	fmt.Printf("%d ",node.E)
	PreOrder(node.Left)
	PreOrder(node.Right)
}
// 非递归的前序遍历
func (this *Tree) PreOrderNR() {
	stack := make([]*Node, 0)
	stack = append(stack, this.root)
	for len(stack) > 0 {
		curNode := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		fmt.Printf("%d ",curNode.E)
		if curNode.Right != nil {
			stack = append(stack, curNode.Right)
		}
		if curNode.Left != nil {
			stack = append(stack, curNode.Left)
		}
	}
	fmt.Println()
}
```

那么接下来就是中序遍历，中序遍历就是把对元素的操作操作放在了中间，先不断递归左孩子，然后对元素进行操作，最后递归右孩子。所以很容易的推理出来，中序遍历是对元素从小到大的排序遍历。这个性质可以作为判断二分搜索树的一个条件。

同理可以推出还有后续遍历，这里就不再赘述了。直接看代码

```go
// 2.中序遍历
func MidOrder(node *Node)  {
	if node == nil {
		return
	}

	MidOrder(node.Left)
	fmt.Printf("%d ",node.E)
	MidOrder(node.Right)
}
// 3.后序遍历
func BackOrder(node *Node)  {
	if node == nil {
		return
	}
	BackOrder(node.Left)
	BackOrder(node.Right)
	fmt.Printf("%d ",node.E)
}
```

最后我们一起来聊聊层序遍历，顾名思义，层序遍历就是对树结构一层一层的遍历。要做到这一点，我们可以很方便的想到利用队列来实现这一过程。我们每次从队列中取出一个元素时，接着就把该元素的左右两孩子推入队列中，然后依次取出元素，以此来做到按层把元素遍历一遍。

一起来看看动画爸爸给我们的演示，最直观。

![](http://exia.gz01.bdysite.com/uploads/big/37e56b12709d542802c5b882e94a3739.gif)

代码:

```go
// 二分搜索树的层序遍历
func (this *Tree)LevelOrder(){
	queue := make([]*Node,0)
	queue = append(queue, this.root)
	for len(queue) > 0 {
		curNode := queue[0]
		queue = queue[1:]
		fmt.Printf("%d ",curNode.E)
		if curNode.Left != nil {
			queue = append(queue,curNode.Left)
		}
		if curNode.Right != nil {
			queue = append(queue,curNode.Right)
		}
	}
}
```

### 1.2.6二分搜索树删除元素

接下来，我们聊聊元素的删除。在此之前，我们先得引入一个概念，元素的前驱或者后继节点。

> 后继节点：一个节点右子树中，最小的节点为该节点的后继节点。后继节点是比该节点所有大的元素中最小的元素。

之所以要引入这个概念，原因是在删除元素时，我们需要找一个元素替代被删除位置的元素，但是由于二分搜索树的特性，不能随便找元素过来代替，必须得找一个和被删除元素最接近的元素来替代其位置。所以找前驱或后继替代都可以。

比如说，如下图，47的后继节点就是55，右子树中最小的元素。

![](http://exia.gz01.bdysite.com/uploads/big/f3e3cd20b6e8a1c597e83c40fd724147.png)

我们先看如下两种简单的删除情况：

1. 删除最大节点
2. 删除最小节点

通过二分搜索树的定义和图示我们可以知道——最大节点在二叉树最右边，最小节点在二叉树最左边，也就是是**最大节点没有右子树，最小节点没有左子树**

```go
// 删除掉以node为根的二分搜索树的最小节点
// 返回删除节点后新的二分搜索树的根
func (this *Tree) rmMin(node *Node) *Node {
	if node.Left == nil {
		nodeRight := node.Right
		node.Right = nil
		this.size--
		return nodeRight
	}
	node.Left = this.rmMin(node.Left)
	return node
}
// 删除掉以node为根的二分搜索树的最大节点
// 返回删除节点后新的二分搜索树的根
func (this *Tree) rmMax(node *Node) *Node {
	if node.Right == nil {
		nodeLeft := node.Left
		node.Left = nil
		this.size--
		return nodeLeft
	}
	node.Right = this.rmMax(node.Right)
	return node
}
```

那么我们再看更复杂的情况，也就是删除任意节点时，有可能左右子树都不为空，我们可以按照图示思路去删除

![](http://exia.gz01.bdysite.com/uploads/big/b5594ce301cdb91d97f41c284bb7d7c6.gif)

第一步，找到待删除元素；
第二步，找到待删除元素的后继节点；
第三步，把待删除元素的左子树赋值给后继节点的左子树，把待删除元素的右子树最小元素删除后，赋值给后继节点的右子树。

代码:

```go
func (this *Tree) remove(node *Node,e int) *Node {
	if node == nil {
		return nil
	}
	if e > node.E {
		node.Right = this.remove(node.Right,e)
		return node
	} else if e < node.E {
		node.Left = this.remove(node.Left,e)
		return node
	} else {
		// 如果左子树为空的时候
		if node.Left == nil {
			nodeRight := node.Right
			node.Right = nil
			this.size--
			return nodeRight
		}
		// 如果右子树为空
		if node.Right == nil {
			nodeLeft := node.Left
			node.Left = nil
			this.size--
			return nodeLeft
		}
		// 如果左右子树都不为空，那么我们需要找到node的后继
		// 就是所有比node值大的节点中值最小的那个，显然它在node的右子树中
		nodeNext := minimum(node.Right)
		nodeNext.Right = this.rmMin(node.Right)
		nodeNext.Left = node.Left
		node.Left = nil
		node.Right = nil
		return nodeNext
	}
}
```

