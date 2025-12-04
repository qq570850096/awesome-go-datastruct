# Set 集合实现

> 两种底层对比：顺序链表 vs 二分搜索树。

## 链表集合 `List.go`
- 结构：按插入顺序维护 `head`，操作简单。
- 片段：
```go
func (s *ListSet) Add(e int) {
    if !s.Contains(e) { ... } // 头插
}
func (s *ListSet) Remove(e int) { ... }
```
- 复杂度：`Add/Contains/Remove` 最坏 O(n)，适合小规模或演示。

## BST 集合 `BST.go`
- 结构：手写二分搜索树，未自平衡。
- 片段：
```go
func (t *BST) Add(e int) {
    t.root = add(t.root, e)
}
func (t *BST) Contains(e int) bool { ... } // 递归比较左右子树
func (t *BST) Remove(e int)       { ... } // 删除命中节点后用前驱/后继替换
```
- 复杂度：平均 O(log n)，退化链（有序插入）时 O(n)。

## 使用建议
- 数据量小或无序：链表集合即可。
- 需要更好查找性能：优先 BST 版；若需稳定 O(log n)，可迁移到 `AVL` 或引入自平衡策略。

## 运行
- 当前无自带测试，可自写：`go test ./Set`。
