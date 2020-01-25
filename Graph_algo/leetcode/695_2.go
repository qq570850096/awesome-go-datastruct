package leetcode

// LeetCode695第二种解法，就地解决

func maxAreaOfIsland2(grid [][]int) int {
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

	// 开辟布尔型空间，和矩阵一一对应
	visited := make([][]bool,r)
	for i := 0; i < r; i++ {
		visited[i] = make([]bool,c)
	}
	res := 0
	var dfs func(x,y int) int
	dfs = func(x,y int) int {
		visited[x][y] = true
		tmp := 1
		for i := 0; i < 4; i++ {
			newX,newY := x + dirs[i][0], y + dirs[i][1]
			// 如果新点在矩阵中，且未被遍历过，且他是陆地，就可继续深度遍历
			if inArea(newX,newY) && !visited[newX][newY] && grid[newX][newY] == 1 {
				tmp += dfs(newX,newY)
			}
		}
		return tmp
	}
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if !visited[i][j] && grid[i][j] == 1 {
				res = Max(res,dfs(i,j))
			}
		}
	}
	return res
}

func Max (i, j int) int {
	if i > j {
		return i
	}
	return j
}