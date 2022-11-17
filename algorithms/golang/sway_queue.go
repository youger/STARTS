package main

func wiggleMaxLength(nums []int) int {
	count := len(nums)
	if count <= 1 {
		return 1
	}
	subCount := count
	lastDiff := nums[1] - nums[0]
	for i := 0; i < count-1; i++ {
		var diff = nums[i+1] - nums[i]
		if diff*lastDiff >= 0 {
			subCount--
		}
	}
	return subCount
}
