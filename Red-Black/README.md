
## 红黑树
### 红黑树的定义

1. 每个节点或者是红色的，或者是黑色的；
2. 根节点是黑色的；
3. 每个叶子结点（最后的空节点）是黑色的；
4. 如果一个节点是红色的，那么他的孩子都是黑色的；
5. 从任意一个节点到叶子节点经过的黑色节点是一样的。

只看红黑树的定义，如果就能把红黑树的数据结构想出来那你真是一个天才。不得不说红黑树的定义比较难懂。所以我们先从比较简单的2-3树开始看起，然后再去证明红黑树与2-3树的等价性。

### 2-3树

2-3树是一种**绝对平衡的一种数据结构**

![](http://www.liuanqihappybirthday.top/uploads/big/5a72ebda2c60ed419e5fbf2f6acc9915.png)

#### 基本性质

节点可以存放一个元素或者两个元素



![img](http://www.liuanqihappybirthday.top/uploads/big/54fcd7ed92ae67eb24d8ccd597187d1a.png)

#### 添加节点

**新节点永远不会添加到空子树中去，（新节点的添加不会自己新建一个子树）他只会和最后一个子树融合**

1. 对于一个空的2-3树，添加的元素作为根节点。

   

   ![img](http://www.liuanqihappybirthday.top/uploads/big/9d4a56ab323af5a99a65cc7aeaa622ca.png)

   

2. 添加37。
   因为37比42小，所以应该添加到42的左子树上，但是此时2-3树42没有左子树，所以将37和42进行融合，成为一个节点。

   

   ![img](http://www.liuanqihappybirthday.top/uploads/big/2d8528df836e0ddbe8fbbfa73d9b3625.png)

   

3. 添加12
   因为12比37小，所以应该添加到37的左子树，但是37的左子树为空，所以和37节点进行融合。暂时形成可以存储三个元素的节点。

   

   ![img](http://www.liuanqihappybirthday.top/uploads/big/e2ec4e723b8dbfe8a5049f998226b20f.png)

   2-3树最多只能存储2个元素，所以该节点要进行分裂。

   

   ![img](http://www.liuanqihappybirthday.top/uploads/big/415dd3a3deaa199578d27a8a71bdd81e.png)

4. 添加18
   按上述道理添加到12的右子树，但是12节点的右子树为空，所以和12节点进行融合。

   

   ![img](http://www.liuanqihappybirthday.top/uploads/big/62175e677344a9fc878b23efd65695bb.png)

5. 添加6
   按照上述推论和12节点融合，暂时形成有三个元素的节点。

   

   ![img](http://www.liuanqihappybirthday.top/uploads/big/e99465f33160e57a0000f72ce53b1cba.png)

   1. 进行分裂

   

   ![img](http://www.liuanqihappybirthday.top/uploads/big/9eaaf063073fff83c44dfc9c320e8374.png)

   

   2. 当前2-3树不是绝对平衡的树，分裂之后6,12,18形成新的树，12作为当前子树的根节点，要和他的双亲进行融合。

      

      ![img](http://www.liuanqihappybirthday.top/uploads/big/11cf18f0514b6fb6d61a6bcaa25695b8.png)

      

      融合后，回复平衡。

6. 添加11

   

   ![img](http://www.liuanqihappybirthday.top/uploads/big/0674270666019e821c24cdb50a074be7.png)

   image.png

7. 添加5

   1. 



![img](http://www.liuanqihappybirthday.top/uploads/big/16c06159a5f064db3a1c5bf55f4ce2c2.png)



2.





![img](http://www.liuanqihappybirthday.top/uploads/big/1cdc7419eaaaee722ef6e6d28d45921d.png)



3.





![img](http://www.liuanqihappybirthday.top/uploads/big/546d19f6234d2dc1f13640e1be985a34.png)


4.



![img](http://www.liuanqihappybirthday.top/uploads/big/2fdfdb63bec9e96c775a10674be90175.png)

### 红黑树与2-3树的等价性

因为普通的二叉搜索树只能存储两个指针域，所以，想让二叉树的节点存储三个指针域可以将两个二叉树的节点看成是一个2-3树的节点，如图：

红黑树是一种平衡二叉树，只有一种节点。这种节点有两个儿子，和2-3树中的2-节点对应。

![img](http://www.liuanqihappybirthday.top/uploads/big/5a1ee7a083d43bd26695ea628c3792c2.png)

如何表示3-节点呢？我们尝试一种特殊的边：**默认情况下节点的颜色均为黑色。我们将某个节点染为红色，表示它和父亲的的链接是红色的**，就像下图：

![img](http://www.liuanqihappybirthday.top/uploads/big/706dd7fce0e9581a9ca3f689a9808386.png)
为了便于区分，可以用红色节点表示该节点何双亲表示同一个节点。



![img](http://www.liuanqihappybirthday.top/uploads/big/a1af3d11ee23871f3462b6ea871455ba.png)

所以一颗2-3树转换为红黑树后如图：

![img](http://www.liuanqihappybirthday.top/uploads/big/a44dc62a5f4e6c91aec4baa3387f45dd.png)


等价于

![img](http://www.liuanqihappybirthday.top/uploads/big/180978299cf6b635ece407bb66bf2238.png)

### 查询操作

由于红黑树是二叉搜索树，因此查询操作就是二叉搜索树的查询操作。时间复杂度为$\Theta(\text{log}N)$。

### 基本操作

在介绍红黑树的插入和删除操作前，首先介绍红黑树的一些基本操作。

#### 旋转

红黑树的旋转只有两种：右旋和左旋
红黑树的旋转操作是为了在保证二叉搜索树和红黑树的性质的前提下，来转换红链接的位置。

![rbtree-rotate-1](http://www.liuanqihappybirthday.top/uploads/big/37e2d08ef7a94c17717bb77dece16d0c.png)

可以看出右旋就是将节点的左儿子提上来，将自己变做它的右儿子，将左儿子的右子树接到自己的左子树中，同时转变红链接。可以将其想象成把`4->2`这条边右旋了一下。左旋也是类似的做法。同时右旋和左旋可以视为一对逆操作，因为一次左旋和一次右旋可以变回原来的样子。

代码：

```go
//   node                     x
//  /   \     左旋转         /  \
// T1   x   --------->   node   T3
//     / \              /   \
//    T2 T3            T1   T2
func (this *Tree) leftRotate (node *Node) *Node {
	x := node.right
	node.right = x.left
	x.left = node
	x.color = node.color
	node.color = RED
	return x
}
// 红黑树的右旋转过程
//     node                   x
//    /   \     右旋转       /  \
//   x    T2   ------->   y   node
//  / \                       /  \
// y  T1                     T1  T2
func (this *Tree) rightRotate (node *Node) *Node{
	x := node.left
	node.left = x.right
	x.right = node
	x.color = node.color
	node.color = RED
	return x
}

```

#### 反色

如同在2-3树中一样，红黑树要能够处理4-节点。
对于4-节点，我们只有两种操作：合成一个4-节点和分解一个4-节点。

![rbtree-flip-1](http://www.liuanqihappybirthday.top/uploads/big/82c1d9cd5bdd832bf058ef0a69ade3c3.png)

对照一下2-3树，这个操作就显而易见了。

![img](http://www.liuanqihappybirthday.top/uploads/big/e2ec4e723b8dbfe8a5049f998226b20f.png)

![img](http://www.liuanqihappybirthday.top/uploads/big/415dd3a3deaa199578d27a8a71bdd81e.png)

也许你会注意到反色操作会将两个儿子的父节点变为红色，是因为在2-3树中，中间取出来的键要向上传递并结合进去。此外，反色操作会导致出现右边的红链接，然而这没有关系，因为4-节点是临时的，我们最终会通过左旋其变为左边的红链接或者再次反色将这个4-节点分解。

```go
// 红黑树颜色翻转
func (this *Tree) flipColors (node *Node) {
	node.color = RED
	node.left.color,node.right.color = BLACK,BLACK
}
```

### 插入操作

为了探究红黑树的插入操作，我们依然回到2-3树。在2-3树中，我们将新插入的节点与上面的节点合并，然后再做调整。为了表示合并，我们将**新插入的节点均设为红色**，表示与上面的节点相连接。

然而插入后，新的红节点可能会违反我们的规定，因此需要在回溯的时候进行调整。

![](http://www.liuanqihappybirthday.top/uploads/big/6aa7a84d32bd37dc0b33d23f7d3d8e92.png)

#### 情况一：调整右边的红链接

当我们发现某个节点的**左儿子是黑色**但**右儿子是红色**时，我们要将右边的红色链接转到左边来：

![rbtree-insert-1](https://git.oschina.net/riteme/blogimg/raw/master/rbtree-and-2-3-tree/rbtree-insert-1.png)

如上图，通过对`b`**左旋**，完成了对红链接位置的纠正。
这样做是为了方便接下来的操作。

#### 情况二：分解4-节点

在情况一中，我们要求节点的左儿子是黑色。这是因为当**左儿子和右儿子都是红色**时，就代表着一个4-节点，为此我们可以直接将其反色来分解它：

![rbtree-insert-2](https://git.oschina.net/riteme/blogimg/raw/master/rbtree-and-2-3-tree/rbtree-insert-2.png)

如果该操作是在根节点上，那么整棵红黑树的黑高将会加1。

#### 情况三：连续的红色左儿子

在情况一中，我们能够把所有的右边的红色节点转到左边来，这样就好判断是否存在4-节点。除了情况二中的4-节点外，**连续的两个红色左儿子**也将表示一个4-节点：

![rbtree-insert-3](https://git.oschina.net/riteme/blogimg/raw/master/rbtree-and-2-3-tree/rbtree-insert-3.png)

对此，我们的做法是将节点**右旋**，从而变为了情况二。

```go
// 向红黑树中添加新元素
// 回忆2-3树中我们添加节点，
// 或者添加到一个2节点中，形成一个3节点，
// 或者添加到一个三节点中，形成一个暂时的4节点
// 所以我们永远先添加一个红节点，先融合再调整
func (this *Tree) Push (k,v int) {
	this.root = this.push(this.root,k,v)
	// 我们还需要保持红黑树的根节点为黑节点
	this.root.color = BLACK
}

func (this *Tree) push (node *Node,k,v int) *Node {
	// 因为红黑树的红节点都需要左倾，所以插入较大节点后需要进行一次左旋操作
	if node == nil {
		this.size++
		return InitNode(k,v)
	}

	if k > node.k {
		node.right = this.push(node.right,k,v)
	} else if k < node.k {
		node.left = this.push(node.left,k,v)
	} else {
		node.v = v
	}

	if isRed(node.right) && !isRed(node.left) {
		node = this.leftRotate(node)
	}
	if isRed(node.left) && isRed(node.left.left) {
		node = this.rightRotate(node)
	}
	if isRed(node.left) && isRed(node.right) {
		this.flipColors(node)
	}
	return node
}
```

