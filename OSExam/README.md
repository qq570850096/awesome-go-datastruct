## 操作系统调度算法

### 概述

本模块实现了操作系统课程中常见的进程/作业调度算法，以文件驱动的方式进行调度仿真，便于课堂演示和作业练习。

### 调度算法

| 算法 | 文件 | 特点 | 适用场景 |
|-----|------|------|---------|
| FCFS | `FCFS.go` | 先来先服务，按到达顺序调度 | 批处理系统 |
| SJF | `SJF.go` | 最短作业优先，优化平均等待时间 | 批处理系统 |
| HPF | `FP.go` | 高优先级优先，支持权重周转 | 实时系统 |

### 数据结构

#### 进程/作业

```go
type Process struct {
    pid        int      // 进程ID
    submitTime float64  // 提交时间
    runTime    float64  // 运行时间
    startTime  float64  // 开始时间
    finishTime float64  // 完成时间
    waitTime   float64  // 等待时间
    trTime     float64  // 周转时间 = 完成时间 - 提交时间
    wtrTime    float64  // 带权周转时间 = 周转时间 / 运行时间
    priority   int      // 优先级
    reached    bool     // 是否到达
    visited    bool     // 是否访问过
}
```

### 核心算法

#### FCFS（先来先服务）

```go
type FCFS struct {
    pending []Process  // 待处理作业队列
    ready   []int      // 就绪队列索引
}

// 按到达时间排序
func (this *FCFS) Less(i, j int) bool {
    return this.pending[i].submitTime < this.pending[j].submitTime
}

// 执行调度
func (this *FCFS) FCFS() {
    sort.Sort(this)  // 按到达时间排序
    for i := 0; i < len(this.pending); i++ {
        // 计算等待时间、完成时间、周转时间
        if i == 0 {
            this.pending[i].finishTime = this.pending[i].runTime + this.pending[i].submitTime
            this.pending[i].waitTime = 0
        } else {
            // 上一个作业完成后才能开始
            if this.pending[i-1].finishTime > this.pending[i].submitTime {
                this.pending[i].finishTime = this.pending[i-1].finishTime + this.pending[i].runTime
                this.pending[i].waitTime = this.pending[i-1].finishTime - this.pending[i].submitTime
            } else {
                this.pending[i].finishTime = this.pending[i].runTime + this.pending[i].submitTime
                this.pending[i].waitTime = 0
            }
        }
    }
}
```

#### SJF（最短作业优先）

```go
// 查找下一个最短且已到达的作业
func (s *SJF) FindNextSJF(finish float64) int {
    minTime := math.MaxFloat64
    index := -1
    for i, p := range s.pending {
        if p.submitTime <= finish && !p.visited && p.runTime < minTime {
            minTime = p.runTime
            index = i
        }
    }
    return index
}
```

### 输入文件格式

作业描述文件（如 `test.txt`）：
```
1 0.0 3.0       // 作业ID 到达时间 运行时间
2 2.0 6.0
3 4.0 4.0
4 6.0 5.0
5 8.0 2.0
```

带优先级的文件格式：
```
1 0.0 3.0 3     // 作业ID 到达时间 运行时间 优先级
2 2.0 6.0 5
3 4.0 4.0 2
```

### 使用示例

```go
// FCFS 调度
fcfs := &OSExam.FCFS{}
fcfs.InitFromFile("test.txt")
fcfs.FCFS()

// SJF 调度
sjf := &OSExam.SJF{}
sjf.InitFromFile("test.txt")
sjf.SJF()

// 优先级调度
fp := &OSExam.FP{}
fp.InitFromFile("test1.txt")
fp.HPF()
```

### 性能指标

| 指标 | 计算公式 |
|-----|---------|
| 周转时间 | 完成时间 - 提交时间 |
| 带权周转时间 | 周转时间 / 运行时间 |
| 等待时间 | 周转时间 - 运行时间 |
| 平均周转时间 | Σ周转时间 / 作业数 |

### 文件系统实验

`fileSystem.go` 提供简化的文件系统操作模拟：
- `Cuse/DIR` 结构体
- `DisDir/DisFile` 目录显示
- `DeleteUser` 用户删除

### 运行测试

```bash
go test ./OSExam
```
