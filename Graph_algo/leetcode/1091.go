package leetcode

func shortestPathBinaryMatrix(grid [][]int) int {
	dirs := [8][4]int{{-1,0},{-1,1},{0,1},{1,1},
		{1,0},{1,-1},{0,-1},{-1,-1}}
	R,C := len(grid), len(grid[0])
	visited := make([][]bool,R)
	dis := make([][]int,R)
	for i := range visited{
		visited[i] = make([]bool,C)
		dis[i] = make([]int, C)
	}
	if grid[0][0] == 1 {
		return -1
	}
	if R == 1 && C == 1 {
		return 1
	}
	var InArea func(x,y int) bool
	InArea = func(x, y int) bool {
		return x >= 0 && x < R && y >= 0 && y < C
	}
	queue := make([]int,0)
	// 将二维数组映射到一维中, 则第一个点为0
	queue = append(queue,0)
	visited[0][0] = true
	dis[0][0] = 1
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		// 一维长度转二维坐标
		curX,curY := cur/C, cur%C
		for i:=0;i<8;i++ {
			newX,newY := curX + dirs[i][0], curY + dirs[i][1]
			if InArea(newX, newY) && !visited[newX][newY] &&
				grid[newX][newY] == 0{
				queue = append(queue,newX*C + newY)
				visited[newX][newY] = true
				dis[newX][newY] = dis[curX][curY] + 1

				if newX == R-1 && newY == C-1 {
					 return dis[newX][newY]
				}
			}
		}
	}
	return -1
}


