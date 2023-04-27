package array

func rotate(nums []int, k int) {
	count := len(nums)
	end := count - 1
	step := k % count
	if step <= 0 {
		return
	}
	reverse(nums, 0, count-step-1)
	reverse(nums, count-step, end)
	reverse(nums, 0, end)
}

func reverse(nums []int, start int, end int) {
	for end > start {
		tmp := nums[start]
		nums[start] = nums[end]
		nums[end] = tmp
		start++
		end--
	}
}

func rotate2(nums []int, k int) {
	count := len(nums)
	step := k % count
	remain := count - step
	// shift left
	if step > remain {

	}
	reverseBlock(nums, 0, count, k)
	// shift right
	// start := 0
	// for remain > 0 {
	// 	for i := start; i < step; i++ {
	// 		tmp := nums[i]
	// 		nums[i] = nums[remain+i]
	// 		nums[remain+i] = tmp
	// 	}
	// 	start = step
	// 	step = remain - step
	// 	remain = remain + step
	// }
}

func reverseBlock(nums []int, start int, size int, blockSize int) {
	if size <= 1 {
		return
	}
	distance := size - blockSize
	for i := 0; i < blockSize; i++ {
		relativeStart := start + i
		symmetricStart := distance + relativeStart
		tmp := nums[relativeStart]
		nums[relativeStart] = nums[symmetricStart]
		nums[symmetricStart] = tmp
	}
	size -= blockSize
	if size < 2*blockSize {
		blockSize = size - blockSize
	}
	reverseBlock(nums, start, size, blockSize)
}

func rotate1(nums []int, k int) {
	count := len(nums)
	step := k % count
	remain := count - step
	res := nums[step:]
	for i := count - 1; i > step; i-- {
		nums[i] = nums[i-remain]
	}
	for i := 0; i < step; i++ {
		nums[i] = res[i]
	}
}
