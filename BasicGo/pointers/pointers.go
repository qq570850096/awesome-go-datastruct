package pointers

type Node struct {
	Value int
	Next  *Node
}

func IncrementValue(n int) int {
	n++
	return n
}

func IncrementPointer(n *int) bool {
	if n == nil {
		return false
	}
	*n = *n + 1
	return true
}

func Swap(a, b *int) bool {
	if a == nil || b == nil {
		return false
	}
	*a, *b = *b, *a
	return true
}

func Link(values ...int) *Node {
	var head *Node
	for i := len(values) - 1; i >= 0; i-- {
		head = &Node{Value: values[i], Next: head}
	}
	return head
}

func Values(head *Node) []int {
	var result []int
	for node := head; node != nil; node = node.Next {
		result = append(result, node.Value)
	}
	return result
}
