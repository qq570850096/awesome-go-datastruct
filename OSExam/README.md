# OSExam 操作系统实验题

> 以文件驱动的调度仿真 + 简化文件系统操作，便于课堂/作业演示。

## 调度算法
- 共性：各调度器实现 `InitFromFile` 读取作业列表，再根据策略出队。
- 数据结构：`container/heap` 接口（`Len/Less/Swap/Push/Pop`）实现优先队列。

### FCFS（先来先服务）`FCFS.go`
- 队列按到达顺序出队。
```go
func (q *FCFS) Less(i, j int) bool { return q.Arr[i].StartTime < q.Arr[j].StartTime }
func (q *FCFS) FCFS() { ... } // 依次弹出，累积等待/周转时间
```

### SJF（最短作业优先）`SJF.go`
- 就绪队列按作业运行时间升序选择下一个。
```go
func (s *SJF) FindNextSJF(finish float64) int { ... } // 选择最短且已到达
```

### FP/HPF（高优先级优先）`FP.go`
- 按优先级（数字越小权重越高）调度，支持权重周转时间计算。

## 文件系统实验
- 位置：`fileSystem.go`
- 结构体 `Cuse/DIR`，提供 `DisDir/DisFile/DeleteUser` 等基础操作，用于模拟目录状态变化。

## 测试与输入
- 输入样例：`test.txt`、`test1.txt`（作业到达时间、运行时间、优先级）。
- 运行全部测试：`go test ./OSExam`
- 自定义场景：修改/新增输入文件后，调用对应 `InitFromFile` 读取，再运行调度函数。
