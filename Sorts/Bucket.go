package Sorts

func GradeSort(grade []int) {
	// 初始化出试卷范围[0-100]
	arr := make([]int,101)
	for _,v := range grade {
		// 每一个成绩装到对应的桶中
		arr[v]++
	}
	// 把成绩装回去
	index := 0
	for i,v := range arr {
		// 桶空了就继续下一个
		if v == 0 {
			continue
		}else {
			v--
			grade[index] = i
			index++
		}
	}
}

func main() {
	grade := []int{4,66,66,67,55,55,66,99,100,4,67}
}
