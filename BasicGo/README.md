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
- 互斥计数：`Counter` 使用 `sync.Mutex`，展示竞态修复。
```go
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
```
- `WaitGroupExam` 展示等待所有协程完成；`ContextCancel`（在测试中）演示取消。
- 运行：`go test ./BasicGo/GoRoutine`

## reflect：反射读写
- 位置：`reflect/Reflect.go`
- 核心：`reflect.TypeOf` / `ValueOf`，演示修改字段与读取类型信息。
- 运行：`go test ./BasicGo/reflect`

## 快速上手
- 全部模块：`go test ./BasicGo/...`
- 推荐先读源码再跑测试，理解每个示例修复的具体坑点。
