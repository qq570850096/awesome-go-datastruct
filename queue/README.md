## 循环队列（Circular Queue）

### 定义

**队列**（Queue）是一种遵循**先进先出**（FIFO, First In First Out）原则的线性数据结构。队列只允许在队尾插入元素（入队），在队头删除元素（出队），就像排队买票，先来的人先服务。

**循环队列**通过将数组首尾相连，解决了普通队列的"假溢出"问题，实现了空间的高效利用。

队列的核心操作：
1. **EnQueue**：在队尾插入元素
2. **DeQueue**：从队头删除元素
3. **Front**：查看队头元素

### 为什么使用队列？

队列是计算机科学中最基础的数据结构之一，广泛应用于各种场景：

**应用场景：**
- 任务调度（操作系统进程调度）
- 消息队列（异步通信）
- 广度优先搜索（BFS）
- 缓冲区（打印机队列、IO 缓冲）
- 生产者-消费者模型

**生活类比：** 银行排队叫号——先取号的先办理业务；地铁进站——先到的乘客先上车。

### 特性

| 操作 | 时间复杂度 | 说明 |
|------|-----------|------|
| EnQueue | O(1) | 入队（摊还） |
| DeQueue | O(1) | 出队 |
| Front | O(1) | 查看队头 |
| Rear | O(1) | 查看队尾 |
| IsEmpty | O(1) | 判断是否为空 |
| IsFull | O(1) | 判断是否已满 |
| 空间复杂度 | O(n) | n 为队列容量 |

### 数据结构

```go
// Queue 循环队列实现
type Queue struct {
    container []int // 存储元素的数组
    front     int   // 队头指针
    rear      int   // 队尾指针
    size      int   // 当前元素数量
}

// NewQueue 创建指定容量的循环队列
func NewQueue(k int) *Queue {
    return &Queue{
        container: make([]int, k),
        front:     0,
        rear:      0,
        size:      0,
    }
}
```

### 核心方法实现

#### EnQueue 入队操作

```go
// EnQueue 将元素添加到队尾
// 如果队列已满返回 false，否则返回 true
func (this *Queue) EnQueue(value int) bool {
    if this.container == nil || this.IsFull() {
        return false
    }
    // 使用取模实现循环
    this.container[this.rear%len(this.container)] = value
    this.rear = this.rear%len(this.container) + 1
    this.size++
    return true
}
```

#### DeQueue 出队操作

```go
// DeQueue 从队头移除元素
// 返回操作是否成功以及出队的元素
func (this *Queue) DeQueue() (bool, int) {
    if this.container == nil || this.IsEmpty() {
        return false, 0
    }
    ret := this.container[this.front%len(this.container)]
    this.front = this.front%len(this.container) + 1
    this.size--
    return true, ret
}
```

#### 辅助方法

```go
// IsEmpty 判断队列是否为空
func (this *Queue) IsEmpty() bool {
    return this.size == 0
}

// IsFull 判断队列是否已满
func (this *Queue) IsFull() bool {
    return this.size == len(this.container)
}

// Front 获取队头元素
func (this *Queue) Front() (bool, int) {
    if this.IsEmpty() {
        return false, 0
    }
    return true, this.container[this.front%len(this.container)]
}

// Rear 获取队尾元素
func (this *Queue) Rear() (bool, int) {
    if this.IsEmpty() {
        return false, 0
    }
    // rear 指向下一个待插入位置，所以要减1
    idx := (this.rear - 1 + len(this.container)) % len(this.container)
    return true, this.container[idx]
}
```

### 循环队列示意图

```
容量为 5 的循环队列：

初始状态：front = rear = 0, size = 0
┌───┬───┬───┬───┬───┐
│   │   │   │   │   │
└───┴───┴───┴───┴───┘
  ↑
front,rear

EnQueue(1), EnQueue(2), EnQueue(3)：
┌───┬───┬───┬───┬───┐
│ 1 │ 2 │ 3 │   │   │  size = 3
└───┴───┴───┴───┴───┘
  ↑           ↑
front       rear

DeQueue() -> 1, DeQueue() -> 2：
┌───┬───┬───┬───┬───┐
│   │   │ 3 │   │   │  size = 1
└───┴───┴───┴───┴───┘
          ↑   ↑
        front rear

EnQueue(4), EnQueue(5), EnQueue(6)：（循环利用空间）
┌───┬───┬───┬───┬───┐
│ 6 │   │ 3 │ 4 │ 5 │  size = 4
└───┴───┴───┴───┴───┘
  ↑   ↑
rear front
```

### 测试用例

