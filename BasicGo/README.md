## Go 语言特性练习

### 概述

本模块收集 Go 语言核心特性的最小可运行示例，每个子目录专注于一个主题，包含源码和单元测试。遵循"代码即文档"原则，聚焦 Go 开发中常见的坑点和最佳实践。

### 目录结构

| 子目录 | 主题 | 关键内容 |
|-------|------|---------|
| [defer](./defer/) | 延迟执行与异常恢复 | `defer`、`panic`、`recover` |
| [GoRoutine](./GoRoutine/) | 协程与并发同步 | 闭包陷阱、`sync.Mutex`、`WaitGroup`、`context` |
| [reflect](./reflect/) | 反射机制 | `TypeOf`、`ValueOf`、字段读写 |
| [errors](./errors/) | 错误处理 | `errors.Is/As`、错误包装、自定义错误 |
| [context](./context/) | 上下文与超时 | `WithTimeout`、`WithValue`、`Done()` |
| [channelselect](./channelselect/) | select 与 channel | `FanIn`、日志聚合、心跳触发 |
| [generics](./generics/) | 泛型容器 | `Stack[T]`、`MapSlice`、类型约束 |
| [testingdemo](./testingdemo/) | 测试与基准 | 表驱动测试、`Benchmark` |

---

### defer：延迟执行与异常恢复

**文件**：`defer/defer.go`

**核心概念**：
- `defer` 语句在函数返回前执行，常用于资源清理
- 结合 `recover()` 可捕获 `panic`，防止程序崩溃

```go
// 异常恢复示例
func Error() {
    defer func() {
        if err := recover(); err != nil {
            fmt.Println("Recovered from", err)
        }
    }()
    panic(errors.New("Something Wrong!"))
}
```

**常见陷阱**：
- `defer` 参数在声明时求值，而非执行时
- 多个 `defer` 按 LIFO 顺序执行

**运行测试**：`go test ./BasicGo/defer`

---

### GoRoutine：协程与并发同步

**文件**：`GoRoutine/Goroutine.go`

#### 闭包捕获问题

```go
// 正确：传值捕获，每个协程获得独立的 i
func Thread() {
    for i := 0; i < 10; i++ {
        go func(i int) {
            fmt.Println(i)
        }(i)
    }
}

// 错误：共享循环变量，输出结果不可预测
func ThreadWrong() {
    for i := 0; i < 10; i++ {
        go func() {
            fmt.Println(i) // 所有协程共享同一个 i
        }()
    }
}
```

#### 互斥锁与 WaitGroup

```go
// 使用 sync.Mutex 保证原子操作
func Counter() int {
    var mut sync.Mutex
    counter := 0
    for i := 0; i < 5000; i++ {
        go func() {
            mut.Lock()
            defer mut.Unlock()
            counter++
        }()
    }
    time.Sleep(time.Second)
    return counter
}

// 使用 WaitGroup 等待所有协程完成
func WaitGroupExam() int {
    var wg sync.WaitGroup
    var mut sync.Mutex
    counter := 0
    for i := 0; i < 5000; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            mut.Lock()
            defer mut.Unlock()
            counter++
        }()
    }
    wg.Wait()
    return counter
}
```

#### CSP 并发模式

```go
// 异步服务：带缓冲 channel 实现非阻塞发送
func AsnyService() <-chan string {
    retCh := make(chan string, 1) // 缓冲区为1，异步执行
    go func() {
        ret := Service()
        retCh <- ret
    }()
    return retCh
}

// 生产者-消费者模式
func Producer(ch chan<- int, group *sync.WaitGroup) {
    go func() {
        for i := 0; i < 100; i++ {
            ch <- i
        }
        defer close(ch)
        defer group.Done()
    }()
}

// 对象池：使用 buffered channel 实现
type ObjPool struct {
    bufChan chan *ReusableObj
}

func (p *ObjPool) GetObj(timeout time.Duration) (*ReusableObj, error) {
    select {
    case ret := <-p.bufChan:
        return ret, nil
    case <-time.After(timeout):
        return nil, errors.New("time out!")
    }
}
```

**运行测试**：`go test ./BasicGo/GoRoutine`

---

### reflect：反射机制

**文件**：`reflect/Reflect.go`

**核心 API**：

| 函数 | 用途 |
|-----|------|
| `reflect.TypeOf(v)` | 获取变量的类型信息 |
| `reflect.ValueOf(v)` | 获取变量的值信息 |
| `Value.Elem()` | 获取指针指向的值 |
| `Value.FieldByName()` | 按名称获取字段 |
| `Value.SetInt()` | 设置整数值（需可寻址） |

**使用注意**：
- 修改值必须使用指针，且调用 `Elem()` 获取可寻址的值
- 反射操作有性能开销，避免在热路径使用

**运行测试**：`go test ./BasicGo/reflect`

---

### errors：错误处理

**文件**：`errors/errors.go`

**Go 1.13+ 错误处理模式**：

