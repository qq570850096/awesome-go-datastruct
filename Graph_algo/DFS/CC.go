package DFS

import (
	"Graph_algo/Adj"
	"fmt"
)
// 深度优先遍历一张图，并且找到所有连通分量
type CC struct {
	graph *Adj.Hash
	visited []int
	cccount int
}

func (C *CC) Cccount() int {
	for _,v := range C.visited {
		fmt.Printf("%d ",v)
	}
	fmt.Printf("\n")
	return C.cccount
}

func (C *CC) Init(graph *Adj.Hash){
	C.graph = graph
	C.visited = make([]int,C.graph.V())
	for i,_ := range C.visited {
		C.visited[i]= -1
	}
	for i:=0;i<C.graph.V();i++ {
		if C.visited[i] == -1 {
			C.Dfs(i,C.cccount)
			C.cccount++
		}
	}
}
func (C *CC) Dfs (v int, cccountid int) {
	C.visited[v] = cccountid

	for _,w := range C.graph.LinkedVertex(v) {
		if C.visited[w] == -1 {
			C.Dfs(w,cccountid)
		}
	}
}

