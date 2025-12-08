## 红黑树（Red-Black Tree）

### 定义

**红黑树**是一种自平衡的二叉搜索树，由 Rudolf Bayer 于1972年发明。它通过在每个节点上增加一个颜色属性（红色或黑色），并通过特定规则来保持树的近似平衡。

**红黑树的五条性质：**
1. 每个节点要么是红色，要么是黑色
2. 根节点是黑色
3. 每个叶子节点（NIL节点）是黑色
4. 如果一个节点是红色，则它的两个子节点都是黑色（不能有连续的红色节点）
5. 从任意节点到其所有后代叶子节点的路径上，黑色节点数目相同（黑高相同）

### 为什么使用红黑树？

红黑树在 AVL 树的基础上放宽了平衡条件，允许局部不平衡，从而减少了旋转操作的次数。这使得红黑树在频繁插入删除的场景下性能更优。

**应用场景：**
- Java 的 TreeMap、TreeSet
- C++ STL 的 map、set
- Linux 内核的进程调度（CFS）
- epoll 的文件描述符管理
- Nginx 的定时器管理

### 特性

| 操作 | 时间复杂度 |
|------|-----------|
| 查找 | O(log n) |
| 插入 | O(log n) |
| 删除 | O(log n) |
| 空间复杂度 | O(n) |

> 红黑树保证最长路径不超过最短路径的2倍，因此高度最大为 2log(n+1)。

### 红黑树与2-3树的等价性

理解红黑树的最佳方式是将其与2-3树对应起来：
- **2-节点** 对应黑色节点
- **3-节点** 对应一个黑色节点 + 一个红色左子节点

```
2-3树的3-节点:     红黑树表示:
    [a,b]              b(黑)
   / | \              /
  L  M  R           a(红)
                   /    \
                  L      M
```

### 数据结构

```go
const RED = true
const BLACK = false

// Node 表示红黑树节点
type Node struct {
    k     int   // 键
    v     int   // 值
    left  *Node // 左子节点
    right *Node // 右子节点
    color bool  // 颜色：true=红色，false=黑色
}

// Tree 表示红黑树
type Tree struct {
    size int
    root *Node
}

// 创建新节点（默认为红色）
func InitNode(k, v int) *Node {
    return &Node{
        k:     k,
        v:     v,
        left:  nil,
        right: nil,
        color: RED,  // 新节点默认红色，表示要与父节点"融合"
    }
}

// 判断节点是否为红色
func isRed(node *Node) bool {
    if node == nil {
        return BLACK  // 空节点视为黑色
    }
    return node.color
}
```

### 核心方法实现

#### 左旋转

当右子节点为红色时，需要左旋转：

```go
// 左旋转
//   node                     x
//  /   \     左旋转         /  \
// T1   x   --------->   node   T3
//     / \              /   \
//    T2 T3            T1   T2
func (t *Tree) leftRotate(node *Node) *Node {
    x := node.right
    node.right = x.left
    x.left = node
    x.color = node.color  // x继承node的颜色
    node.color = RED      // node变为红色
    return x
}
```

#### 右旋转

当左子节点和左孙节点都为红色时，需要右旋转：

```go
// 右旋转
//     node                   x
//    /   \     右旋转       /  \
//   x    T2   ------->    y   node
//  / \                       /  \
// y  T1                     T1  T2
func (t *Tree) rightRotate(node *Node) *Node {
    x := node.left
    node.left = x.right
    x.right = node
    x.color = node.color  // x继承node的颜色
    node.color = RED      // node变为红色
    return x
}
```

#### 颜色翻转

当左右子节点都为红色时，需要颜色翻转：

```go
// 颜色翻转（分解4-节点）
func (t *Tree) flipColors(node *Node) {
    node.color = RED                      // 父节点变红（向上传递）
    node.left.color = BLACK              // 左子节点变黑
    node.right.color = BLACK             // 右子节点变黑
}
```

#### 插入操作

```go
// 向红黑树中添加元素
func (t *Tree) Push(k, v int) {
    t.root = t.push(t.root, k, v)
    t.root.color = BLACK  // 根节点始终为黑色
}

func (t *Tree) push(node *Node, k, v int) *Node {
    // 递归终止：创建新的红色节点
    if node == nil {
        t.size++
        return InitNode(k, v)
    }

    // 递归插入
    if k > node.k {
        node.right = t.push(node.right, k, v)
    } else if k < node.k {
        node.left = t.push(node.left, k, v)
    } else {
        node.v = v  // 键已存在，更新值
    }

    // 维护红黑树性质（三种情况）

    // 情况1：右子节点红色，左子节点黑色 -> 左旋
    if isRed(node.right) && !isRed(node.left) {
        node = t.leftRotate(node)
    }

    // 情况2：左子节点红色，左孙节点也红色 -> 右旋
    if isRed(node.left) && isRed(node.left.left) {
        node = t.rightRotate(node)
    }

    // 情况3：左右子节点都红色 -> 颜色翻转
    if isRed(node.left) && isRed(node.right) {
        t.flipColors(node)
    }

    return node
}
```

