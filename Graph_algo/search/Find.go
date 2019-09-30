package search

// 寻找单源路径
type Path struct {
	*SingleSource
	T int
}

func (C *Path) Visited() []bool {
	return C.visited
}

func (C *Path) Dfs(v int, parent int) bool {
	C.visited[v] = true
	C.pre[v] = parent
	if v == C.T {
		return true
	}
	for _, w := range C.graph.LinkedVertex(v) {
		if !C.visited[w] {
			if C.Dfs(w, v) {
				return false
			}
		}
	}
	return false
}

func (C *Path) IsConnectedTo() bool {
	return C.visited[C.T]
}

func (C *Path) Path() (res []int) {
	if !C.IsConnectedTo() {
		return
	}
	cur := C.T
	for cur != C.s {
		res = append(res, cur)
		cur = C.pre[cur]
	}
	res = append(res, C.s)
	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-i-1] = res[len(res)-i-1], res[i]
	}
	return
}
