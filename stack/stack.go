package stack

type Stack struct {
	// 用来装元素的切片
	container []byte
	// 栈顶标记位
	top int
	// 容量限制
	size int
}

// 初始化时，要传入容量限制
func NewStack(size int) Stack {
	return Stack{
		container: make([]byte,size),
		top:       0,
		size:      size,
	}
}

func (s *Stack) Push (e byte) bool {
	if s.IsFull() {
		return false
	}
	s.container[s.top] = e
	s.top++
	return true
}


func (s *Stack) Pop () (flag bool,ret byte) {
	// 如果栈是空的，那么就不能继续 Pop 了
	if s.IsEmpty() {
		return false,ret
	}
	ret = s.container[s.top-1]
	s.top--
	return true,ret
}


func (s *Stack) IsEmpty () bool {
	if s.top == 0 {
		return true
	}
	return false
}


func (s *Stack) IsFull () bool {
	if s.top == s.size {
		return true
	}
	return false
}

func IsValid(s string) bool {
	stack := NewStack(100)
	// 遍历括号字符串
	for _,v := range s {
		if v == '(' {
			// 由于golang中的字符串默认是unicode编码，所以我们要做一个强制类型转换
			stack.Push(byte(v))
		}
		if v == ')' {
			if flag,top := stack.Pop(); flag == true &&top == '(' {
				continue
			} else {
				return false
			}
		}
	}
	// 字符串遍历完后如果栈也空了，说明括号匹配
	if stack.IsEmpty() {
		return true
	}
	// 如果栈不空，说明栈里还有多余的左括号
	return false
}
