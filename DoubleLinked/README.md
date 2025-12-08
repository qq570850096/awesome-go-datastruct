## 双向链表与缓存策略（Double Linked List & Cache）

### 定义

**双向链表**（Double Linked List）是一种链式数据结构，每个节点包含三个部分：数据、指向前驱节点的指针、指向后继节点的指针。相比单向链表，双向链表可以双向遍历，支持 O(1) 时间复杂度的任意位置删除。

本模块基于双向链表实现了三种常用的缓存淘汰策略：
- **LRU**（Least Recently Used）：最近最少使用
- **LFU**（Least Frequently Used）：最不经常使用
- **FIFO**（First In First Out）：先进先出

### 为什么使用双向链表实现缓存？

**双向链表的优势：**
- O(1) 时间删除任意已知节点（无需遍历找前驱）
- O(1) 时间在头部或尾部插入
- 配合哈希表可实现 O(1) 的缓存查找和更新

**应用场景：**
- 操作系统页面置换算法
- 数据库缓冲池管理
- Redis 内存淘汰策略
- Web 服务器缓存
- CDN 缓存策略

### 特性

| 操作 | 时间复杂度 | 说明 |
|------|-----------|------|
| 头部插入 | O(1) | addHead |
| 尾部插入 | O(1) | addTail |
| 删除节点 | O(1) | 已知节点指针 |
| 查找节点 | O(1) | 配合哈希表 |
| 空间复杂度 | O(n) | n 为节点数量 |

### 数据结构

#### 双向链表节点

```go
// Node 双向链表节点
type Node struct {
    key   interface{} // 键（用于缓存淘汰时从哈希表删除）
    value interface{} // 值
    prev  *Node       // 前驱指针
    next  *Node       // 后继指针
}

func InitNode(key, value interface{}) *Node {
    return &Node{
        key:   key,
        value: value,
    }
}
```

#### 双向链表

```go
// List 双向链表
type List struct {
    capacity int   // 容量限制
    head     *Node // 头节点
    tail     *Node // 尾节点
    size     int   // 当前大小
}

func InitList(capacity int) *List {
    return &List{
        capacity: capacity,
        size:     0,
    }
}
```

### 核心方法实现

#### 头部插入

```go
// addHead 将节点插入到链表头部
func (this *List) addHead(node *Node) *Node {
    if this.head == nil {
        // 空链表
        this.head = node
        this.tail = node
        this.head.prev = nil
        this.tail.next = nil
    } else {
        node.next = this.head
        this.head.prev = node
        this.head = node
        this.head.prev = nil
    }
    this.size++
    return node
}
```

#### 尾部插入

```go
// addTail 将节点插入到链表尾部
func (this *List) addTail(node *Node) *Node {
    if this.tail == nil {
        this.tail = node
        this.head = node
        this.head.prev = nil
        this.tail.next = nil
    } else {
        this.tail.next = node
        node.prev = this.tail
        this.tail = node
        this.tail.next = nil
    }
    this.size++
    return node
}
```

#### 删除节点

```go
// remove 删除指定节点
// 如果 node == nil，默认删除尾节点
func (this *List) remove(node *Node) *Node {
    if node == nil {
        node = this.tail
    }
    if node == this.tail {
        this.removeTail()
    } else if node == this.head {
        this.removeHead()
    } else {
        // 中间节点：更新前后节点的指针
        node.next.prev = node.prev
        node.prev.next = node.next
        this.size--
    }
    return node
}

// removeTail 删除尾节点
func (this *List) removeTail() *Node {
    if this.tail == nil {
        return nil
    }
    node := this.tail
    if node.prev != nil {
        this.tail = node.prev
        this.tail.next = nil
    } else {
        // 只有一个节点
        this.tail = nil
        this.head = nil
    }
    this.size--
    return node
}
```

### LRU 缓存实现

LRU（最近最少使用）策略：每次访问将节点移到头部，淘汰时删除尾部节点。

```go
// LRUCache LRU 缓存
type LRUCache struct {
    capacity int                      // 容量
    find     map[interface{}]*Node    // 哈希表：key -> 节点
    list     *List                    // 双向链表
}

func InitLRU(capacity int) *LRUCache {
    return &LRUCache{
        capacity: capacity,
        list:     InitList(capacity),
        find:     make(map[interface{}]*Node),
    }
}

// Get 获取缓存
// 命中则将节点移到头部，未命中返回 -1
func (this *LRUCache) Get(key interface{}) interface{} {
    if node, ok := this.find[key]; ok {
        // 命中：移到头部
        this.list.Remove(node)
        this.list.AppendToHead(node)
        return node.value
    }
    return -1
}

// Put 写入缓存
func (this *LRUCache) Put(key, value interface{}) {
    if node, ok := this.find[key]; ok {
        // key 已存在：更新值并移到头部
        this.list.Remove(node)
        node.value = value
        this.list.AppendToHead(node)
    } else {
        // key 不存在：创建新节点
        node := InitNode(key, value)
        // 检查容量
        if this.list.size >= this.list.capacity {
            // 淘汰尾部节点
            oldNode := this.list.Remove(nil)
            delete(this.find, oldNode.key)
        }
        this.list.AppendToHead(node)
        this.find[key] = node
    }
}
```

