package stack

type Stack struct {
	container []byte

	top int

	size int
}

func NewStack(size int) Stack {
	return Stack{
		container: make([]byte, size),
		top:       0,
		size:      size,
	}
}

func (s *Stack) Push(e byte) bool {
	if s.IsFull() {
		return false
	}
	s.container[s.top] = e
	s.top++
	return true
}

func (s *Stack) Pop() (flag bool, ret byte) {

	if s.IsEmpty() {
		return false, ret
	}
	ret = s.container[s.top-1]
	s.top--
	return true, ret
}

func (s *Stack) IsEmpty() bool {
	if s.top == 0 {
		return true
	}
	return false
}

func (s *Stack) IsFull() bool {
	if s.top == s.size {
		return true
	}
	return false
}

func IsValid(s string) bool {
	stack := NewStack(100)

	for _, v := range s {
		if v == '(' {

			stack.Push(byte(v))
		}
		if v == ')' {
			if flag, top := stack.Pop(); flag == true && top == '(' {
				continue
			} else {
				return false
			}
		}
	}

	if stack.IsEmpty() {
		return true
	}

	return false
}
