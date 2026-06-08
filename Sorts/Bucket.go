package Sorts

func GradeSort(grade []int) {

	arr := make([]int, 101)
	for _, v := range grade {

		arr[v]++
	}

	index := 0
	for i, count := range arr {

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
