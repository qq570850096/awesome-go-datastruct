package BehavioralType

import "bytes"

type Context struct {
	text string
}

type AbstractExpress interface {
	Interpreter(*Context) int
}

type TerminalExpression struct {
	arg int
}

func (t *TerminalExpression) Interpreter(ctx *Context) int {
	return t.arg
}

type NonTerminalExpression struct {
	left  AbstractExpress
	right AbstractExpress
}

func (n NonTerminalExpression) Interpreter(ctx *Context) int {

	if !bytes.Equal([]byte(ctx.text), []byte("")) {
		return n.left.Interpreter(ctx) + n.right.Interpreter(ctx)
	}
	return 0
}
