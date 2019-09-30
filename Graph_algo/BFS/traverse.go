package BFS

import (
	"Graph_algo/Adj"
	"fmt"
)

// 广度优先遍历取所有节点
func Traverse(hash *Adj.Hash, start int) (order []int) {
	var (
		visited []bool
		que     []int
		temp    int
	)
	visited = make([]bool, hash.V())
	if err := hash.ValidateVertex(start); err != nil {
		panic(err)
	}
	que = append(que, start)
	visited[start] = true
	// 最外层循环保证图中有多个联通分量也可以访问到所有节点
	for i := 0; i < hash.V(); i++ {
		if !visited[i] {
			que = append(que, i)
			visited[i] = true
		}
		for len(que) > 0 {
			temp = que[0]
			order = append(order, temp)
			que = que[1:]
			for _, v := range hash.LinkedVertex(temp) {
				if !visited[v] {
					que = append(que, v)
					visited[v] = true
				}
			}
		}
	}
	fmt.Println(visited)
	return
}
