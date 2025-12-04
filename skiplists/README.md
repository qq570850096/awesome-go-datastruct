# skiplists 跳表

> 概率平衡有序表，多层前进指针实现近似 O(log n)。

## 数据结构
- 节点：`value`、`score`（排序键）、`level` 层的前进指针数组。
- 头节点持最高层，便于从上到下逐层逼近插入/查找位置。

## 关键方法
```go
func (sl *SkipList) Search(score float64) (*Node, bool) { ... }
func (sl *SkipList) Insert(score float64, value interface{}) *Node { ... }
func (sl *SkipList) Delete(score float64) *Node { ... }
```
- 插入：从顶层向下寻找插入点，随机层高 `randomLevel()` 决定新节点高度。
- 删除：同样先收集前驱指针，再逐层断链。

## 随机层高
- `randomLevel()` 以固定概率递增层数，保证期望高度 `O(log n)`，避免显式旋转/重建。
- 可调整概率参数以权衡高度与更新成本。

## 运行与调试
- 运行：`go test ./skiplists`
- 建议打印节点 `level` 分布或手动插入有序/重复分布，观察链路变化。
