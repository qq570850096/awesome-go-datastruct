package leetcode
// 如果你有无穷多的水，一个3 公升的提捅，一个5 公升的提捅，两只提捅形状上下都不均 匀，问你如何才能准确称出4 公升的水？
func FindPath(target int) []int  {
	queue := make([]int,0)
	visited := make([]bool,100)
	pre := make([]int,100)
	var result func(end int) []int
	result = func(end int) []int {
		if end == -1 {
			return nil
		}
		res := make([]int,0)
		cur := end
		for cur != 0 {
			res = append(res,cur)
			cur = pre[cur]
		}
		res = append(res, 0)
		reverse(res)
		return res
	}

	queue = append(queue, 0)
	visited[0] = true
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		curX,curY := cur/10, cur%10
		nexts := make([]int,0)
		// TODO: 找到next的六种情况
		nexts = append(nexts, 5*10 + curY)
		nexts = append(nexts, curX*10 + 3)
		nexts = append(nexts, 0*10 +curY)
		nexts = append(nexts, curX*10+0)
		x := min(3-curY,curX)
		nexts = append(nexts, (curX-x)*10+curY+x)
		y := min(5-curX,curY)
		nexts = append(nexts, (curX+y)*10+curY-y)
		for _,v := range nexts {
			if !visited[v] {
				queue = append(queue,v)
				pre[v] = cur
				visited[v] = true

				if v /10 == target || v%10 == target {
					return result(v)
				}
			}
		}
	}
	return nil
}

func min(x,y int) int {
	if x<y {
		return x
	}
	return  y
}

func reverse(arr []int) {
	for i,j := 0,len(arr)-1; i<j ;{
		arr[i],arr[j] = arr[j], arr[i]
		i++
		j--
	}
}
