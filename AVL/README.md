## AVL 自平衡二叉搜索树

### 定义

**AVL树**（Adelson-Velsky and Landis Tree）是一种自平衡二叉搜索树，由苏联数学家 G.M. Adelson-Velsky 和 E.M. Landis 于1962年发明。在AVL树中，任意节点的左右子树高度差（平衡因子）的绝对值不超过1，这保证了树的高度始终维持在 O(log n) 级别。

### 为什么使用 AVL 树？

普通的二分搜索树在最坏情况下会退化成链表，导致查找、插入、删除操作的时间复杂度退化为 O(n)。AVL树通过旋转操作维持平衡，确保所有操作的时间复杂度稳定在 O(log n)。

**应用场景：**
- 需要频繁进行插入、删除操作的有序集合
- 数据库索引
- 内存中的有序数据管理
- 需要保证最坏情况性能的场景

### 特性

1. **高度平衡**：任意节点的平衡因子（左子树高度 - 右子树高度）只能是 -1、0、1
2. **自动调整**：插入或删除后通过旋转自动恢复平衡
3. **查找效率稳定**：不会退化成链表

| 操作 | 平均时间复杂度 | 最坏时间复杂度 |
|------|---------------|---------------|
| 查找 | O(log n) | O(log n) |
| 插入 | O(log n) | O(log n) |
| 删除 | O(log n) | O(log n) |
| 空间复杂度 | O(n) | O(n) |

### 数据结构

```go
// Node 表示 AVL 树节点，记录高度以便计算平衡因子
type Node struct {
    E      int   // 节点存储的元素值
    height int   // 以该节点为根的子树高度
    Left   *Node // 左子节点
    Right  *Node // 右子节点
}

// Tree 是简单的整型 AVL 树实现
type Tree struct {
    root *Node // 根节点
    size int   // 树中元素个数
}
```

### 核心方法实现

#### 高度与平衡因子计算

```go
// 获取节点高度，空节点高度为0
func height(n *Node) int {
    if n == nil {
        return 0
    }
    return n.height
}

// 更新节点高度 = max(左子树高度, 右子树高度) + 1
func updateHeight(n *Node) {
    lh, rh := height(n.Left), height(n.Right)
    if lh > rh {
        n.height = lh + 1
        return
    }
    n.height = rh + 1
}

// 计算平衡因子：左子树高度 - 右子树高度
// 正数表示左重，负数表示右重
func balanceFactor(n *Node) int {
    return height(n.Left) - height(n.Right)
}
```

#### 旋转操作

AVL树的核心在于四种旋转操作：

```go
// 右旋转（LL情况）
//        y                              x
//       / \                           /   \
//      x   T4     向右旋转 (y)        z     y
//     / \       - - - - - - - ->    / \   / \
//    z   T3                       T1  T2 T3 T4
//   / \
// T1   T2
func rightRotate(y *Node) *Node {
    x := y.Left
    t3 := x.Right

    x.Right = y
    y.Left = t3

    updateHeight(y)
    updateHeight(x)
    return x
}

// 左旋转（RR情况）
//    y                             x
//  /  \                          /   \
// T1   x      向左旋转 (y)       y     z
//     / \   - - - - - - - ->   / \   / \
//   T2   z                   T1  T2 T3 T4
//       / \
//      T3 T4
func leftRotate(y *Node) *Node {
    x := y.Right
    t2 := x.Left

    x.Left = y
    y.Right = t2

    updateHeight(y)
    updateHeight(x)
    return x
}
```

#### 再平衡操作

```go
// 对节点进行再平衡，处理四种不平衡情况
func rebalance(n *Node) *Node {
    if n == nil {
        return nil
    }
    updateHeight(n)
    bf := balanceFactor(n)

    // LL：左子树的左侧插入导致不平衡，右旋转
    // LR：左子树的右侧插入导致不平衡，先左旋后右旋
    if bf > 1 {
        if balanceFactor(n.Left) < 0 {
            n.Left = leftRotate(n.Left) // LR情况：先左旋
        }
        return rightRotate(n) // LL情况：右旋
    }

    // RR：右子树的右侧插入导致不平衡，左旋转
    // RL：右子树的左侧插入导致不平衡，先右旋后左旋
    if bf < -1 {
        if balanceFactor(n.Right) > 0 {
            n.Right = rightRotate(n.Right) // RL情况：先右旋
        }
        return leftRotate(n) // RR情况：左旋
    }
    return n
}
```

