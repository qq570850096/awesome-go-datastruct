package BehavioralType

import "bytes"

// Context stores the input consumed by expression objects.
type Context struct {
	text string
}

// AbstractExpress is the expression interface. Each expression interprets the
// context and returns an integer result.
type AbstractExpress interface {
	Interpreter(*Context) int
}

// TerminalExpression represents a leaf expression with a fixed value.
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

// Interpreter combines child expressions when the context contains input text.
func (n NonTerminalExpression) Interpreter(ctx *Context) int {

	if !bytes.Equal([]byte(ctx.text), []byte("")) {
		return n.left.Interpreter(ctx) + n.right.Interpreter(ctx)
	}
	return 0
}
