## 跳表（Skip List）

### 定义

**跳表**（Skip List）是由 William Pugh 于1990年发明的一种概率性数据结构，它通过在有序链表的基础上增加多级索引，实现了类似平衡树的查找效率，但实现更加简单。

跳表的核心思想：通过"空间换时间"，在链表上建立多层"快速通道"，使得查找可以跳跃式前进，而不必逐个遍历。

### 为什么使用跳表？

**与平衡树相比的优势：**
- 实现简单，代码量少
- 插入删除时不需要复杂的旋转操作
- 支持范围查询，天然有序
- 并发友好，锁粒度更细

**应用场景：**
- Redis 的有序集合（Sorted Set）底层实现
- LevelDB / RocksDB 的 MemTable
- 需要有序数据且频繁插入删除的场景

### 特性

| 操作 | 平均时间复杂度 | 最坏时间复杂度 |
|------|---------------|---------------|
| 查找 | O(log n) | O(n) |
| 插入 | O(log n) | O(n) |
| 删除 | O(log n) | O(n) |
| 空间复杂度 | O(n) | O(n log n) |

> 跳表通过概率性方式实现平衡，期望高度为 O(log n)。

### 数据结构

```go
const (
    maxLevel int     = 16   // 最大层数，足够应对 2^16 个元素
    p        float32 = 0.25 // 层数增长概率
)

// Node 表示跳表中的节点
type Node struct {
    Score   float64       // 排序分值（作为排序键）
    Value   interface{}   // 存储的数据
    forward []*Node       // 每层的前进指针
}

// SkipList 表示一份跳表
type SkipList struct {
    header *Node // 头节点（哨兵节点，不存储实际数据）
    len    int   // 当前跳表长度（不含头节点）
    level  int   // 当前跳表的最高层数
}

// 创建新节点
func newElement(score float64, value interface{}, level int) *Node {
    return &Node{
        Score:   score,
        Value:   value,
        forward: make([]*Node, level),
    }
}

// 创建空跳表
func New() *SkipList {
    return &SkipList{
        header: &Node{forward: make([]*Node, maxLevel)},
    }
}
```

### 核心方法实现

#### 随机层高

跳表的关键在于随机决定新节点的层高，这保证了跳表的概率平衡性：

```go
// 随机生成层高
// 以概率 p 向上增长，期望层高为 1/(1-p)
func randomLevel() int {
    level := 1
    for rand.Float32() < p && level < maxLevel {
        level++
    }
    return level
}
```

#### 查找操作

从最高层开始，逐层向下查找：

```go
// 在跳表中查找给定分值对应的节点
// 返回 (*Node, true) 表示找到，(nil, false) 表示未找到
func (sl *SkipList) Search(score float64) (element *Node, ok bool) {
    x := sl.header

    // 从最高层开始向下查找
    for i := sl.level - 1; i >= 0; i-- {
        // 在当前层向右移动，直到下一个节点的分值 >= 目标分值
        for x.forward[i] != nil && x.forward[i].Score < score {
            x = x.forward[i]
        }
    }

    // 移动到第0层的下一个节点
    x = x.forward[0]

    // 检查是否命中
    if x != nil && x.Score == score {
        return x, true
    }
    return nil, false
}
```

#### 插入操作

```go
// 将分值和数据插入跳表，返回对应的节点指针
func (sl *SkipList) Insert(score float64, value interface{}) *Node {
    // update数组记录每层需要更新的前驱节点
    update := make([]*Node, maxLevel)
    x := sl.header

    // 从最高层向下，记录每层的前驱节点
    for i := sl.level - 1; i >= 0; i-- {
        for x.forward[i] != nil && x.forward[i].Score < score {
            x = x.forward[i]
        }
        update[i] = x
    }
    x = x.forward[0]

    // 如果分值已存在，直接更新数据
    if x != nil && x.Score == score {
        x.Value = value
        return x
    }

    // 随机生成新节点的层高
    level := randomLevel()

    // 如果新层高超过当前最高层，需要更新header的前驱
    if level > sl.level {
        level = sl.level + 1
        update[sl.level] = sl.header
        sl.level = level
    }

    // 创建新节点
    e := newElement(score, value, level)

    // 逐层插入新节点
    for i := 0; i < level; i++ {
        e.forward[i] = update[i].forward[i]
        update[i].forward[i] = e
    }

    sl.len++
    return e
}
```

#### 删除操作

