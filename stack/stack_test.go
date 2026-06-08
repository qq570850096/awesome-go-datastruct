package stack

import (
	"fmt"
	"testing"
)

func TestNewStack(t *testing.T) {
	stack := NewStack(3)

	fmt.Println(stack.Pop())

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())

	fmt.Println(stack.Pop())
}
