package AVL

// Node 表示 AVL 树节点，记录高度以便计算平衡因子。
type Node struct {
	E      int
	height int
	Left   *Node
	Right  *Node
}

// Tree 是简单的整型 AVL 树实现。
type Tree struct {
	root *Node
	size int
}

func newNode(e int) *Node {
	return &Node{E: e, height: 1}
}

func height(n *Node) int {
	if n == nil {
		return 0
	}
	return n.height
}

func updateHeight(n *Node) {
	lh, rh := height(n.Left), height(n.Right)
	if lh > rh {
		n.height = lh + 1
		return
	}
	n.height = rh + 1
}

func balanceFactor(n *Node) int {
	return height(n.Left) - height(n.Right)
}

// rightRotate 和 leftRotate 负责局部旋转，恢复平衡。
func rightRotate(y *Node) *Node {
	x := y.Left
	t3 := x.Right

	x.Right = y
	y.Left = t3

	updateHeight(y)
	updateHeight(x)
	return x
}

func leftRotate(y *Node) *Node {
	x := y.Right
	t2 := x.Left

	x.Left = y
	y.Right = t2

	updateHeight(y)
	updateHeight(x)
	return x
}

func rebalance(n *Node) *Node {
	if n == nil {
		return nil
	}
	updateHeight(n)
	bf := balanceFactor(n)

	// 左重
	if bf > 1 {
		if balanceFactor(n.Left) < 0 {
			n.Left = leftRotate(n.Left)
		}
		return rightRotate(n)
	}
	// 右重
	if bf < -1 {
		if balanceFactor(n.Right) > 0 {
			n.Right = rightRotate(n.Right)
		}
		return leftRotate(n)
	}
	return n
}

func (t *Tree) Size() int {
	return t.size
}

func (t *Tree) IsEmpty() bool {
	return t.size == 0
}

func (t *Tree) Root() *Node {
	return t.root
}

func (t *Tree) Add(e int) {
	t.root = t.add(t.root, e)
}

func (t *Tree) add(node *Node, e int) *Node {
	if node == nil {
		t.size++
		return newNode(e)
	}

	if e < node.E {
		node.Left = t.add(node.Left, e)
	} else if e > node.E {
		node.Right = t.add(node.Right, e)
	} else {
		return node
	}
	return rebalance(node)
}

func (t *Tree) Contains(e int) bool {
	return contains(t.root, e)
}

func contains(node *Node, e int) bool {
	if node == nil {
		return false
	}
	if e == node.E {
		return true
	}
	if e < node.E {
		return contains(node.Left, e)
	}
	return contains(node.Right, e)
}

func (t *Tree) Remove(e int) {
	t.root = t.remove(t.root, e)
}

func (t *Tree) remove(node *Node, e int) *Node {
	if node == nil {
		return nil
	}

	if e < node.E {
		node.Left = t.remove(node.Left, e)
	} else if e > node.E {
		node.Right = t.remove(node.Right, e)
	} else {
		// 命中节点
		if node.Left == nil {
			right := node.Right
			node.Right = nil
			t.size--
			return right
		}
		if node.Right == nil {
			left := node.Left
			node.Left = nil
			t.size--
			return left
		}
		// 用后继替换当前节点值，然后在右子树删除后继
		succ := minimum(node.Right)
		node.E = succ.E
		node.Right = t.remove(node.Right, succ.E)
	}
	return rebalance(node)
}

func minimum(node *Node) *Node {
	cur := node
	for cur.Left != nil {
		cur = cur.Left
	}
	return cur
}

func inorder(node *Node, res *[]int) {
	if node == nil {
		return
	}
	inorder(node.Left, res)
	*res = append(*res, node.E)
	inorder(node.Right, res)
}

func (t *Tree) InOrder() []int {
	var res []int
	inorder(t.root, &res)
	return res
}

func (t *Tree) IsBST() bool {
	arr := t.InOrder()
	for i := 1; i < len(arr); i++ {
		if arr[i-1] > arr[i] {
			return false
		}
	}
	return true
}

func isBalanced(node *Node) bool {
	if node == nil {
		return true
	}
	lh, rh := height(node.Left), height(node.Right)
	if lh-rh > 1 || rh-lh > 1 {
		return false
	}
	return isBalanced(node.Left) && isBalanced(node.Right)
}

func (t *Tree) IsBalanced() bool {
	return isBalanced(t.root)
}