### LFU 缓存实现

LFU（最不经常使用）策略：按访问频率管理，淘汰访问次数最少的节点。

```go
// LFUNode LFU 节点（包含频率信息）
type LFUNode struct {
    freq int   // 访问频率
    node *Node // 链表节点
}

// LFUCache LFU 缓存
type LFUCache struct {
    capacity int                       // 容量
    find     map[interface{}]*LFUNode  // key -> LFUNode
    freq_map map[int]*List             // 频率 -> 链表
    size     int                       // 当前大小
}

func InitLFUCache(capacity int) *LFUCache {
    return &LFUCache{
        capacity: capacity,
        find:     map[interface{}]*LFUNode{},
        freq_map: map[int]*List{},
    }
}

// updateFreq 更新节点频率
func (this *LFUCache) updateFreq(node *LFUNode) {
    freq := node.freq
    // 从当前频率链表删除
    node.node = this.freq_map[freq].Remove(node.node)
    if this.freq_map[freq].size == 0 {
        delete(this.freq_map, freq)
    }
    // 频率+1，加入新链表
    freq++
    node.freq = freq
    if _, ok := this.freq_map[freq]; !ok {
        this.freq_map[freq] = InitList(math.MaxInt32)
    }
    this.freq_map[freq].Append(node.node)
}

// Get 获取缓存
func (this *LFUCache) Get(key interface{}) interface{} {
    node, ok := this.find[key]
    if !ok {
        return -1
    }
    this.updateFreq(node)
    return node.node.value
}

// Put 写入缓存
func (this *LFUCache) Put(key, value interface{}) {
    if this.capacity == 0 {
        return
    }
    if _, ok := this.find[key]; ok {
        // 更新已存在的 key
        node := this.find[key]
        node.node.value = value
        this.updateFreq(node)
    } else {
        // 新增 key
        if this.capacity == this.size {
            // 找到最小频率并淘汰
            minFreq := findMinFreq(this.freq_map)
            list := this.freq_map[minFreq]
            evicted := list.Pop()
            if evicted != nil {
                delete(this.find, evicted.key)
                if list.size == 0 {
                    delete(this.freq_map, minFreq)
                }
                this.size--
            }
        }
        // 创建新节点
        node := InitLFUNode(key, value)
        node.freq = 1
        this.find[key] = node
        if _, ok := this.freq_map[node.freq]; !ok {
            this.freq_map[node.freq] = InitList(math.MaxInt32)
        }
        this.freq_map[node.freq].Append(node.node)
        this.size++
    }
}
```

### 缓存策略对比

| 策略 | 淘汰依据 | 优点 | 缺点 |
|------|---------|------|------|
| LRU | 最近访问时间 | 实现简单，局部性好 | 偶发访问可能污染缓存 |
| LFU | 访问频率 | 保留热点数据 | 新数据难以晋升 |
| FIFO | 进入顺序 | 实现最简单 | 不考虑访问模式 |

### 双向链表操作示意图

```
初始状态（空链表）：
head -> nil
tail -> nil

AppendToHead(A)：
head -> [A] <- tail

AppendToHead(B)：
head -> [B] <-> [A] <- tail

AppendToTail(C)：
head -> [B] <-> [A] <-> [C] <- tail

Remove(A)：
head -> [B] <-> [C] <- tail
```

### 测试用例

```go
func TestLRUCache(t *testing.T) {
    lru := InitLRU(2)

    lru.Put("a", 1)
    lru.Put("b", 2)
    fmt.Println(lru.Get("a")) // 1（a 移到头部）

    lru.Put("c", 3) // 淘汰 b（最近最少使用）
    fmt.Println(lru.Get("b")) // -1（已被淘汰）
    fmt.Println(lru.Get("c")) // 3

    fmt.Println("链表状态:", lru.String())
}

func TestLFUCache(t *testing.T) {
    lfu := InitLFUCache(2)

    lfu.Put("a", 1)
    lfu.Put("b", 2)
    lfu.Get("a") // a 频率变为 2

    lfu.Put("c", 3) // 淘汰 b（频率最低）
    fmt.Println(lfu.Get("b")) // -1
    fmt.Println(lfu.Get("a")) // 1
    fmt.Println(lfu.Get("c")) // 3
}
```

### 运行方式

```bash
go test ./DoubleLinked
```

### LeetCode 实战

#### [146. LRU 缓存](https://leetcode-cn.com/problems/lru-cache/)

实现 LRU 缓存机制：

