package basics

import "strings"

// Sum 演示变参函数的写法。
func Sum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

// SplitName 演示多返回值和命名返回值。
// 假设输入格式为 "first last"，仅在第一个空格处分割。
func SplitName(full string) (first, last string) {
	parts := strings.Fields(full)
	if len(parts) == 0 {
		return "", ""
	}
	if len(parts) == 1 {
		return parts[0], ""
	}
	return parts[0], parts[1]
}

// NewCounter 返回一个闭包函数，每次调用返回自增后的计数值。
func NewCounter(start int) func() int {
	counter := start
	return func() int {
		counter++
		return counter
	}
}

