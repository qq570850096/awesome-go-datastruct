## 二分搜索树（Binary Search Tree）

### 定义

**二分搜索树**（Binary Search Tree，简称BST）是一种特殊的二叉树，它满足以下性质：
1. 每个节点的左子树中所有节点的值都小于该节点的值
2. 每个节点的右子树中所有节点的值都大于该节点的值
3. 左右子树也分别是二分搜索树

这种有序性使得二分搜索树能够高效地进行查找、插入和删除操作。

### 为什么使用二分搜索树？

二分搜索树通过有序存储数据，实现了类似"二分查找"的高效查询。与数组的二分查找不同，BST是动态数据结构，可以高效地进行插入和删除操作。

**应用场景：**
- 实现动态有序集合
- 字典/映射的底层实现
- 数据库索引
- 文件系统目录结构
- 符号表

**生活类比：** 图书馆的图书按照编号排列，超市的商品按照类别分区——这些都是"有序存储"的思想，能够大大加快查找速度。

### 特性

| 操作 | 平均时间复杂度 | 最坏时间复杂度 |
|------|---------------|---------------|
| 查找 | O(log n) | O(n) |
| 插入 | O(log n) | O(n) |
| 删除 | O(log n) | O(n) |
| 遍历 | O(n) | O(n) |

> **注意**：最坏情况发生在树退化为链表时（如按顺序插入元素），此时所有操作退化为O(n)。

### 数据结构

```go
// Node 表示二分搜索树的节点
type Node struct {
    E     int   // 节点存储的元素值
    Left  *Node // 左子节点
    Right *Node // 右子节点
}

// Tree 表示二分搜索树
type Tree struct {
    root *Node // 根节点
    size int   // 树中元素个数
}

// 初始化节点
func InitNode(E int) *Node {
    return &Node{
        E:     E,
        Left:  nil,
        Right: nil,
    }
}
```

### 核心方法实现

#### 插入元素

```go
// 向二分搜索树添加元素（公有方法）
func (t *Tree) AddE(e int) {
    t.root = t.add(t.root, e)
}

// 向以node为根的二分搜索树中插入元素（递归实现）
func (t *Tree) add(node *Node, e int) *Node {
    // 递归终止条件：找到空位置，创建新节点
    if node == nil {
        t.size++
        return InitNode(e)
    }

    // 递归插入
    if e > node.E {
        node.Right = t.add(node.Right, e)
    } else if e < node.E {
        node.Left = t.add(node.Left, e)
    }
    // e == node.E 时不做操作（不允许重复元素）

    return node
}
```

#### 查找元素

```go
// 查找二分搜索树中是否包含元素e
func (t *Tree) Contains(e int) bool {
    return t.contains(t.root, e)
}

// 递归查找
func (t *Tree) contains(node *Node, e int) bool {
    if node == nil {
        return false
    }
    if e == node.E {
        return true
    } else if e > node.E {
        return t.contains(node.Right, e)
    } else {
        return t.contains(node.Left, e)
    }
}
```

#### 遍历操作

```go
// 前序遍历：根 -> 左 -> 右
func PreOrder(node *Node) {
    if node == nil {
        return
    }
    fmt.Printf("%d ", node.E)
    PreOrder(node.Left)
    PreOrder(node.Right)
}

// 非递归前序遍历（使用栈）
func (t *Tree) PreOrderNR() {
    stack := make([]*Node, 0)
    stack = append(stack, t.root)
    for len(stack) > 0 {
        curNode := stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        fmt.Printf("%d ", curNode.E)
        // 先压右子树，再压左子树（保证左子树先出栈）
        if curNode.Right != nil {
            stack = append(stack, curNode.Right)
        }
        if curNode.Left != nil {
            stack = append(stack, curNode.Left)
        }
    }
}

// 中序遍历：左 -> 根 -> 右（结果有序）
func MidOrder(node *Node) {
    if node == nil {
        return
    }
    MidOrder(node.Left)
    fmt.Printf("%d ", node.E)
    MidOrder(node.Right)
}

// 后序遍历：左 -> 右 -> 根
func BackOrder(node *Node) {
    if node == nil {
        return
    }
    BackOrder(node.Left)
    BackOrder(node.Right)
    fmt.Printf("%d ", node.E)
}

// 层序遍历（使用队列）
func (t *Tree) LevelOrder() {
    queue := make([]*Node, 0)
    queue = append(queue, t.root)
    for len(queue) > 0 {
        curNode := queue[0]
        queue = queue[1:]
        fmt.Printf("%d ", curNode.E)
        if curNode.Left != nil {
            queue = append(queue, curNode.Left)
        }
        if curNode.Right != nil {
            queue = append(queue, curNode.Right)
        }
    }
}
```

#### 查找最值

