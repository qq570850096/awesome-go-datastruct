## 单向链表（Singly Linked List）

### 定义

**单向链表**（Singly Linked List）是一种线性数据结构，由若干个节点组成，每个节点包含两部分：数据域和指向下一个节点的指针。链表的元素在内存中不必连续存储，而是通过指针链接在一起。

单向链表的核心特点：
1. **动态大小**：无需预先指定长度
2. **单向遍历**：只能从头向尾遍历
3. **高效插入删除**：O(1) 时间（已知前驱节点）

### 为什么使用链表？

**链表 vs 数组的对比：**

| 特性 | 链表 | 数组 |
|------|------|------|
| 内存分配 | 动态，按需分配 | 静态，预先分配 |
| 插入/删除 | O(1)（已知位置） | O(n)（需移动元素） |
| 随机访问 | O(n) | O(1) |
| 内存利用 | 有指针开销 | 紧凑存储 |
| 缓存友好 | 较差 | 较好 |

**应用场景：**
- 实现栈、队列、双端队列
- 内存管理（空闲块链表）
- 多项式运算
- 稀疏矩阵表示
- 浏览器历史记录

### 特性

| 操作 | 时间复杂度 | 说明 |
|------|-----------|------|
| 头部插入 | O(1) | AddFirst |
| 尾部插入 | O(n) | AddLast（无尾指针） |
| 任意位置插入 | O(n) | 需先找到位置 |
| 头部删除 | O(1) | RemoveFirst |
| 任意位置删除 | O(n) | 需先找到位置 |
| 查找 | O(n) | 顺序遍历 |
| 空间复杂度 | O(n) | n 为节点数 |

### 数据结构

```go
// Node 链表节点
type Node struct {
    E    int   // 数据域
    Next *Node // 指向下一个节点的指针
}

// List 单向链表（使用虚拟头节点）
type List struct {
    dummyHead *Node // 虚拟头节点（哨兵节点）
    size      int   // 链表长度
}

// 初始化节点
func initNode(e int) *Node {
    return &Node{
        E:    e,
        Next: nil,
    }
}

// InitList 创建空链表
func InitList() *List {
    return &List{
        dummyHead: initNode(0), // 虚拟头节点不存储实际数据
        size:      0,
    }
}

// 获取链表大小
func (l List) Size() int {
    return l.size
}

// 判断链表是否为空
func (this *List) IsEmpty() bool {
    return this.size == 0
}
```

### 核心方法实现

#### 插入操作

```go
// AddIndex 在指定位置插入元素
// index 从 0 开始，表示插入到第 index 个位置
func (this *List) AddIndex(index, e int) {
    if index > this.size || index < 0 {
        panic("索引越界，不能插入")
    }

    prev := this.dummyHead
    node := initNode(e)

    // 找到第 index 个位置的前一个节点
    for i := 0; i < index; i++ {
        prev = prev.Next
    }

    // 插入新节点
    node.Next = prev.Next
    prev.Next = node
    this.size++
}

// AddFirst 在链表头部插入元素 - O(1)
func (this *List) AddFirst(e int) {
    this.AddIndex(0, e)
}

// AddLast 在链表尾部插入元素 - O(n)
func (this *List) AddLast(e int) {
    this.AddIndex(this.size, e)
}
```

#### 查询操作

```go
// Get 获取第 index 个元素
func (this *List) Get(index int) int {
    if index >= this.size || index < 0 {
        panic("索引越界，不能查询")
    }

    cur := this.dummyHead.Next
    for i := 0; i < index; i++ {
        cur = cur.Next
    }
    return cur.E
}

// GetFirst 获取第一个元素
func (this *List) GetFirst() int {
    return this.Get(0)
}

// GetLast 获取最后一个元素
func (this *List) GetLast() int {
    return this.Get(this.size - 1)
}

// Contains 判断链表是否包含元素 e
func (this *List) Contains(e int) bool {
    cur := this.dummyHead.Next
    for cur != nil {
        if cur.E == e {
            return true
        }
        cur = cur.Next
    }
    return false
}
```

#### 删除操作

```go
// Remove 删除第 index 个元素，返回被删除的元素值
func (this *List) Remove(index int) int {
    if index >= this.size || index < 0 {
        panic("索引越界，不能删除")
    }

    prev := this.dummyHead
    for i := 0; i < index; i++ {
        prev = prev.Next
    }

    retNode := prev.Next
    prev.Next = retNode.Next
    retNode.Next = nil // 帮助 GC
    this.size--

    return retNode.E
}

// RemoveFirst 删除第一个元素
func (this *List) RemoveFirst() int {
    return this.Remove(0)
}

// RemoveLast 删除最后一个元素
func (this *List) RemoveLast() int {
    return this.Remove(this.size - 1)
}

// RemoveElement 删除指定值的第一个元素
func (this *List) RemoveElement(e int) {
    prev := this.dummyHead
    for prev.Next != nil {
        if prev.Next.E == e {
            break
        }
        prev = prev.Next
    }
    if prev.Next != nil {
        delNode := prev.Next
        prev.Next = delNode.Next
        delNode.Next = nil
        this.size--
    }
}
```

#### 修改操作

```go
// Set 修改第 index 个元素的值
func (this *List) Set(index, e int) {
    if index >= this.size || index < 0 {
        panic("索引越界，不能修改")
    }

    cur := this.dummyHead.Next
    for i := 0; i < index; i++ {
        cur = cur.Next
    }
    cur.E = e
}
```

#### 链表排序

