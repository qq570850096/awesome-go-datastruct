package Sorts

import "testing"

func BenchmarkSort(t *testing.B) {
	var arr []int
	for i:=9 ; i>=0;i--{
		arr = append(arr,i)
	}
	t.Log(arr)
	// 如需测试归并排序可调用：MergeSort(arr,0,len(arr)-1)
	ShellSort(arr,len(arr))
	t.Log(arr)
}
