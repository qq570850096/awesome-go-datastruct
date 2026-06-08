package basics

import "strconv"

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

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
