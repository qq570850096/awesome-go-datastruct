package generics

// Stack 展示最常见的泛型容器实现。
type Stack[T any] struct {
	data []T
}

func (s *Stack[T]) Push(v T) {
	s.data = append(s.data, v)
}

func (s *Stack[T]) Pop() (T, bool) {
	var zero T
	if len(s.data) == 0 {
		return zero, false
	}
	v := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return v, true
}

func (s *Stack[T]) Len() int {
	return len(s.data)
}

// MapSlice 展示 type parameter 组合：输入类型 T，输出类型 R。
func MapSlice[T any, R any](items []T, fn func(T) R) []R {
	res := make([]R, 0, len(items))
	for _, item := range items {
		res = append(res, fn(item))
	}
	return res
}

// FilterSlice 用 comparable 约束示例。
func FilterSlice[T comparable](items []T, keep func(T) bool) []T {
	out := items[:0]
	for _, item := range items {
		if keep(item) {
			out = append(out, item)
		}
	}
	return out
}
