package Sorts

import "testing"

func TestSort(t *testing.T) {
	var arr []int
	for i:=9 ; i>=0;i--{
		arr = append(arr,i)
	}
	t.Log(arr)
	//MergeSort(arr,0,len(arr)-1)
	ShellSort(arr,len(arr))
	t.Log(arr)
}
