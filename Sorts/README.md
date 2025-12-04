# Sorts 排序算法合集

> 通过一组可运行实现对比复杂度、稳定性与适用场景。

## 收录算法与文件
- O(n^2)：`Bubble.go`、`Insert.go`、`Select.go`
- 分治：`Merge.go`（递归 + 自底向上）、`Quick.go`（双路/三路）
- 其他：`Shell.go`（gap 递减）、`Bucket.go`（计数/桶排序示例）
- 基准：`Sort_test.go`

## 关键片段
- 归并（自顶向下）：
```go
func MergeSort(arr []int, l, r int) {
    if l >= r { return }
    mid := (l + r) / 2
    MergeSort(arr, l, mid)
    MergeSort(arr, mid+1, r)
    merge(arr, aux, l, mid, r)
}
```
- 三路快排（重复元素友好）：
```go
func quickSort3Ways(arr []int, l, r int) {
    if l >= r { return }
    // < v | == v | > v
    ...
}
```
- 希尔：
```go
for gap := n/2; gap > 0; gap /= 2 { ... }
```

## 复杂度与适用性
- 冒泡/插入/选择：O(n^2)，小数据或近乎有序。
- 归并/快排：平均 O(n log n)；归并稳定、占用额外空间，快排就地但需注意随机化。
- 希尔/桶：依赖 gap/桶设计，对特定分布有优势。

## 运行与对比
- 基准：`go test ./Sorts -bench .`
- 可在 `Sort_test.go` 调整数据规模、重复度，比较不同算法耗时与结果正确性。
