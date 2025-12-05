package generics

import (
	"strconv"
	"testing"
)

func TestStack(t *testing.T) {
	var stack Stack[int]
	stack.Push(1)
	stack.Push(2)

	if stack.Len() != 2 {
		t.Fatalf("expected len 2, got %d", stack.Len())
	}
	if v, ok := stack.Pop(); !ok || v != 2 {
		t.Fatalf("unexpected pop result %v %v", v, ok)
	}
	if v, ok := stack.Pop(); !ok || v != 1 {
		t.Fatalf("unexpected pop result %v %v", v, ok)
	}
	if _, ok := stack.Pop(); ok {
		t.Fatalf("expected empty stack")
	}
}

func TestMapAndFilter(t *testing.T) {
	nums := []int{1, 2, 3}
	str := MapSlice(nums, func(v int) string {
		return strconv.Itoa(v)
	})
	if str[0] != "1" || len(str) != 3 {
		t.Fatalf("unexpected map result %v", str)
	}

	filtered := FilterSlice(nums, func(v int) bool { return v%2 == 1 })
	if len(filtered) != 2 || filtered[0] != 1 || filtered[1] != 3 {
		t.Fatalf("unexpected filter result %v", filtered)
	}
}
