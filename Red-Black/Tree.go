package Red_Black

const RED  = true
const BLACK  = false

type Node struct {
	k int
	v int
	left *Node
	right *Node
	// 这里我们用true代表红，false代表黑
	color bool
}
func InitNode(k,v int) *Node {
	return &Node{
		v:v,
		k:k,
		left:nil,
		right:nil,
		color:RED,
	}
}
type Tree struct {
	size int
	root *Node
}

func (this *Tree) Size() int {
	return this.size
}
func (this *Tree) IsEmpty() bool {
	return this.size == 0
}
func isRed (node *Node) bool {
	if node == nil {
		return BLACK
	}
	return node.color
}

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

// 红黑树颜色翻转
func (this *Tree) flipColors (node *Node) {
	node.color = RED
	node.left.color,node.right.color = BLACK,BLACK
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

func (this *Tree) getNode (node *Node,k int) *Node {
	if node == nil {
		return nil
	}
	if k == node.k {
		return node
	} else if k < node.k {
		return this.getNode(node.left,k)
	} else {
		return this.getNode(node.right,k)
	}
}
func (this *Tree) Contains (key int) bool {
	return this.getNode(this.root,key) != nil
}

func (this *Tree) GetValue (key int) *int {
	node := this.getNode(this.root,key)
	if node == nil {
		return nil
	} else {
		return &node.v
	}
}
func (this *Tree) SetNewValue (key,value int) {
	node := this.getNode(this.root,key)
	if node == nil {
		panic("没有这个key值")
	}
	node.v = value
}

func (this *Tree) minimum (node *Node) *Node {
	if node.left == nil {
		return node
	}
	return this.minimum(node.left)
}

func (this *Tree)removeMin(node *Node) *Node {
	if node.left == nil {
		rightNode := node.right
		node.right = nil
		this.size--
		return rightNode
	}
	node.left = this.removeMin(node.left)
	return node
}

func (this *Tree) Remove (node *Node,key int) *Node {
	if node == nil {
		return node
	}
	if key < node.k {
		node.left = this.Remove(node.left,key)
		return node
	} else if key > node.k {
		node.right = this.Remove(node.right,key)
		return node.right
	} else {
		if node.left == nil {
			rightNode := node.right
			node.right = nil
			this.size--
			return rightNode
		}
		if node.right == nil {
			leftNode := node.left
			node.left = nil
			this.size--
			return leftNode
		}
		successor := this.minimum(node.right)
		successor.right = this.removeMin(node.right)
		successor.left = node.left
		return successor
	}
}