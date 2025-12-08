## Go 语言特性练习

### 概述

本模块收集 Go 语言核心特性的最小可运行示例，每个子目录专注于一个主题，包含源码和单元测试。遵循"代码即文档"原则，聚焦 Go 开发中常见的坑点和最佳实践。

### 目录结构

| 子目录 | 主题 | 关键内容 |
|-------|------|---------|
| [basics](./basics/) | 语言基础语法 | 变量、常量、控制流、函数、多返回值 |
| [structs](./structs/) | 结构体与方法 | 值/指针接收者、组合/嵌入、`json` tag |
| [interface](./interface/) | 接口与多态 | 行为抽象、空接口、`error` 实现 |
| [slicemap](./slicemap/) | slice 与 map 深入 | nil vs 空、底层共享、原地过滤、map 计数与 Set |
| [defer](./defer/) | 延迟执行与异常恢复 | `defer`、`panic`、`recover` |
| [GoRoutine](./GoRoutine/) | 协程与并发同步 | 闭包陷阱、`sync.Mutex`、`WaitGroup`、`context` |
| [reflect](./reflect/) | 反射机制 | `TypeOf`、`ValueOf`、字段读写 |
| [errors](./errors/) | 错误处理 | `errors.Is/As`、错误包装、自定义错误 |
| [context](./context/) | 上下文与超时 | `WithTimeout`、`WithValue`、`Done()` |
| [channelselect](./channelselect/) | select 与 channel | `FanIn`、日志聚合、心跳触发 |
| [generics](./generics/) | 泛型容器 | `Stack[T]`、`MapSlice`、类型约束 |
| [testingdemo](./testingdemo/) | 测试与基准 | 表驱动测试、`Benchmark` |

---

### basics：语言基础语法

**文件**：`basics/vars.go`、`basics/control.go`、`basics/funcs.go`

**涵盖内容**：
- 变量与常量：`var`、短变量声明 `:=`、零值、`const`
- 控制流：`if` 带初始化、`for` 三种写法、`switch` 与 `type switch`
- 函数：多返回值、命名返回值、变参函数、闭包捕获

```go
// 零值与常量
func ZeroValues() (int, string, bool) {
    var i int
    var s string
    var b bool
    return i, s, b // 0, "", false
}

const Pi = 3.14

// 变参 + for-range
func Sum(nums ...int) int {
    total := 0
    for _, n := range nums {
        total += n
    }
    return total
}

// 闭包：返回一个自增计数器
func NewCounter(start int) func() int {
    counter := start
    return func() int {
        counter++
        return counter
    }
}
```

**运行测试**：`go test ./BasicGo/basics`

---

### structs：结构体与方法

**文件**：`structs/receiver.go`、`structs/embedding.go`、`structs/tag.go`

**核心概念**：
- 值接收者 vs 指针接收者：是否修改调用方持有的对象
- 结构体嵌入：通过组合复用能力，而不是继承
- struct tag：配合 `encoding/json` 做序列化控制

```go
// 值接收者：不会修改原对象
func (u User) RenameValue(newName string) {
    u.Name = newName
}

// 指针接收者：直接修改原对象
func (u *User) RenamePointer(newName string) {
    u.Name = newName
}

// 通过嵌入实现"带日志"的 Service
type Logger struct {
    Prefix string
}

func (l *Logger) Log(msg string) string {
    return l.Prefix + ": " + msg
}

type Service struct {
    Logger        // 嵌入字段，方法会被"提升"
    Name string
}

// Account 展示 json tag 与字段忽略
type Account struct {
    ID    int    `json:"id"`
    Name  string `json:"name,omitempty"`
    Token string `json:"-"` // 不参与 JSON 编解码
}
```

**运行测试**：`go test ./BasicGo/structs`

---

### interface：接口与多态

**文件**：`interface/shape.go`、`interface/empty.go`、`interface/error_interface.go`

**核心概念**：
- 接口隐式实现：类型只要实现方法即可满足接口
- 空接口 `interface{}` / `any`：搭配 `type switch` 做动态分派
- `error` 也是接口：自定义错误类型与错误链（`%w`、`errors.Is/As`）

```go
// 行为接口：只关心 Area，忽略具体形状
type Shape interface {
    Area() float64
}

type Rect struct{ Width, Height float64 }
type Circle struct{ Radius float64 }

func (r Rect) Area() float64   { return r.Width * r.Height }
func (c Circle) Area() float64 { return math.Pi * c.Radius * c.Radius }

// 空接口 + type switch
func Describe(v any) string {
    switch x := v.(type) {
    case nil:
        return "nil"
    case int:
        return fmt.Sprintf("int:%d", x)
    case string:
        return fmt.Sprintf("string:%s", x)
    default:
        return "unknown"
    }
}

// 自定义错误类型实现 error 接口
type OpError struct {
    Op   string
    Code int
    Msg  string
}

func (e OpError) Error() string {
    return fmt.Sprintf("%s failed(%d): %s", e.Op, e.Code, e.Msg)
}
```

**运行测试**：`go test ./BasicGo/interface`

---

### slicemap：slice 与 map 深入

**文件**：`slicemap/slice.go`、`slicemap/map.go`

**核心概念**：
- `nil` 切片 vs 空切片：`len` 一样，为 0，但是否等于 `nil` 不同
- 切片底层共享：多个切片视图指向同一底层数组
- 原地过滤：利用切片复用底层数组，避免额外 allocations
- map 计数与集合：`map[string]int`、`map[string]struct{}` 的常见用法

```go
// 返回 nil 切片与空切片
func MakeNilAndEmpty() (nilSlice, emptySlice []int) {
    var s []int          // nil
    e := make([]int, 0)  // 非 nil，len=0
    return s, e
}

// 底层数组共享示例
func ShareUnderlying() (base, sub, grown []int) {
    base = []int{1, 2, 3, 4}
    sub = base[:2]
    sub[0] = 10      // 修改 sub 同时影响 base
    grown = append(sub, 99)
    return base, sub, grown
}

// 原地过滤奇数
func FilterInPlace(nums []int, keep func(int) bool) []int {
    j := 0
    for _, v := range nums {
        if keep(v) {
            nums[j] = v
            j++
        }
    }
    return nums[:j]
}

// 用 map 模拟 Set
type Set map[string]struct{}
```

**运行测试**：`go test ./BasicGo/slicemap`

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
1. `basics` → 打好语法与函数基础
2. `structs` → 学会用结构体建模数据，理解值/指针接收者
3. `interface` → 掌握行为抽象与接口多态
4. `slicemap` → 熟悉集合操作与常见坑（切片共享、map 用法）
5. `defer` → 理解延迟执行和异常恢复
6. `GoRoutine` / `channelselect` / `context` → 系统掌握并发模型
7. `errors` / `testingdemo` → 错误与测试习惯
8. `generics` / `reflect` → 进阶特性与工具箱能力
