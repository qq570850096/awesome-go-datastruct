package Sorts

import "testing"

func BenchmarkSort(t *testing.B) {
	var arr []int
	for i := 9; i >= 0; i-- {
		arr = append(arr, i)
	}
	t.Log(arr)

	ShellSort(arr, len(arr))
	t.Log(arr)
}
