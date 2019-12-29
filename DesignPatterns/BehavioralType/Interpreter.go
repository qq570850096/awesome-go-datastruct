package BehavioralType

import "bytes"

type Context struct {
	text string
}

//抽象解释器
type AbstractExpress interface {
	Interpreter(*Context) int
}

// 终结符，即我们的参数构造类
type TerminalExpression struct {
	arg int
}

func (t *TerminalExpression) Interpreter(ctx *Context) int {
	return t.arg
}

// 非终结符，即我们的运算符构造类
type NonTerminalExpression struct {
	left AbstractExpress
	right AbstractExpress
}

func (n NonTerminalExpression) Interpreter(ctx *Context) int {
	// 实现具体的a+b的解释执行操作
	if !bytes.Equal([]byte(ctx.text),[]byte("")) {
		return n.left.Interpreter(ctx) + n.right.Interpreter(ctx)
	}
	return 0
}