```go
// Sort 链表快速排序
func (this *List) Sort() {
    if this.dummyHead == nil || this.dummyHead.Next == nil {
        return
    }
    qsortList(this.dummyHead, nil)
}

func qsortList(head, tail *Node) {
    // 链表范围是 [head, tail)
    if head != tail && head.Next != tail {
        mid := partitionList(head, tail)
        qsortList(head, mid)
        qsortList(mid.Next, tail)
    }
}

func partitionList(head, tail *Node) *Node {
    key := head.E
    loc := head
    for i := head.Next; i != tail; i = i.Next {
        if i.E < key {
            loc = loc.Next
            i.E, loc.E = loc.E, i.E
        }
    }
    loc.E, head.E = head.E, loc.E
    return loc
}
```

### 链表操作示意图

```
虚拟头节点的作用：统一插入/删除操作，无需特殊处理头节点

初始状态：
dummyHead -> nil

AddFirst(1)：
dummyHead -> [1] -> nil

AddLast(2), AddLast(3)：
dummyHead -> [1] -> [2] -> [3] -> nil

AddIndex(1, 5)：在位置 1 插入 5
dummyHead -> [1] -> [5] -> [2] -> [3] -> nil

Remove(1)：删除位置 1 的元素
dummyHead -> [1] -> [2] -> [3] -> nil
```

### 测试用例

```go
func TestLinkedList(t *testing.T) {
    list := InitList()

    // 测试插入
    list.AddFirst(1)
    list.AddLast(3)
    list.AddIndex(1, 2)
    fmt.Println(list.String()) // 1 -> 2 -> 3 -> NULL

    // 测试查询
    fmt.Println("第 1 个元素:", list.Get(1))     // 2
    fmt.Println("包含 2:", list.Contains(2))    // true
    fmt.Println("包含 5:", list.Contains(5))    // false

    // 测试修改
    list.Set(1, 5)
    fmt.Println(list.String()) // 1 -> 5 -> 3 -> NULL

    // 测试删除
    list.Remove(1)
    fmt.Println(list.String()) // 1 -> 3 -> NULL

    // 测试排序
    list.AddFirst(5)
    list.AddLast(2)
    fmt.Println("排序前:", list.String()) // 5 -> 1 -> 3 -> 2 -> NULL
    list.Sort()
    fmt.Println("排序后:", list.String()) // 1 -> 2 -> 3 -> 5 -> NULL
}
```

### 运行方式

```bash
go test ./Linked
```

### LeetCode 实战

#### [206. 反转链表](https://leetcode-cn.com/problems/reverse-linked-list/)

反转单链表：

```go
// 迭代法
func reverseList(head *ListNode) *ListNode {
    var prev *ListNode = nil
    curr := head

    for curr != nil {
        next := curr.Next // 保存下一个节点
        curr.Next = prev  // 反转指针
        prev = curr       // prev 前进
        curr = next       // curr 前进
    }

    return prev
}

// 递归法
func reverseListRecursive(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    newHead := reverseListRecursive(head.Next)
    head.Next.Next = head
    head.Next = nil
    return newHead
}
```

#### [21. 合并两个有序链表](https://leetcode-cn.com/problems/merge-two-sorted-lists/)

将两个有序链表合并为一个有序链表：

```go
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    dummy := &ListNode{} // 虚拟头节点
    curr := dummy

    for list1 != nil && list2 != nil {
        if list1.Val <= list2.Val {
            curr.Next = list1
            list1 = list1.Next
        } else {
            curr.Next = list2
            list2 = list2.Next
        }
        curr = curr.Next
    }

    // 拼接剩余部分
    if list1 != nil {
        curr.Next = list1
    } else {
        curr.Next = list2
    }

    return dummy.Next
}
```

#### [141. 环形链表](https://leetcode-cn.com/problems/linked-list-cycle/)

判断链表是否有环（快慢指针）：

```go
func hasCycle(head *ListNode) bool {
    if head == nil || head.Next == nil {
        return false
    }

    slow, fast := head, head.Next
    for slow != fast {
        if fast == nil || fast.Next == nil {
            return false
        }
        slow = slow.Next
        fast = fast.Next.Next
    }
    return true
}
```

#### [142. 环形链表 II](https://leetcode-cn.com/problems/linked-list-cycle-ii/)

找出环的入口节点：

```go
func detectCycle(head *ListNode) *ListNode {
    slow, fast := head, head

    // 第一阶段：判断是否有环
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
        if slow == fast {
            // 第二阶段：找入口
            slow = head
            for slow != fast {
                slow = slow.Next
                fast = fast.Next
            }
            return slow
        }
    }
    return nil
}
```

#### [19. 删除链表的倒数第 N 个节点](https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list/)

使用双指针一次遍历：

```go
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    dummy := &ListNode{Next: head}
    fast, slow := dummy, dummy

    // fast 先走 n+1 步
    for i := 0; i <= n; i++ {
        fast = fast.Next
    }

    // 同时移动，直到 fast 到达末尾
    for fast != nil {
        fast = fast.Next
        slow = slow.Next
    }

    // 删除 slow.Next
    slow.Next = slow.Next.Next
    return dummy.Next
}
```

#### [234. 回文链表](https://leetcode-cn.com/problems/palindrome-linked-list/)

判断链表是否为回文：

```go
func isPalindrome(head *ListNode) bool {
    if head == nil || head.Next == nil {
        return true
    }

    // 找中点
    slow, fast := head, head
    for fast.Next != nil && fast.Next.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }

    // 反转后半部分
    secondHalf := reverseList(slow.Next)

    // 比较两半
    p1, p2 := head, secondHalf
    for p2 != nil {
        if p1.Val != p2.Val {
            return false
        }
        p1 = p1.Next
        p2 = p2.Next
    }

    return true
}
```