```go
// 错误包装
err := fmt.Errorf("operation failed: %w", originalErr)

// 错误判断
if errors.Is(err, targetErr) {
    // 判断错误链中是否包含特定错误
}

// 错误类型断言
var validErr *ValidationError
if errors.As(err, &validErr) {
    // 从错误链中提取特定类型的错误
}
```

**最佳实践**：
- 使用 sentinel error（包级变量）表示特定错误
- 自定义错误类型实现 `Error()` 方法
- 使用 `%w` 包装错误保留错误链

**运行测试**：`go test ./BasicGo/errors`

---

### context：上下文与超时控制

**文件**：`context/timeout.go`

#### 超时控制

```go
// FetchWithTimeout 在指定超时内执行任务
func FetchWithTimeout(parent context.Context, timeout time.Duration,
    work func(context.Context) (string, error)) (string, error) {
    ctx, cancel := context.WithTimeout(parent, timeout)
    defer cancel()

    ch := make(chan result, 1)
    go func() {
        val, err := work(ctx)
        ch <- result{data: val, err: err}
    }()

    select {
    case <-ctx.Done():
        return "", ctx.Err()  // 超时或取消
    case res := <-ch:
        return res.data, res.err
    }
}
```

#### 请求追踪

```go
// 使用自定义 key 类型存储上下文值
type requestIDKey struct{}

func ContextWithRequestID(parent context.Context, id string) context.Context {
    return context.WithValue(parent, requestIDKey{}, id)
}

func RequestIDFromContext(ctx context.Context) string {
    if v, ok := ctx.Value(requestIDKey{}).(string); ok {
        return v
    }
    return ""
}
```

**运行测试**：`go test ./BasicGo/context`

---

### channelselect：select 与方向 channel

**文件**：`channelselect/select.go`

**核心模式**：

| 模式 | 说明 |
|-----|------|
| `FanIn` | 合并多个只读 channel 到一个输出 |
| `AggregateLogs` | 日志聚合的实际场景 |
| `Ticker` | 心跳/轮询触发器 |
| `Drain` | `for-range` 消费已关闭的 channel |

**channel 方向约束**：
- `chan<- T`：只写 channel
- `<-chan T`：只读 channel
- 编译时检查，防止误用

**运行测试**：`go test ./BasicGo/channelselect`

---

### generics：泛型容器（Go 1.18+）

**文件**：`generics/stack.go`

#### 泛型栈

```go
type Stack[T any] struct {
    data []T
}

func (s *Stack[T]) Push(v T) {
    s.data = append(s.data, v)
}

func (s *Stack[T]) Pop() (T, bool) {
    var zero T
    if len(s.data) == 0 {
        return zero, false
    }
    v := s.data[len(s.data)-1]
    s.data = s.data[:len(s.data)-1]
    return v, true
}
```

#### 泛型函数

```go
// 多类型参数：T 输入类型，R 输出类型
func MapSlice[T any, R any](items []T, fn func(T) R) []R {
    res := make([]R, 0, len(items))
    for _, item := range items {
        res = append(res, fn(item))
    }
    return res
}

// comparable 约束：支持 == 和 != 操作
func FilterSlice[T comparable](items []T, keep func(T) bool) []T {
    out := items[:0]
    for _, item := range items {
        if keep(item) {
            out = append(out, item)
        }
    }
    return out
}
```

**类型约束**：
- `any`：任意类型
- `comparable`：可比较类型（支持 `==`）
- 自定义接口约束

**运行测试**：`go test ./BasicGo/generics`

---

### testingdemo：表驱动测试与基准

**文件**：`testingdemo/calculator.go`

#### 表驱动测试模式

```go
func TestCalc(t *testing.T) {
    tests := []struct {
        name     string
        a, b     int
        op       string
        expected int
        wantErr  bool
    }{
        {"add", 1, 2, "+", 3, false},
        {"sub", 5, 3, "-", 2, false},
        {"mul", 2, 3, "*", 6, false},
        {"div", 6, 2, "/", 3, false},
        {"divByZero", 1, 0, "/", 0, true},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got, err := Calc(tt.a, tt.b, tt.op)
            if (err != nil) != tt.wantErr {
                t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
            }
            if got != tt.expected {
                t.Errorf("got %v, want %v", got, tt.expected)
            }
        })
    }
}
```

#### 基准测试

```go
func BenchmarkCalc(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Calc(100, 50, "+")
    }
}
```

**运行测试**：
```bash
go test ./BasicGo/testingdemo          # 运行测试
go test ./BasicGo/testingdemo -bench=. # 运行基准测试
```

---

### 快速上手

```bash
# 运行所有模块测试
go test ./BasicGo/...

# 运行单个模块
go test ./BasicGo/GoRoutine
go test ./BasicGo/context
go test ./BasicGo/generics

# 运行基准测试
go test ./BasicGo/testingdemo -bench=.
```

**推荐学习顺序**：
1. `defer` → 理解延迟执行和异常恢复
2. `GoRoutine` → 掌握并发基础和常见陷阱
3. `context` → 学习超时控制和取消机制
4. `channelselect` → 深入 CSP 并发模式
5. `generics` → 了解 Go 1.18+ 泛型特性
