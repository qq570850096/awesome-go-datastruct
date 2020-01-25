package leetcode

func openLock(deadends []string, target string) int {
	deadSet := make(map[string]bool)

	for _,v := range deadends  {
		deadSet[v] = true
	}

	if deadSet[target] == true {
		return -1
	}

	if deadSet["0000"] == true {
		return -1
	}

	if target == "0000" {
		return 0
	}

	// BFS
	queue := make([]string,0)
	queue = append(queue, "0000")
	visited := make(map[string]int)
	visited["0000"] = 0
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		next := make([]string,0)
		var curs []byte
		curs = []byte(cur)
		// 将next字符串依次找到
		for i,v := range curs  {
			curs[i] = ( v - '0' + 1 )%10 + '0'
			next = append(next,string(curs))

			curs[i] = v
			curs[i] = (v - '0' + 9) % 10 + '0'
			next = append(next,string(curs))
			curs[i] = v

		}
		// TODO:Next

		for _,v := range next  {
			if _,ok := visited[v];!ok && !deadSet[v] {
				queue = append(queue,v)
				visited[v] = visited[cur] + 1
				if v == target {
					return visited[v]
				}
			}
		}
	}
	return -1
}