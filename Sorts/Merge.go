package Sorts

func merge(arr, aux []int, l, mid, r int) {

	aux = make([]int, r-l+1)
	for i := l; i <= r; i++ {
		aux[i-l] = arr[i]
	}

	i, j := l, mid+1
	for k := l; k <= r; k++ {

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

func MergeSort(arr []int, l, r int) {
	if l >= r {
		return
	}
	mid := (r + l) / 2
	MergeSort(arr, l, mid)
	MergeSort(arr, mid+1, r)

	if arr[mid] > arr[mid+1] {
		var aux []int
		merge(arr, aux, l, mid, r)
	}
}

func MergeSortBU(arr []int, n int) {
	aux := make([]int, n)

	for sz := 1; sz <= n; sz += sz {
		for i := 0; i < n-sz; i += sz + sz {
			if arr[i+sz-1] > arr[i+sz] {
				merge(arr, aux, i, i+sz-1, min(i+sz+sz-1, n-1))
			}
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
