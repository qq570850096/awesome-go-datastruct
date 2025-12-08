## 集合（Set）

### 定义

**集合**（Set）是一种不包含重复元素的数据结构，它主要支持以下操作：
- 添加元素
- 删除元素
- 判断元素是否存在

集合的核心特性是**元素唯一性**——同一个元素只能出现一次。

### 为什么使用集合？

集合适用于需要**去重**或**成员检测**的场景：

**应用场景：**
- 统计不重复的用户访问量（UV）
- 判断用户是否已注册
- 标签系统（一个物品的标签不重复）
- 词汇表、字典
- 图论中记录已访问的节点

### 特性

本模块提供两种底层实现，性能特性不同：

| 实现方式 | Add | Remove | Contains | 空间复杂度 |
|---------|-----|--------|----------|-----------|
| 链表实现 | O(n) | O(n) | O(n) | O(n) |
| BST实现 | O(log n)* | O(log n)* | O(log n)* | O(n) |

> *BST实现的时间复杂度为平均情况，最坏情况（退化为链表）为O(n)

### 数据结构

#### 集合接口定义

```go
// Set 定义集合接口
type Set interface {
    Add(e int)            // 添加元素
    Remove(e int)         // 删除元素
    Contains(e int) bool  // 判断是否包含元素
    GetSize() int         // 获取集合大小
    IsEmpty() bool        // 判断是否为空
}
```

#### 基于BST的实现

```go
import "algo/BinarySearch"

// BST 基于二分搜索树实现的集合
type BST struct {
    bst *BinarySearch.Tree
}

// 创建BST集合
func NewBSTSet() *BST {
    return &BST{
        bst: &BinarySearch.Tree{},
    }
}
```

#### 基于链表的实现

```go
import "algo/Linked"

// ListSet 基于链表实现的集合
type ListSet struct {
    list *Linked.List
}

// 创建链表集合
func NewListSet() *ListSet {
    return &ListSet{
        list: Linked.InitList(),
    }
}
```

### 核心方法实现

#### BST集合实现

```go
// 添加元素
// BST本身不允许重复元素，直接添加即可
func (s *BST) Add(e int) {
    s.bst.AddE(e)
}

// 删除元素
func (s *BST) Remove(e int) {
    s.bst.Remove(e)
}

// 判断是否包含元素
func (s *BST) Contains(e int) bool {
    return s.bst.Contains(e)
}

// 获取集合大小
func (s *BST) GetSize() int {
    return s.bst.Size()
}

// 判断是否为空
func (s *BST) IsEmpty() bool {
    return s.bst.IsEmpty()
}
```

#### 链表集合实现

```go
// 添加元素（需要先检查是否存在）
func (s *ListSet) Add(e int) {
    if !s.Contains(e) {
        s.list.AddFirst(e)  // 头插法
    }
}

// 删除元素
func (s *ListSet) Remove(e int) {
    s.list.RemoveElement(e)
}

// 判断是否包含元素（需要遍历链表）
func (s *ListSet) Contains(e int) bool {
    return s.list.Contains(e)
}

// 获取集合大小
func (s *ListSet) GetSize() int {
    return s.list.Size()
}

// 判断是否为空
func (s *ListSet) IsEmpty() bool {
    return s.list.Size() == 0
}
```

### 两种实现对比

```
链表集合结构:
head -> [5] -> [3] -> [8] -> [1] -> nil
添加元素2: 需要遍历检查是否存在 -> O(n)
查找元素8: 需要遍历链表 -> O(n)

BST集合结构:
        5
       / \
      3   8
     /
    1
添加元素2: 二分查找位置 -> O(log n)
查找元素8: 二分查找 -> O(log n)
```

### 测试用例

