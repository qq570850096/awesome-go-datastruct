## 工具函数

### 概述

本模块提供数据结构实现中常用的工具函数和接口定义，避免在各模块中重复编写比较逻辑。

### 目录结构

| 文件 | 说明 |
|-----|------|
| `Interfaces.go` | 包占位文件 |
| `Interfaces/Interfaces.go` | 通用比较函数实现 |

### 通用比较函数

```go
// Compare 比较两个相同类型的值
// 返回值：a > b 返回 1，a < b 返回 -1，a == b 返回 0
func Compare(a interface{}, b interface{}) int {
    aType := reflect.TypeOf(a).String()
    bType := reflect.TypeOf(b).String()

    if aType != bType {
        panic("cannot compare different type params")
    }

    switch a.(type) {
    case int:
        if a.(int) > b.(int) {
            return 1
        } else if a.(int) < b.(int) {
            return -1
        }
        return 0
    case string:
        // 字符串字典序比较
    case float64:
        // 浮点数比较
    default:
        panic("unsupported type params")
    }
}
```

### 支持的类型

| 类型 | 说明 |
|-----|------|
| `int` | 整数比较 |
| `string` | 字符串字典序比较 |
| `float64` | 浮点数比较 |

### 使用示例

```go
import "awesome-go-datastruct/Utils/Interfaces"

// 整数比较
result := Interfaces.Compare(3, 5)  // 返回 -1

// 字符串比较
result := Interfaces.Compare("abc", "abd")  // 返回 -1

// 封装为 less 函数
func less(a, b interface{}) bool {
    return Interfaces.Compare(a, b) < 0
}
```

### 使用注意

- 调用前确保两个参数类型一致，类型不同会 panic
- 需要支持其他类型时，在 `switch` 分支中扩展即可

### 设计说明

在 Go 1.18 之前，Go 语言不支持泛型，因此使用 `interface{}` + 反射来实现通用比较。

Go 1.18+ 推荐使用泛型约束：
```go
import "golang.org/x/exp/constraints"

func Compare[T constraints.Ordered](a, b T) int {
    if a > b { return 1 }
    if a < b { return -1 }
    return 0
}
```
