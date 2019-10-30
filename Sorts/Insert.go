package Sorts

func InsertSort (arr []int) {
	// 对于切片，他本身就是引用传递，所以可以用是否为nil判断
	if arr == nil {
		return
	}
	for i:=1; i<len(arr);i++{
		tmp,j := arr[i],i
		if arr[j-1] > tmp {
			for j>=1 && arr[j-1]>tmp {
				arr[j] = arr[j-1]
				j--
			}
		}
		arr[j] = tmp
	}
}
