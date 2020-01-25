package leetcode


func maxAreaOfIsland(grid [][]int) int {
	dirs := [4][2]int {{-1,0},{0,1},{1,0},{0,-1}}
	// 如果二维数组不存在，直接返回
	if len(grid) == 0 {
		return 0
	}

	r := len(grid)
	if r == 0 {
		return 0
	}
	c := len(grid[0])
	if c == 0 {
		return 0
	}
	// 判断一点是否在可选范围
	inArea := func(x,y int) bool {
		return x >= 0 && x < r && y >= 0 && y < c
	}
	// 构造一个图的邻接表
	constructGraph := func() map[int][]int {
		g := make(map[int][]int)
		for v := 0; v < r*c; v++ {
			x,y := v/c , v%c
			if grid[x][y] == 1 {
				for d := 0; d < 4; d++ {
					newX := x + dirs[d][0]
					newY := y + dirs[d][1]
					if inArea(newX, newY) && grid[newX][newY] == 1 {
						next := newX * c + newY
						g[v] = append(g[v],next)
						g[next] = append(g[next],v)
					}
				}
			}
		}
		return g
	}

	g := constructGraph()
	res := 0
	visited := make([]bool,r*c)
	var dfs func(v int) int
	dfs = func(v int) int {
		visited[v] = true
		tmp := 1
		for _,w := range g[v] {
			if !visited[w] {
				tmp += dfs(w)
			}
		}
		return tmp
	}
	for v := 0; v < r*c; v++ {
		x,y := v/c, v%c
		if !visited[v] && grid[x][y] == 1 {
			res = max(res,dfs(v))
		}
	}
	return res
}

func max (x,y int) int {
	if x > y {
		return x
	}
	return y
}


