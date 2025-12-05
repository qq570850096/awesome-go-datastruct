package testingdemo

import "errors"

// Operation 定义为字符串，方便 table-driven 测试扩展。
type Operation string

const (
	Add Operation = "add"
	Sub Operation = "sub"
	Mul Operation = "mul"
	Div Operation = "div"
)

var ErrUnknownOp = errors.New("unknown operation")
var ErrDivideByZero = errors.New("divide by zero")

func Calc(a, b float64, op Operation) (float64, error) {
	switch op {
	case Add:
		return a + b, nil
	case Sub:
		return a - b, nil
	case Mul:
		return a * b, nil
	case Div:
		if b == 0 {
			return 0, ErrDivideByZero
		}
		return a / b, nil
	default:
		return 0, ErrUnknownOp
	}
}