```go
// 查找最小值（最左节点）
func (t *Tree) FindMin() int {
    if t.IsEmpty() {
        panic("二叉树为空")
    }
    return minimum(t.root).E
}

func minimum(node *Node) *Node {
    if node.Left == nil {
        return node
    }
    return minimum(node.Left)
}

// 查找最大值（最右节点）
func (t *Tree) FindMax() int {
    if t.IsEmpty() {
        panic("二叉树为空")
    }
    return maximum(t.root).E
}

func maximum(node *Node) *Node {
    if node.Right == nil {
        return node
    }
    return maximum(node.Right)
}
```

#### 删除操作

删除节点时需要考虑三种情况：
1. 待删除节点只有左子树
2. 待删除节点只有右子树
3. 待删除节点左右子树都有（用后继节点替换）

```go
// 删除指定元素
func (t *Tree) Remove(e int) {
    t.root = t.remove(t.root, e)
}

func (t *Tree) remove(node *Node, e int) *Node {
    if node == nil {
        return nil
    }

    if e > node.E {
        node.Right = t.remove(node.Right, e)
        return node
    } else if e < node.E {
        node.Left = t.remove(node.Left, e)
        return node
    } else {
        // 找到待删除节点

        // 情况1：只有右子树
        if node.Left == nil {
            nodeRight := node.Right
            node.Right = nil
            t.size--
            return nodeRight
        }

        // 情况2：只有左子树
        if node.Right == nil {
            nodeLeft := node.Left
            node.Left = nil
            t.size--
            return nodeLeft
        }

        // 情况3：左右子树都存在
        // 找到后继节点（右子树中的最小值）
        nodeNext := minimum(node.Right)
        nodeNext.Right = t.rmMin(node.Right)
        nodeNext.Left = node.Left
        node.Left = nil
        node.Right = nil
        return nodeNext
    }
}

// 删除最小节点
func (t *Tree) rmMin(node *Node) *Node {
    if node.Left == nil {
        nodeRight := node.Right
        node.Right = nil
        t.size--
        return nodeRight
    }
    node.Left = t.rmMin(node.Left)
    return node
}
```

### 测试用例

```go
func TestBST(t *testing.T) {
    tree := &Tree{}

    // 插入元素
    elements := []int{47, 35, 60, 22, 37, 55, 67}
    for _, e := range elements {
        tree.AddE(e)
    }

    fmt.Println("前序遍历:")
    tree.PreOrder() // 47 35 22 37 60 55 67

    fmt.Println("中序遍历（有序）:")
    tree.MidOrder() // 22 35 37 47 55 60 67

    fmt.Println("层序遍历:")
    tree.LevelOrder() // 47 35 60 22 37 55 67

    fmt.Println("最小值:", tree.FindMin()) // 22
    fmt.Println("最大值:", tree.FindMax()) // 67

    // 删除节点
    tree.Remove(47)
    fmt.Println("删除47后中序遍历:")
    tree.MidOrder() // 22 35 37 55 60 67
}
```

### 运行方式

```bash
go test ./BinarySearch
```

### LeetCode 实战

#### [98. 验证二叉搜索树](https://leetcode-cn.com/problems/validate-binary-search-tree/)

验证一棵树是否为有效的二分搜索树：

```go
func isValidBST(root *TreeNode) bool {
    return validate(root, math.MinInt64, math.MaxInt64)
}

func validate(node *TreeNode, min, max int64) bool {
    if node == nil {
        return true
    }
    val := int64(node.Val)
    if val <= min || val >= max {
        return false
    }
    return validate(node.Left, min, val) && validate(node.Right, val, max)
}
```

#### [700. 二叉搜索树中的搜索](https://leetcode-cn.com/problems/search-in-a-binary-search-tree/)

在BST中搜索给定值：

```go
func searchBST(root *TreeNode, val int) *TreeNode {
    if root == nil || root.Val == val {
        return root
    }
    if val < root.Val {
        return searchBST(root.Left, val)
    }
    return searchBST(root.Right, val)
}
```

#### [701. 二叉搜索树中的插入操作](https://leetcode-cn.com/problems/insert-into-a-binary-search-tree/)

向BST中插入新值：

```go
func insertIntoBST(root *TreeNode, val int) *TreeNode {
    if root == nil {
        return &TreeNode{Val: val}
    }
    if val < root.Val {
        root.Left = insertIntoBST(root.Left, val)
    } else {
        root.Right = insertIntoBST(root.Right, val)
    }
    return root
}
```

#### [450. 删除二叉搜索树中的节点](https://leetcode-cn.com/problems/delete-node-in-a-bst/)

从BST中删除给定值的节点：

```go
func deleteNode(root *TreeNode, key int) *TreeNode {
    if root == nil {
        return nil
    }
    if key < root.Val {
        root.Left = deleteNode(root.Left, key)
    } else if key > root.Val {
        root.Right = deleteNode(root.Right, key)
    } else {
        // 找到待删除节点
        if root.Left == nil {
            return root.Right
        }
        if root.Right == nil {
            return root.Left
        }
        // 找到后继节点（右子树最小值）
        minNode := root.Right
        for minNode.Left != nil {
            minNode = minNode.Left
        }
        root.Val = minNode.Val
        root.Right = deleteNode(root.Right, minNode.Val)
    }
    return root
}
```
