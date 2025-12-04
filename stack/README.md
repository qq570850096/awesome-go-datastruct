# stack 顺序栈

> 最小栈实现 + 括号匹配示例，便于迁移到表达式解析等场景。

## 核心结构
```go
type Stack struct {
    data []byte
    top  int
}
func NewStack(size int) Stack { ... }
```

## 主要方法
- `Push/Pop`：判满/判空后操作切片并维护 `top`。
- `IsEmpty/IsFull`：检查当前元素数。
- 拓展示例：`IsValid(s string) bool`，用栈校验括号合法性。

## 示例
```go
st := NewStack(3)
st.Push('('); st.Push(')')
ok := IsValid(\"([]{})\") // true
```

## 运行
- `go test ./stack`
- 可在测试中调整栈容量或自定义符号表，验证边界条件（空栈出栈、满栈入栈）。