#### 查找操作

```go
// 获取指定键的节点
func (t *Tree) getNode(node *Node, k int) *Node {
    if node == nil {
        return nil
    }
    if k == node.k {
        return node
    } else if k < node.k {
        return t.getNode(node.left, k)
    } else {
        return t.getNode(node.right, k)
    }
}

// 判断是否包含指定键
func (t *Tree) Contains(key int) bool {
    return t.getNode(t.root, key) != nil
}

// 获取指定键的值
func (t *Tree) GetValue(key int) *int {
    node := t.getNode(t.root, key)
    if node == nil {
        return nil
    }
    return &node.v
}

// 更新指定键的值
func (t *Tree) SetNewValue(key, value int) {
    node := t.getNode(t.root, key)
    if node == nil {
        panic("没有这个key值")
    }
    node.v = value
}
```

### 插入过程图解

```
插入顺序: 7, 3, 18, 10, 22, 8, 11

步骤1: 插入7（根节点变黑）
    7(黑)

步骤2: 插入3（红色左子节点）
    7(黑)
   /
  3(红)

步骤3: 插入18（红色右子节点）-> 颜色翻转
    7(黑)                    7(黑)
   /    \         翻转后     /    \
  3(红)  18(红)   ------>  3(黑)  18(黑)

步骤4: 插入10
       7(黑)
      /    \
    3(黑)  18(黑)
           /
         10(红)

... 继续维护平衡 ...
```

### 测试用例

```go
func TestRedBlackTree(t *testing.T) {
    tree := &Tree{}

    // 插入元素
    tree.Push(7, 70)
    tree.Push(3, 30)
    tree.Push(18, 180)
    tree.Push(10, 100)
    tree.Push(22, 220)
    tree.Push(8, 80)
    tree.Push(11, 110)

    fmt.Println("树的大小:", tree.Size())        // 7
    fmt.Println("包含键10:", tree.Contains(10))  // true
    fmt.Println("包含键15:", tree.Contains(15))  // false

    if val := tree.GetValue(10); val != nil {
        fmt.Println("键10的值:", *val)  // 100
    }

    // 更新值
    tree.SetNewValue(10, 1000)
    fmt.Println("更新后键10的值:", *tree.GetValue(10))  // 1000
}
```

测试结果：
```
=== RUN   TestRedBlackTree
树的大小: 7
包含键10: true
包含键15: false
键10的值: 100
更新后键10的值: 1000
--- PASS: TestRedBlackTree (0.00s)
PASS
```

### 运行方式

```bash
go test ./Red-Black
```

### 红黑树 vs AVL树 vs 跳表

| 特性 | 红黑树 | AVL树 | 跳表 |
|-----|-------|-------|-----|
| 平衡条件 | 宽松 | 严格 | 概率 |
| 查找效率 | O(log n) | O(log n) | O(log n) |
| 插入旋转次数 | 最多2次 | 最多O(log n)次 | 无旋转 |
| 删除旋转次数 | 最多3次 | 最多O(log n)次 | 无旋转 |
| 实现复杂度 | 中等 | 中等 | 简单 |
| 适用场景 | 插入删除频繁 | 查找为主 | 范围查询、并发 |

### LeetCode 实战

红黑树本身较少直接出现在 LeetCode 中，但其思想在平衡树相关题目中常用：

#### [面试题 04.06. 后继者](https://leetcode-cn.com/problems/successor-lcci/)

查找二叉搜索树中指定节点的后继节点：

```go
func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
    var successor *TreeNode
    for root != nil {
        if root.Val > p.Val {
            successor = root  // 可能的后继
            root = root.Left
        } else {
            root = root.Right
        }
    }
    return successor
}
```

#### [剑指 Offer 54. 二叉搜索树的第k大节点](https://leetcode-cn.com/problems/er-cha-sou-suo-shu-de-di-kda-jie-dian-lcof/)

```go
func kthLargest(root *TreeNode, k int) int {
    var result int
    var count int

    // 反向中序遍历：右 -> 根 -> 左
    var traverse func(node *TreeNode)
    traverse = func(node *TreeNode) {
        if node == nil || count >= k {
            return
        }
        traverse(node.Right)
        count++
        if count == k {
            result = node.Val
            return
        }
        traverse(node.Left)
    }

    traverse(root)
    return result
}
```
