package Sorts

func GradeSort(grade []int) {
	// 初始化出试卷范围[0-100]
	arr := make([]int, 101)
	for _, v := range grade {
		// 每一个成绩装到对应的桶中
		arr[v]++
	}
	// 把成绩装回去
	index := 0
	for i, count := range arr {
		// 桶空了就继续下一个
		if count == 0 {
			continue
		}
		for count > 0 {
			grade[index] = i
			index++
			count--
		}
	}
}
