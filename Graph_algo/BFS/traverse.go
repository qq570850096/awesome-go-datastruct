package BFS

import (
	"algo/Graph_algo/Adj"
	"fmt"
)

// 广度优先遍历取所有节点
func Traverse(hash *Adj.Hash, start int) (order []int) {
	visited := make([]bool, hash.V())
	bfs := func(source int) {
		que := []int{source}
		visited[source] = true
		for len(que) > 0 {
			temp := que[0]
			que = que[1:]
			order = append(order, temp)
			for _, v := range hash.LinkedVertex(temp) {
				if !visited[v] {
					visited[v] = true
					que = append(que, v)
				}
			}
		}
	}
	if err := hash.ValidateVertex(start); err != nil {
		panic(err)
	}
	bfs(start)
	for i := 0; i < hash.V(); i++ {
		if !visited[i] {
			bfs(i)
		}
	}
	fmt.Println(visited)
	return
}
