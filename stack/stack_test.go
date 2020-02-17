package stack

import (
	"fmt"
	"testing"
)

func TestNewStack(t *testing.T) {
	stack := NewStack(3)
	// 先测试栈为空的时候能否Pop
	fmt.Println(stack.Pop())
	// 测试Push是否正常
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	// 如果栈为正常的，这里Pop打印顺序应该是3,2,1
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())

	fmt.Println(stack.Pop())
}
