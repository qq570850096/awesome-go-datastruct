package slicemap

// MakeNilAndEmpty 返回一个 nil 切片和一个空切片，用于展示它们的区别。
func MakeNilAndEmpty() (nilSlice, emptySlice []int) {
	var s []int      // nil 切片
	e := make([]int, 0) // 空切片，非 nil
	return s, e
}

// ShareUnderlying 展示切片共享底层数组时的联动效应。
// 返回的三个切片共享一部分底层数组。
func ShareUnderlying() (base, sub, grown []int) {
	base = []int{1, 2, 3, 4}
	sub = base[:2]
	sub[0] = 10
	grown = append(sub, 99)
	return base, sub, grown
}

// FilterInPlace 使用原地覆盖的方式过滤切片内容，避免额外分配。
func FilterInPlace(nums []int, keep func(int) bool) []int {
	j := 0
	for _, v := range nums {
		if keep(v) {
			nums[j] = v
			j++
		}
	}
	return nums[:j]
}

