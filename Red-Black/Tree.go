package Red_Black

const RED = true
const BLACK = false

type Node struct {
	k     int
	v     int
	left  *Node
	right *Node

	color bool
}

func InitNode(k, v int) *Node {
	return &Node{
		v:     v,
		k:     k,
		left:  nil,
		right: nil,
		color: RED,
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
func isRed(node *Node) bool {
	if node == nil {
		return BLACK
	}
	return node.color
}

func (this *Tree) Push(k, v int) {
	this.root = this.push(this.root, k, v)

	this.root.color = BLACK
}

func (this *Tree) push(node *Node, k, v int) *Node {

	if node == nil {
		this.size++
		return InitNode(k, v)
	}

	if k > node.k {
		node.right = this.push(node.right, k, v)
	} else if k < node.k {
		node.left = this.push(node.left, k, v)
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

func (this *Tree) leftRotate(node *Node) *Node {
	x := node.right
	node.right = x.left
	x.left = node
	x.color = node.color
	node.color = RED
	return x
}

func (this *Tree) flipColors(node *Node) {
	node.color = RED
	node.left.color, node.right.color = BLACK, BLACK
}

func (this *Tree) rightRotate(node *Node) *Node {
	x := node.left
	node.left = x.right
	x.right = node
	x.color = node.color
	node.color = RED
	return x
}

func (this *Tree) getNode(node *Node, k int) *Node {
	if node == nil {
		return nil
	}
	if k == node.k {
		return node
	} else if k < node.k {
		return this.getNode(node.left, k)
	} else {
		return this.getNode(node.right, k)
	}
}
func (this *Tree) Contains(key int) bool {
	return this.getNode(this.root, key) != nil
}

func (this *Tree) GetValue(key int) *int {
	node := this.getNode(this.root, key)
	if node == nil {
		return nil
	} else {
		return &node.v
	}
}
func (this *Tree) SetNewValue(key, value int) {
	node := this.getNode(this.root, key)
	if node == nil {
		panic("key not found")
	}
	node.v = value
}

func (this *Tree) minimum(node *Node) *Node {
	if node.left == nil {
		return node
	}
	return this.minimum(node.left)
}

func (this *Tree) removeMin(node *Node) *Node {
	if node.left == nil {
		rightNode := node.right
		node.right = nil
		this.size--
		return rightNode
	}
	node.left = this.removeMin(node.left)
	return node
}

func (this *Tree) Remove(node *Node, key int) *Node {
	if node == nil {
		return node
	}
	if key < node.k {
		node.left = this.Remove(node.left, key)
		return node
	} else if key > node.k {
		node.right = this.Remove(node.right, key)
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
