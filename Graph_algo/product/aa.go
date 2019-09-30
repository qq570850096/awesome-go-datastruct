package main

func removeDuplicates(nums []int) int {
	if len(nums)<=2{
		return len(nums)
	}

	var k,count int = 0,1

	for i:=1;i<len(nums); {

		if nums[i]!=nums[k]{
			if count <= 2{
				nums[k+count] = nums[i]
				k = k+count
				count = 0
			} else {
				nums[k+2] = nums[i]
				count = 0
				k = k+2
			}

		}
		if nums[i] == nums[k] {
			count++
			i++
		}
	}
	return k+1
}

func main() {
	a := []int{1,1,2,4}
	removeDuplicates(a)
}
