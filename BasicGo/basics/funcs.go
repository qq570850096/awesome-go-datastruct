package basics

import "strings"

func Sum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

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

func NewCounter(start int) func() int {
	counter := start
	return func() int {
		counter++
		return counter
	}
}
