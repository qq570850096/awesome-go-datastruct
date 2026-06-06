package Adj

func indexOfIntSlice(items []int, target int) int {
	for index, item := range items {
		if item == target {
			return index
		}
	}
	return -1
}
