package basics

import "strconv"

// Max 返回两个整数中较大的一个。
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// FizzBuzz 演示 if / else 与取模运算。
// 3 的倍数返回 "Fizz"，5 的倍数返回 "Buzz"，同时是 3 和 5 的倍数返回 "FizzBuzz"，否则返回数字本身。
func FizzBuzz(n int) string {
	switch {
	case n%15 == 0:
		return "FizzBuzz"
	case n%3 == 0:
		return "Fizz"
	case n%5 == 0:
		return "Buzz"
	default:
		return strconv.Itoa(n)
	}
}

// TypeName 使用 type switch 返回值的大致类型名称。
func TypeName(v any) string {
	switch v.(type) {
	case int:
		return "int"
	case string:
		return "string"
	case bool:
		return "bool"
	default:
		return "unknown"
	}
}

