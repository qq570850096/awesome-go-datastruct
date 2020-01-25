package leetcode

func isBipartite(graph [][]int) bool {
	v := len(graph)
	visited := make([]bool,v+1)
	colors := make([]int,v+1)
	for i:=0; i < v ; i++ {
		if !visited[i] {
			if !dfs(i, 0,graph,colors,visited) {
				return false
			}
		}
	}
	return true
}

// 从顶点v开始，判断图中是否有环
func dfs(v int, color int, graph [][]int, colors []int,visited []bool) bool {
	visited[v] = true
	colors[v] = color
	for _, w := range graph[v] {
		if !visited[w] {
			if !dfs(w, 1-color,graph,colors,visited) {
				return false
			}
		} else if colors[w] == colors[v] {
			return false
		}
	}
	return true
}