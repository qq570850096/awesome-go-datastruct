package search

import "Graph_algo/Adj"

type SingleSource struct {
	graph   *Adj.Hash
	visited []bool
	s       int
	pre     []int
}

func (C *SingleSource) Pre() []int {
	return C.pre
}

func (C *SingleSource) IsConnectedTo(t int) bool {
	if err := C.graph.ValidateVertex(t); err != nil {
		panic(err)
	}
	return C.visited[t]
}
func (C *SingleSource) Path(t int) (res []int) {
	if !C.IsConnectedTo(t) {
		return
	}
	cur := t
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

func (C *SingleSource) Init(graph *Adj.Hash, s int) {

	C.graph = graph
	if err := C.graph.ValidateVertex(s); err != nil {
		panic(err)
	}
	C.visited = make([]bool, C.graph.V())
	C.pre = make([]int, C.graph.V())
	for i, _ := range C.pre {
		C.pre[i] = -1
	}
	C.s = s
	C.Dfs(s, s)
}

func (C *SingleSource) Dfs(v int, parent int) {
	C.visited[v] = true
	C.pre[v] = parent
	for _, w := range C.graph.LinkedVertex(v) {
		if !C.visited[w] {
			C.Dfs(w, v)
		}
	}
}
