package Sorts

func BubbleSort(arr []int)  {
	if arr == nil || len(arr) == 1{
		return
	}
	
	for i := range arr {
		for j := len(arr)-1;j>i;j-- {
			if arr[j] < arr[j-1] {
				arr[j],arr[j-1] = arr[j-1],arr[j]
			}
		}
	}
}