#### 插入操作

```go
// 向AVL树中添加元素
func (t *Tree) Add(e int) {
    t.root = t.add(t.root, e)
}

func (t *Tree) add(node *Node, e int) *Node {
    // 递归终止条件：找到插入位置
    if node == nil {
        t.size++
        return newNode(e)
    }

    // 递归查找插入位置
    if e < node.E {
        node.Left = t.add(node.Left, e)
    } else if e > node.E {
        node.Right = t.add(node.Right, e)
    } else {
        return node // 元素已存在，不重复插入
    }

    // 回溯时进行再平衡
    return rebalance(node)
}
```

#### 删除操作

```go
// 从AVL树中删除元素
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
        // 命中待删除节点
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
        // 左右子树都不为空：用后继节点（右子树最小值）替换
        succ := minimum(node.Right)
        node.E = succ.E
        node.Right = t.remove(node.Right, succ.E)
    }
    // 回溯时进行再平衡
    return rebalance(node)
}

// 查找以node为根的子树的最小节点
func minimum(node *Node) *Node {
    cur := node
    for cur.Left != nil {
        cur = cur.Left
    }
    return cur
}
```

#### 验证方法

```go
// 中序遍历，返回有序数组
func (t *Tree) InOrder() []int {
    var res []int
    inorder(t.root, &res)
    return res
}

// 验证是否为二分搜索树
func (t *Tree) IsBST() bool {
    arr := t.InOrder()
    for i := 1; i < len(arr); i++ {
        if arr[i-1] > arr[i] {
            return false
        }
    }
    return true
}

// 验证是否平衡
func (t *Tree) IsBalanced() bool {
    return isBalanced(t.root)
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
```

### 测试用例

```go
func TestAVLTree(t *testing.T) {
    tree := &Tree{}

    // 插入元素
    elements := []int{30, 20, 40, 10, 25, 35, 50}
    for _, e := range elements {
        tree.Add(e)
    }

    // 验证BST性质
    fmt.Println("是否为BST:", tree.IsBST())       // true
    fmt.Println("是否平衡:", tree.IsBalanced())   // true
    fmt.Println("中序遍历:", tree.InOrder())      // [10 20 25 30 35 40 50]

    // 删除元素后验证
    tree.Remove(25)
    fmt.Println("删除25后是否平衡:", tree.IsBalanced()) // true
}
```

测试结果：
```
=== RUN   TestAVLTree
是否为BST: true
是否平衡: true
中序遍历: [10 20 25 30 35 40 50]
删除25后是否平衡: true
--- PASS: TestAVLTree (0.00s)
PASS
```

### 运行方式

```bash
go test ./AVL
```

### LeetCode 实战

#### [110. 平衡二叉树](https://leetcode-cn.com/problems/balanced-binary-tree/)

判断一棵二叉树是否为平衡二叉树：

```go
func isBalanced(root *TreeNode) bool {
    return height(root) >= 0
}

func height(node *TreeNode) int {
    if node == nil {
        return 0
    }
    leftHeight := height(node.Left)
    rightHeight := height(node.Right)

    // 如果子树不平衡，返回-1
    if leftHeight == -1 || rightHeight == -1 {
        return -1
    }
    // 如果当前节点不平衡，返回-1
    if abs(leftHeight - rightHeight) > 1 {
        return -1
    }
    return max(leftHeight, rightHeight) + 1
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}
```

#### [1382. 将二叉搜索树变平衡](https://leetcode-cn.com/problems/balance-a-binary-search-tree/)

将BST转换为平衡BST，利用AVL树思想：

```go
func balanceBST(root *TreeNode) *TreeNode {
    // 中序遍历获取有序数组
    var nodes []int
    inorder(root, &nodes)
    // 用有序数组构建平衡BST
    return buildBalancedBST(nodes, 0, len(nodes)-1)
}

func inorder(node *TreeNode, nodes *[]int) {
    if node == nil {
        return
    }
    inorder(node.Left, nodes)
    *nodes = append(*nodes, node.Val)
    inorder(node.Right, nodes)
}

func buildBalancedBST(nodes []int, left, right int) *TreeNode {
    if left > right {
        return nil
    }
    mid := left + (right-left)/2
    node := &TreeNode{Val: nodes[mid]}
    node.Left = buildBalancedBST(nodes, left, mid-1)
    node.Right = buildBalancedBST(nodes, mid+1, right)
    return node
}
```
