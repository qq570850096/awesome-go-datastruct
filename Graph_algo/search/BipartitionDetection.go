package search

import "Graph_algo/Adj"

type BipartitionDetection struct {
	graph     *Adj.Hash
	visited   []bool
	color     []int
	isBippart bool
}

func (C *BipartitionDetection) IsBippart() bool {
	return C.isBippart
}

func (C *BipartitionDetection) Init(graph *Adj.Hash) {

	C.graph = graph
	C.visited = make([]bool, C.graph.V())
	C.color = make([]int, C.graph.V())
	C.isBippart = true
	for i, _ := range C.color {
		C.color[i] = -1
	}
	for i := 0; i < C.graph.V(); i++ {
		if !C.visited[i] {
			if !C.Dfs(i, 0) {
				C.isBippart = false
				break
			}
		}
	}

}

// 从顶点v开始，判断图中是否有环
func (C *BipartitionDetection) Dfs(v int, color int) bool {
	C.visited[v] = true
	C.color[v] = color
	for _, w := range C.graph.LinkedVertex(v) {
		if !C.visited[w] {
			if !C.Dfs(w, 1-color) {
				return false
			}
		} else if C.color[w] == C.color[v] {
			return false
		}
	}
	return true
}
