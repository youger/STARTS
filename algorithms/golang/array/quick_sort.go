package array

func quickSort(nums []int) {
	sortArray(nums, 0, len(nums)-1)
}

func sortArray(nums []int, left int, right int) {
	if left >= right {
		return
	}
	anchor := nums[left]
	i := left
	j := right
	for i < j {
		for i < j && nums[j] >= anchor {
			j--
		}
		if i < j {
			nums[i] = nums[j]
		}
		for i < j && nums[i] < anchor {
			i++
		}
		if i < j {
			nums[j] = nums[i]
		}
	}
	nums[i] = anchor
	sortArray(nums, left, i-1)
	sortArray(nums, i+1, right)
}
