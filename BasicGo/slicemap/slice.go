package slicemap

func MakeNilAndEmpty() (nilSlice, emptySlice []int) {
	var s []int
	e := make([]int, 0)
	return s, e
}

func ShareUnderlying() (base, sub, grown []int) {
	base = []int{1, 2, 3, 4}
	sub = base[:2]
	sub[0] = 10
	grown = append(sub, 99)
	return base, sub, grown
}

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