```go
func TestSet(t *testing.T) {
    // 测试BST集合
    bstSet := NewBSTSet()
    bstSet.Add(5)
    bstSet.Add(3)
    bstSet.Add(8)
    bstSet.Add(1)
    bstSet.Add(3)  // 重复元素，不会添加

    fmt.Println("BST集合大小:", bstSet.GetSize())     // 4
    fmt.Println("包含3:", bstSet.Contains(3))         // true
    fmt.Println("包含10:", bstSet.Contains(10))       // false

    bstSet.Remove(3)
    fmt.Println("删除3后包含3:", bstSet.Contains(3))  // false

    // 测试链表集合
    listSet := NewListSet()
    listSet.Add(5)
    listSet.Add(3)
    listSet.Add(8)
    listSet.Add(1)
    listSet.Add(3)  // 重复元素，不会添加

    fmt.Println("\n链表集合大小:", listSet.GetSize())  // 4
    fmt.Println("包含3:", listSet.Contains(3))        // true
}

// 性能对比测试
func BenchmarkSet(t *testing.T) {
    n := 10000

    // BST集合
    start := time.Now()
    bstSet := NewBSTSet()
    for i := 0; i < n; i++ {
        bstSet.Add(rand.Intn(n))
    }
    for i := 0; i < n; i++ {
        bstSet.Contains(rand.Intn(n))
    }
    fmt.Println("BST集合耗时:", time.Since(start))

    // 链表集合
    start = time.Now()
    listSet := NewListSet()
    for i := 0; i < n; i++ {
        listSet.Add(rand.Intn(n))
    }
    for i := 0; i < n; i++ {
        listSet.Contains(rand.Intn(n))
    }
    fmt.Println("链表集合耗时:", time.Since(start))
}
```

### 使用建议

| 场景 | 推荐实现 | 原因 |
|------|---------|------|
| 数据量小（<100） | 链表 | 实现简单，常数因子小 |
| 数据量大 | BST | O(log n) 性能优势明显 |
| 需要稳定性能 | AVL/红黑树 | BST可能退化 |
| 只需要判断存在性 | 哈希表 | O(1) 最快 |

### 运行方式

```bash
go test ./Set
```

### LeetCode 实战

#### [217. 存在重复元素](https://leetcode-cn.com/problems/contains-duplicate/)

判断数组中是否存在重复元素：

```go
func containsDuplicate(nums []int) bool {
    set := make(map[int]bool)
    for _, num := range nums {
        if set[num] {
            return true  // 发现重复
        }
        set[num] = true
    }
    return false
}
```

#### [349. 两个数组的交集](https://leetcode-cn.com/problems/intersection-of-two-arrays/)

返回两个数组的交集（结果唯一）：

```go
func intersection(nums1 []int, nums2 []int) []int {
    set1 := make(map[int]bool)
    for _, num := range nums1 {
        set1[num] = true
    }

    result := make(map[int]bool)
    for _, num := range nums2 {
        if set1[num] {
            result[num] = true
        }
    }

    // 转换为切片
    ans := make([]int, 0, len(result))
    for num := range result {
        ans = append(ans, num)
    }
    return ans
}
```

#### [804. 唯一摩尔斯密码词](https://leetcode-cn.com/problems/unique-morse-code-words/)

统计不同摩尔斯密码表示的数量：

```go
func uniqueMorseRepresentations(words []string) int {
    morse := []string{
        ".-","-...","-.-.","-..",".","..-.","--.","....","..",
        ".---","-.-",".-..","--","-.","---",".--.","--.-",".-.",
        "...","-","..-","...-",".--","-..-","-.--","--..",
    }

    set := make(map[string]bool)
    for _, word := range words {
        var code strings.Builder
        for _, c := range word {
            code.WriteString(morse[c-'a'])
        }
        set[code.String()] = true
    }
    return len(set)
}
```

#### [705. 设计哈希集合](https://leetcode-cn.com/problems/design-hashset/)

不使用内置哈希表实现集合：

```go
type MyHashSet struct {
    buckets [][]int
    size    int
}

func Constructor() MyHashSet {
    return MyHashSet{
        buckets: make([][]int, 1000),
        size:    1000,
    }
}

func (s *MyHashSet) hash(key int) int {
    return key % s.size
}

func (s *MyHashSet) Add(key int) {
    h := s.hash(key)
    for _, v := range s.buckets[h] {
        if v == key {
            return  // 已存在
        }
    }
    s.buckets[h] = append(s.buckets[h], key)
}

func (s *MyHashSet) Remove(key int) {
    h := s.hash(key)
    for i, v := range s.buckets[h] {
        if v == key {
            s.buckets[h] = append(s.buckets[h][:i], s.buckets[h][i+1:]...)
            return
        }
    }
}

func (s *MyHashSet) Contains(key int) bool {
    h := s.hash(key)
    for _, v := range s.buckets[h] {
        if v == key {
            return true
        }
    }
    return false
}
```