```go
func TestQueue(t *testing.T) {
    queue := NewQueue(5)

    // 测试入队
    queue.EnQueue(1)
    queue.EnQueue(2)
    queue.EnQueue(3)
    fmt.Println("队列是否为空:", queue.IsEmpty()) // false

    // 测试出队
    _, val := queue.DeQueue()
    fmt.Println("出队元素:", val) // 1

    // 测试循环特性
    queue.EnQueue(4)
    queue.EnQueue(5)
    queue.EnQueue(6) // 循环到数组开头

    // 连续出队
    for !queue.IsEmpty() {
        _, v := queue.DeQueue()
        fmt.Printf("%d ", v)
    }
    // 输出: 2 3 4 5 6
}

func main() {
    queue := NewQueue(5)
    // 循环3次，每次添加5个元素，再出队三个元素
    for i := 0; i < 3; i++ {
        for j := 0; j < 5; j++ {
            queue.EnQueue(j)
        }
        for k := 0; k < 3; k++ {
            fmt.Println(queue.DeQueue())
        }
    }
}
```

### 运行方式

```bash
go run ./queue
```

### LeetCode 实战

#### [622. 设计循环队列](https://leetcode-cn.com/problems/design-circular-queue/)

实现完整的循环队列：

```go
type MyCircularQueue struct {
    data  []int
    front int
    rear  int
    size  int
    cap   int
}

func Constructor(k int) MyCircularQueue {
    return MyCircularQueue{
        data:  make([]int, k),
        front: 0,
        rear:  0,
        size:  0,
        cap:   k,
    }
}

func (this *MyCircularQueue) EnQueue(value int) bool {
    if this.IsFull() {
        return false
    }
    this.data[this.rear] = value
    this.rear = (this.rear + 1) % this.cap
    this.size++
    return true
}

func (this *MyCircularQueue) DeQueue() bool {
    if this.IsEmpty() {
        return false
    }
    this.front = (this.front + 1) % this.cap
    this.size--
    return true
}

func (this *MyCircularQueue) Front() int {
    if this.IsEmpty() {
        return -1
    }
    return this.data[this.front]
}

func (this *MyCircularQueue) Rear() int {
    if this.IsEmpty() {
        return -1
    }
    return this.data[(this.rear-1+this.cap)%this.cap]
}

func (this *MyCircularQueue) IsEmpty() bool {
    return this.size == 0
}

func (this *MyCircularQueue) IsFull() bool {
    return this.size == this.cap
}
```

#### [232. 用栈实现队列](https://leetcode-cn.com/problems/implement-queue-using-stacks/)

使用两个栈实现队列的先进先出：

```go
type MyQueue struct {
    stackIn  []int // 入队栈
    stackOut []int // 出队栈
}

func Constructor() MyQueue {
    return MyQueue{}
}

func (this *MyQueue) Push(x int) {
    this.stackIn = append(this.stackIn, x)
}

func (this *MyQueue) Pop() int {
    // 如果出队栈为空，将入队栈元素全部倒入
    if len(this.stackOut) == 0 {
        for len(this.stackIn) > 0 {
            this.stackOut = append(this.stackOut, this.stackIn[len(this.stackIn)-1])
            this.stackIn = this.stackIn[:len(this.stackIn)-1]
        }
    }
    val := this.stackOut[len(this.stackOut)-1]
    this.stackOut = this.stackOut[:len(this.stackOut)-1]
    return val
}

func (this *MyQueue) Peek() int {
    val := this.Pop()
    this.stackOut = append(this.stackOut, val)
    return val
}

func (this *MyQueue) Empty() bool {
    return len(this.stackIn) == 0 && len(this.stackOut) == 0
}
```

#### [225. 用队列实现栈](https://leetcode-cn.com/problems/implement-stack-using-queues/)

使用队列实现栈的后进先出：

```go
type MyStack struct {
    queue []int
}

func Constructor() MyStack {
    return MyStack{}
}

func (this *MyStack) Push(x int) {
    n := len(this.queue)
    this.queue = append(this.queue, x)
    // 将新元素之前的所有元素出队再入队
    for i := 0; i < n; i++ {
        this.queue = append(this.queue, this.queue[0])
        this.queue = this.queue[1:]
    }
}

func (this *MyStack) Pop() int {
    val := this.queue[0]
    this.queue = this.queue[1:]
    return val
}

func (this *MyStack) Top() int {
    return this.queue[0]
}

func (this *MyStack) Empty() bool {
    return len(this.queue) == 0
}
```

#### [933. 最近的请求次数](https://leetcode-cn.com/problems/number-of-recent-calls/)

使用队列统计滑动窗口内的请求数：

```go
type RecentCounter struct {
    queue []int
}

func Constructor() RecentCounter {
    return RecentCounter{queue: []int{}}
}

func (this *RecentCounter) Ping(t int) int {
    this.queue = append(this.queue, t)
    // 移除 3000ms 之前的请求
    for this.queue[0] < t-3000 {
        this.queue = this.queue[1:]
    }
    return len(this.queue)
}
```