```go
// 删除并返回给定分值的节点；不存在则返回 nil
func (sl *SkipList) Delete(score float64) *Node {
    update := make([]*Node, maxLevel)
    x := sl.header

    // 记录每层的前驱节点
    for i := sl.level - 1; i >= 0; i-- {
        for x.forward[i] != nil && x.forward[i].Score < score {
            x = x.forward[i]
        }
        update[i] = x
    }
    x = x.forward[0]

    // 未找到目标节点
    if x == nil || x.Score != score {
        return nil
    }

    // 逐层删除节点
    for i := 0; i < len(x.forward); i++ {
        if update[i].forward[i] != x {
            break
        }
        update[i].forward[i] = x.forward[i]
    }

    // 更新跳表高度（如果最高层变空）
    for sl.level > 1 && sl.header.forward[sl.level-1] == nil {
        sl.level--
    }

    sl.len--
    return x
}
```

#### 辅助方法

```go
// 返回跳表长度
func (sl *SkipList) Size() int {
    return sl.len
}

// 返回第一个节点
func (sl *SkipList) Front() *Node {
    return sl.header.forward[0]
}

// 返回当前节点的下一个节点
func (e *Node) Next() *Node {
    if e != nil {
        return e.forward[0]
    }
    return nil
}
```

### 跳表结构示意图

```
Level 3:  head ────────────────────────────> 67 ──────────> nil
Level 2:  head ──────────> 35 ─────────────> 67 ──────────> nil
Level 1:  head ────> 22 -> 35 -> 47 ──────> 67 -> 78 ────> nil
Level 0:  head -> 12 -> 22 -> 35 -> 47 -> 55 -> 67 -> 78 -> 89 -> nil
```

### 测试用例

```go
func TestSkipList(t *testing.T) {
    sl := New()

    // 插入元素
    sl.Insert(3.0, "three")
    sl.Insert(1.0, "one")
    sl.Insert(2.0, "two")
    sl.Insert(4.0, "four")

    fmt.Println("跳表长度:", sl.Size()) // 4

    // 查找元素
    if node, ok := sl.Search(2.0); ok {
        fmt.Printf("找到: score=%.1f, value=%v\n", node.Score, node.Value)
        // 找到: score=2.0, value=two
    }

    // 遍历跳表（有序）
    fmt.Println("有序遍历:")
    for node := sl.Front(); node != nil; node = node.Next() {
        fmt.Printf("  %.1f: %v\n", node.Score, node.Value)
    }
    // 1.0: one
    // 2.0: two
    // 3.0: three
    // 4.0: four

    // 删除元素
    sl.Delete(2.0)
    fmt.Println("删除2.0后长度:", sl.Size()) // 3
}
```

测试结果：
```
=== RUN   TestSkipList
跳表长度: 4
找到: score=2.0, value=two
有序遍历:
  1.0: one
  2.0: two
  3.0: three
  4.0: four
删除2.0后长度: 3
--- PASS: TestSkipList (0.00s)
PASS
```

### 运行方式

```bash
go test ./skiplists
```

### LeetCode 实战

#### [1206. 设计跳表](https://leetcode-cn.com/problems/design-skiplist/)

实现一个跳表类，支持搜索、添加、删除操作：

```go
type Skiplist struct {
    head  *Node
    level int
}

type Node struct {
    val     int
    forward []*Node
}

const maxLevel = 16
const p = 0.5

func Constructor() Skiplist {
    return Skiplist{
        head: &Node{forward: make([]*Node, maxLevel)},
    }
}

func (sl *Skiplist) Search(target int) bool {
    cur := sl.head
    for i := sl.level - 1; i >= 0; i-- {
        for cur.forward[i] != nil && cur.forward[i].val < target {
            cur = cur.forward[i]
        }
    }
    cur = cur.forward[0]
    return cur != nil && cur.val == target
}

func (sl *Skiplist) Add(num int) {
    update := make([]*Node, maxLevel)
    cur := sl.head
    for i := sl.level - 1; i >= 0; i-- {
        for cur.forward[i] != nil && cur.forward[i].val < num {
            cur = cur.forward[i]
        }
        update[i] = cur
    }

    level := randomLevel()
    if level > sl.level {
        for i := sl.level; i < level; i++ {
            update[i] = sl.head
        }
        sl.level = level
    }

    node := &Node{val: num, forward: make([]*Node, level)}
    for i := 0; i < level; i++ {
        node.forward[i] = update[i].forward[i]
        update[i].forward[i] = node
    }
}

func (sl *Skiplist) Erase(num int) bool {
    update := make([]*Node, maxLevel)
    cur := sl.head
    for i := sl.level - 1; i >= 0; i-- {
        for cur.forward[i] != nil && cur.forward[i].val < num {
            cur = cur.forward[i]
        }
        update[i] = cur
    }
    cur = cur.forward[0]

    if cur == nil || cur.val != num {
        return false
    }

    for i := 0; i < sl.level && update[i].forward[i] == cur; i++ {
        update[i].forward[i] = cur.forward[i]
    }
    for sl.level > 1 && sl.head.forward[sl.level-1] == nil {
        sl.level--
    }
    return true
}

func randomLevel() int {
    level := 1
    for rand.Float64() < p && level < maxLevel {
        level++
    }
    return level
}
```
