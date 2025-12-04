# main 示例入口与环形队列

## 文件说明
- `main.go`：最小调试入口（Hello World），可临时写演示代码。
- `622.go`：LeetCode 622 环形队列实现，暴露完整 API。

## 核心代码（摘录）
```go
type MyCircularQueue struct {
    arr   []int
    front int
    rear  int
    size  int
}

func Constructor(k int) MyCircularQueue { ... }
func (q *MyCircularQueue) EnQueue(v int) bool { ... }  // 入队（满则 false）
func (q *MyCircularQueue) DeQueue() bool { ... }       // 出队（空则 false）
func (q *MyCircularQueue) Front() int { ... }          // 空返回 -1
func (q *MyCircularQueue) Rear() int { ... }           // 空返回 -1
func (q *MyCircularQueue) IsEmpty() bool { ... }
func (q *MyCircularQueue) IsFull() bool { ... }
```

## 运行方式
- 占位入口：`go run ./main`
- 队列演示：在 `main.go` 内创建 `q := Constructor(3)` 并调用各方法，或复制到其他练习文件与 `queue/queue.go` 版本对照。
