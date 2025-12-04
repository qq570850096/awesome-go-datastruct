# Utils 公共工具

## 概览
- 提供通用比较函数，避免在各数据结构中重复编写比较逻辑。
- 当前核心：`Interfaces.Compare` 支持 `int`、`string`、`float64` 的三态比较。

## 代码结构
- `Interfaces.go`：包占位。
- `Interfaces/Interfaces.go`：`Compare(a, b interface{}) int`，使用反射校验类型一致性后返回 `-1/0/1`。

## 使用建议
- 直接调用 `Interfaces.Compare(a, b)`；若类型不同会 panic，调用前确保类型一致。
- 需要支持其他类型时，在 `switch` 分支中扩展即可。

## 示例
```go
import "awesome-go-datastruct/Utils/Interfaces"

func less(a, b interface{}) bool {
    return Interfaces.Compare(a, b) < 0
}
```
