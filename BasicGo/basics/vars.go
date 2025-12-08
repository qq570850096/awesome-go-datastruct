package basics

// ZeroValues 返回几种基础类型的零值，用于展示 Go 的默认初始化行为。
func ZeroValues() (int, string, bool) {
	var i int
	var s string
	var b bool
	return i, s, b
}

// Pi 展示常量声明的基本用法。
const Pi = 3.14

// DoublePi 返回 2 * Pi，用于在测试中验证常量的使用。
func DoublePi() float64 {
	return 2 * Pi
}

