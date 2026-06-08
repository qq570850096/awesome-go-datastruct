package Sorts

func ShellSort(arr []int, n int) {
	h := 1
	for h < n/3 {
		h = 3*h + 1

	}
	for h >= 1 {
		for i := h; i < n; i++ {

			e := arr[i]
			var j int
			for j = i; j >= h && e < arr[j-h]; j -= h {
				arr[j] = arr[j-h]
			}
			arr[j] = e
		}
		h /= 3
	}
}