```go
type LRUCache struct {
    capacity int
    cache    map[int]*DLinkedNode
    head     *DLinkedNode // 虚拟头节点
    tail     *DLinkedNode // 虚拟尾节点
}

type DLinkedNode struct {
    key, value int
    prev, next *DLinkedNode
}

func Constructor(capacity int) LRUCache {
    l := LRUCache{
        capacity: capacity,
        cache:    make(map[int]*DLinkedNode),
        head:     &DLinkedNode{},
        tail:     &DLinkedNode{},
    }
    l.head.next = l.tail
    l.tail.prev = l.head
    return l
}

func (this *LRUCache) Get(key int) int {
    if node, ok := this.cache[key]; ok {
        this.moveToHead(node)
        return node.value
    }
    return -1
}

func (this *LRUCache) Put(key int, value int) {
    if node, ok := this.cache[key]; ok {
        node.value = value
        this.moveToHead(node)
    } else {
        node := &DLinkedNode{key: key, value: value}
        this.cache[key] = node
        this.addToHead(node)
        if len(this.cache) > this.capacity {
            removed := this.removeTail()
            delete(this.cache, removed.key)
        }
    }
}

func (this *LRUCache) addToHead(node *DLinkedNode) {
    node.prev = this.head
    node.next = this.head.next
    this.head.next.prev = node
    this.head.next = node
}

func (this *LRUCache) removeNode(node *DLinkedNode) {
    node.prev.next = node.next
    node.next.prev = node.prev
}

func (this *LRUCache) moveToHead(node *DLinkedNode) {
    this.removeNode(node)
    this.addToHead(node)
}

func (this *LRUCache) removeTail() *DLinkedNode {
    node := this.tail.prev
    this.removeNode(node)
    return node
}
```

#### [460. LFU 缓存](https://leetcode-cn.com/problems/lfu-cache/)

实现 LFU 缓存机制：

```go
type LFUCache struct {
    capacity   int
    minFreq    int
    keyToVal   map[int]int
    keyToFreq  map[int]int
    freqToKeys map[int]*list.List
    keyToNode  map[int]*list.Element
}

func Constructor(capacity int) LFUCache {
    return LFUCache{
        capacity:   capacity,
        keyToVal:   make(map[int]int),
        keyToFreq:  make(map[int]int),
        freqToKeys: make(map[int]*list.List),
        keyToNode:  make(map[int]*list.Element),
    }
}

func (this *LFUCache) Get(key int) int {
    if _, ok := this.keyToVal[key]; !ok {
        return -1
    }
    this.increaseFreq(key)
    return this.keyToVal[key]
}

func (this *LFUCache) Put(key int, value int) {
    if this.capacity <= 0 {
        return
    }
    if _, ok := this.keyToVal[key]; ok {
        this.keyToVal[key] = value
        this.increaseFreq(key)
        return
    }
    if len(this.keyToVal) >= this.capacity {
        this.removeMinFreqKey()
    }
    this.keyToVal[key] = value
    this.keyToFreq[key] = 1
    if this.freqToKeys[1] == nil {
        this.freqToKeys[1] = list.New()
    }
    this.keyToNode[key] = this.freqToKeys[1].PushBack(key)
    this.minFreq = 1
}

func (this *LFUCache) increaseFreq(key int) {
    freq := this.keyToFreq[key]
    this.keyToFreq[key] = freq + 1
    this.freqToKeys[freq].Remove(this.keyToNode[key])
    if this.freqToKeys[freq].Len() == 0 {
        delete(this.freqToKeys, freq)
        if this.minFreq == freq {
            this.minFreq++
        }
    }
    if this.freqToKeys[freq+1] == nil {
        this.freqToKeys[freq+1] = list.New()
    }
    this.keyToNode[key] = this.freqToKeys[freq+1].PushBack(key)
}

func (this *LFUCache) removeMinFreqKey() {
    lst := this.freqToKeys[this.minFreq]
    elem := lst.Front()
    key := elem.Value.(int)
    lst.Remove(elem)
    if lst.Len() == 0 {
        delete(this.freqToKeys, this.minFreq)
    }
    delete(this.keyToVal, key)
    delete(this.keyToFreq, key)
    delete(this.keyToNode, key)
}
```

#### [707. 设计链表](https://leetcode-cn.com/problems/design-linked-list/)

设计双向链表：

```go
type MyLinkedList struct {
    head *Node
    tail *Node
    size int
}

type Node struct {
    val        int
    prev, next *Node
}

func Constructor() MyLinkedList {
    head := &Node{}
    tail := &Node{}
    head.next = tail
    tail.prev = head
    return MyLinkedList{head: head, tail: tail}
}

func (this *MyLinkedList) Get(index int) int {
    if index < 0 || index >= this.size {
        return -1
    }
    cur := this.head.next
    for i := 0; i < index; i++ {
        cur = cur.next
    }
    return cur.val
}

func (this *MyLinkedList) AddAtHead(val int) {
    this.AddAtIndex(0, val)
}

func (this *MyLinkedList) AddAtTail(val int) {
    this.AddAtIndex(this.size, val)
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
    if index > this.size {
        return
    }
    if index < 0 {
        index = 0
    }
    cur := this.head
    for i := 0; i < index; i++ {
        cur = cur.next
    }
    node := &Node{val: val}
    node.next = cur.next
    node.prev = cur
    cur.next.prev = node
    cur.next = node
    this.size++
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
    if index < 0 || index >= this.size {
        return
    }
    cur := this.head
    for i := 0; i < index; i++ {
        cur = cur.next
    }
    cur.next = cur.next.next
    cur.next.prev = cur
    this.size--
}
```
