# DoubleLinked 双向链表与缓存策略

> 代码即文档：链表原语 + 策略组合，直接复用到缓存场景。

## 链表原语
- 位置：`List.go`
- 节点：`key/value` + `prev/next`，支持 O(1) 移除/追加。
- 片段：
```go
func (l *List) AppendToHead(node *Node) *Node { ... }
func (l *List) removeTail() *Node            { ... } // 返回被淘汰节点
func (l *List) Remove(node *Node) *Node      { ... } // 对外暴露删除
```

## 策略实现
- LRU（`LRU.go`）：
  - `Get` 命中移到表头，未命中返回 `nil`。
  - `Put` 满容量时 `removeTail` 淘汰最久未用。
- LFU（`LFU.go`）：
  - 频率桶 `freq -> List`，`updateFreq` 将节点迁移到更高频桶。
  - 淘汰：取最小频率桶尾节点。
- FIFO（`FIFO.go`）：
  - 按写入顺序淘汰尾节点，不刷新命中。

## 示例
```go
lru := InitLRU(2)
lru.Put("a", 1); lru.Put("b", 2)
lru.Get("a")           // 提升 a
lru.Put("c", 3)        // 淘汰 b
fmt.Println(lru.String()) // 观察链表状态
```

## 测试与调试
- 运行：`go test ./DoubleLinked`
- 可调小容量或构造重复访问序列，观察 `String()` 输出的链表顺序与淘汰节点。
