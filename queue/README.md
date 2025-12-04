# queue 循环队列（数组实现）

> 与 `main/622.go` 同题不同解，便于对照学习。

## 核心结构
```go
type Queue struct {
    container []int
    front int
    rear  int
    size  int
}
```
- 取模移动指针，`size` 独立计数，区分空与满。

## 主要方法（片段）
```go
func (q *Queue) EnQueue(v int) bool {
    if q.IsFull() { return false }
    q.container[q.rear%len(q.container)] = v
    q.rear = q.rear%len(q.container) + 1
    q.size++
    return true
}

func (q *Queue) DeQueue() (bool, int) {
    if q.IsEmpty() { return false, 0 }
    ret := q.container[q.front%len(q.container)]
    q.front = q.front%len(q.container) + 1
    q.size--
    return true, ret
}
```

## 示例运行
- 文件内 `main()` 会多轮入队/出队，打印结果验证指针移动：
  - 执行：`go run ./queue`
- 也可在其他文件引入本实现，与 `main/622.go` 的版本对比 API 命名与判满逻辑。
