# AVL 自平衡二叉搜索树

> 融合文档即代码：直接摘取关键结构与方法片段，便于对照实现。

## 场景与特性
- 高度平衡的二叉搜索树，平衡因子始终落在 [-1,1]。
- 查找/插入/删除均摊 O(log n)，适合频繁更新的有序集合。

## 核心数据结构
```go
type Node struct {
    E      int
    height int
    Left   *Node
    Right  *Node
}

type Tree struct {
    root *Node
    size int
}
```

## 关键方法（片段）
- 旋转与回平衡：递归返回前更新高度 + 选择 LL/LR/RL/RR。
```go
func rebalance(n *Node) *Node {
    updateHeight(n)
    bf := balanceFactor(n)        // 左高为正，右高为负
    if bf > 1 {                   // 左重
        if balanceFactor(n.Left) < 0 {
            n.Left = leftRotate(n.Left)  // LR
        }
        return rightRotate(n)            // LL
    }
    if bf < -1 {                  // 右重
        if balanceFactor(n.Right) > 0 {
            n.Right = rightRotate(n.Right) // RL
        }
        return leftRotate(n)              // RR
    }
    return n
}
```

- 插入/删除：递归落位后回溯 rebalance，删除时用后继替换。
```go
func (t *Tree) add(node *Node, e int) *Node {
    if node == nil { t.size++; return newNode(e) }
    if e < node.E { node.Left = t.add(node.Left, e) }
    if e > node.E { node.Right = t.add(node.Right, e) }
    return rebalance(node)
}

func (t *Tree) remove(node *Node, e int) *Node {
    // 命中节点后用右子树最小值替换，再删后继
    ...
    return rebalance(node)
}
```

- 诊断辅助：`InOrder` 验证 BST；`IsBalanced` 递归检测高度差。

## 快速使用
- 运行测试：`go test ./AVL`
- 调试旋转：在测试中插入 `[30,20,40,10,25,35,50]` 后删除 `25`，观察 `IsBalanced()`/`InOrder()`。
