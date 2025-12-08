## 项目入口与示例

### 概述

本目录包含项目的入口文件和独立的 LeetCode 练习代码。

### 文件说明

| 文件 | 说明 |
|-----|------|
| `main.go` | 项目入口，基础示例 |
| `622.go` | LeetCode 622 循环队列实现 |

### 运行方式

```bash
# 运行主程序
go run ./main

# 输出
hello world
```

### LeetCode 622：设计循环队列

实现了一个循环队列的完整功能：

```go
type MyCircularQueue struct {
    arr   []int  // 存储数组
    front int    // 队首指针
    rear  int    // 队尾指针
    size  int    // 当前元素数量
}

// 初始化循环队列，设置容量为 k
func Constructor(k int) MyCircularQueue

// 向循环队列插入元素（满则返回 false）
func (this *MyCircularQueue) EnQueue(value int) bool

// 从循环队列删除元素（空则返回 false）
func (this *MyCircularQueue) DeQueue() bool

// 获取队首元素（空返回 -1）
func (this *MyCircularQueue) Front() int

// 获取队尾元素（空返回 -1）
func (this *MyCircularQueue) Rear() int

// 判断队列是否为空
func (this *MyCircularQueue) IsEmpty() bool

// 判断队列是否已满
func (this *MyCircularQueue) IsFull() bool
```

### 使用示例

```go
// 创建容量为 3 的循环队列
q := Constructor(3)

q.EnQueue(1)  // true
q.EnQueue(2)  // true
q.EnQueue(3)  // true
q.EnQueue(4)  // false（队列已满）

q.Rear()      // 3
q.IsFull()    // true

q.DeQueue()   // true
q.EnQueue(4)  // true

q.Rear()      // 4
```

### 调试用途

`main.go` 可用于临时调试：
1. 导入需要测试的模块
2. 在 `main()` 中编写测试代码
3. 运行 `go run ./main` 查看结果

### 项目导航

本目录为练习入口，各数据结构模块请参考：

| 模块 | 说明 |
|-----|------|
| [queue](../queue/) | 队列模块详细实现 |
| [stack](../stack/) | 栈模块 |
| [Heap](../Heap/) | 堆模块 |
| [Graph_algo](../Graph_algo/) | 图算法模块 |
