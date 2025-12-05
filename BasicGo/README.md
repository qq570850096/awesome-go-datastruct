# BasicGo 语言特性练习

> 代码即文档：每个子目录带最小可运行示例和单测，聚焦常见坑。

## defer：异常恢复
- 位置：`defer/defer.go`
- 核心：利用 `recover` 防止 panic 失控。
```go
func Error() {
    defer func() {
        if err := recover(); err != nil {
            fmt.Println("Recovered from", err)
        }
    }()
    panic(errors.New("Something Wrong!"))
}
```
- 运行：`go test ./BasicGo/defer`

## GoRoutine：协程与同步
- 位置：`GoRoutine/Goroutine.go`
- 闭包捕获对比：`Thread`（传值） vs `ThreadWrong`（共享循环变量）。
- 互斥计数：`Counter` 使用 `sync.Mutex`，`WaitGroupExam` 演示 `WaitGroup`，`Cancel` 展示 `context.WithCancel`。
- CSP 示例：`Producer/Consumer`、`AsnyService`、`ObjPool` 等覆盖 channel 缓冲、关闭与超时。
- 运行：`go test ./BasicGo/GoRoutine`

## reflect：反射读写
- 位置：`reflect/Reflect.go`
- 核心：`reflect.TypeOf` / `ValueOf`，演示修改字段与读取类型信息。
- 运行：`go test ./BasicGo/reflect`

## errors：错误处理
- 位置：`errors/errors.go`
- 展示 `errors.Is/As`、sentinel error、结构化 `ValidationError` 以及 `fmt.Errorf("%w")` 包装。
- 运行：`go test ./BasicGo/errors`

## context：值与超时
- 位置：`context/timeout.go`
- `FetchWithTimeout` 提供统一的超时控制；`ContextWithRequestID` / `RequestIDFromContext` 传递追踪信息；`ExampleCancelableWork` 演示 `ctx.Done()` 的退出机制。
- 运行：`go test ./BasicGo/context`

## channelselect：select 与方向 channel
- 位置：`channelselect/select.go`
- `FanIn` 合并多个只读 channel；`AggregateLogs` 给出日志聚合的实际场景，`Ticker` 可作心跳/轮询触发器，`Drain` 展示 `for-range` 消费已关闭 channel。
- 运行：`go test ./BasicGo/channelselect`

## generics：泛型容器
- 位置：`generics/stack.go`
- 实现 `Stack[T]`、`MapSlice`、`FilterSlice`，覆盖 `any`、多类型参数、`comparable` 约束。
- 运行：`go test ./BasicGo/generics`

## testingdemo：表驱动测试与基准
- 位置：`testingdemo/calculator.go`
- `Calc` 支持加减乘除；`calculator_test.go` 用 table-driven 子测试与 `BenchmarkCalc`，演示 `go test -bench .`。
- 运行：`go test ./BasicGo/testingdemo -bench=.`

## 快速上手
- 全部模块：`go test ./BasicGo/...`
- 推荐先读源码再跑测试，理解每个示例修复的具体坑点。
