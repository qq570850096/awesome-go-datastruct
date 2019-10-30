package Sorts

// 将arr[l...mid]和arr[mid+1...r]两部分进行归并
func merge(arr,aux []int,l,mid,r int)  {
	//* 如果你使用的是c++，这里特别提醒一下
	//* VS不支持动态长度数组, 即不能使用 T aux[r-l+1]的方式申请aux的空间
	//* 使用VS的同学, 请使用new的方式申请aux空间
	//* 使用new申请空间, 不要忘了在merge函数的最后, delete掉申请的空间:)
	aux = make([]int,r-l+1)
	for i:=l;i<=r;i++ {
		aux[i-l] = arr[i]
	}
	// 初始化，i指向左半部分的起始索引位置l，j指向mid+1
	i,j := l,mid+1
	for k:=l ; k<=r; k++ {
		// 如果左半部分元素已经处理完毕
		if i > mid {
			arr[k] = aux[j-l]
			j++
		} else if j > r {
			arr[k] = aux[i-l]
			i++
		} else if aux[i-l] < aux[j-l] {
			arr[k] = aux[i-l]
			i++
		} else {
			arr[k] = aux[j-l]
			j++
		}
	}
}

func MergeSort(arr []int,l,r int)  {
	if l >= r {
		return
	}
	mid := (r+l)/2
	MergeSort(arr,l,mid)
	MergeSort(arr,mid+1,r)
	// 对于arr[mid] <= arr[mid+1]的情况，不进行merge
	// 对于近乎有序的数组非常有效，但是对于一般情况，有一定的性能损失。
	if arr[mid] > arr[mid+1] {
		var aux []int
		merge(arr,aux,l,mid,r)
	}
}

// 自底向上的归并排序
func MergeSortBU(arr []int,n int)  {
	aux := make([]int,n)
	// 一次性申请aux的空间，并将这个辅助空间以参数形式传递给完成归并排序的各个子函数
	for sz := 1 ; sz <= n; sz += sz {
		for i:=0;i<n-sz;i+=sz+sz {
			if arr[i + sz -1] > arr[i+sz]{
				merge(arr,aux,i,i+sz-1,min(i+sz+sz-1,n-1))
			}
		}
	}
}

func min (a,b int ) int {
	if a < b {
		return a
	}
	return b
}
