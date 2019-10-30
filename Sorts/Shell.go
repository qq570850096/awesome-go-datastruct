package Sorts

func ShellSort(arr []int,n int)  {
	h := 1
	for h < n/3 {
		h = 3 * h + 1
		// 计算 increment sequence: 1, 4, 13, 40, 121, 364, 1093...
	}
	for h >= 1 {
		for i := h; i < n ; i++ {
			// 对 arr[i], arr[i-h], arr[i-2*h], arr[i-3*h]... 使用插入排序
			e := arr[i]
			var j int
			for j = i; j>=h && e < arr[j-h]; j-=h {
				arr[j] = arr[j-h]
			}
			arr[j] = e
		}
		h /= 3
	}
}
