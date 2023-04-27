package array

func minMoves(nums []int) int {
	min := nums[0]
	for _, num := range nums {
		if min > num {
			min = num
		}
	}

	step := 0
	for _, num := range nums {
		step += num - min		
	}
	return step
}