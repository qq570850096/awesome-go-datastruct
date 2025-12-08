package interfacedemo

import "fmt"

// Describe 使用空接口 + type switch 打印值的类型信息。
func Describe(v any) string {
	switch x := v.(type) {
	case nil:
		return "nil"
	case int:
		return fmt.Sprintf("int:%d", x)
	case string:
		return fmt.Sprintf("string:%s", x)
	case bool:
		return fmt.Sprintf("bool:%t", x)
	default:
		return "unknown"
	}
}

