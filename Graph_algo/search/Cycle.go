package search

import (
	"Graph_algo/Adj"
)

type Cycle struct {
	graph    *Adj.Hash
	visited  []bool
	hasCycle bool
}

func (C *Cycle) HasCycle() bool {
	return C.hasCycle
}

func (C *Cycle) Init(graph *Adj.Hash) {

	C.graph = graph
	C.hasCycle = false
	C.visited = make([]bool, C.graph.V())
	for i := 0; i < C.graph.V(); i++ {
		if !C.visited[i] {
			if C.Dfs(i, i) {
				C.hasCycle = true
				break
			}
		}
	}

}

// 从顶点v开始，判断图中是否有环
func (C *Cycle) Dfs(v int, parent int) bool {
	C.visited[v] = true
	for _, w := range C.graph.LinkedVertex(v) {
		if !C.visited[w] {
			if C.Dfs(w, v) {
				return true
			}
		} else if w != parent {
			return true
		}
	}
	return false
}
