package BehavioralType

import "testing"

//测试类
func TestNonTerminalExpression_Interpreter(t *testing.T) {
	var (
		left AbstractExpress
		right AbstractExpress
		callExpression AbstractExpress
	)
	left = &TerminalExpression{arg:12}
	right = &TerminalExpression{arg:34}
	callExpression = &NonTerminalExpression{left:left,right:right}

	context := &Context{text:"+"}

	result := callExpression.Interpreter(context)
	t.Log(result)
}
