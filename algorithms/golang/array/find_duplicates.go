package array

func findDuplicates(nums []int) (res []int) {
	res = make([]int, 0)
	count := len(nums)
	flag := make(map[int]int, 0)
	for i := 0; i < count; i++ {
		flag[nums[i]] += 1
	}
	for k, v := range flag {
		if v == 2 {
			res = append(res, k)
		}
	}
	return
}

func findDuplicates2(nums []int) (res []int) {
	res = make([]int, 0)
	for _, v := range nums {
		if v < 0 {
			v = -v
		}
		index := v - 1
		if nums[index] > 0 {
			nums[index] = -nums[index]
		} else {
			res = append(res, v)
		}
	}
	return
}

func findDuplicates3(nums []int) (ans []int) {
	for i := range nums {
		for nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}
	for i, num := range nums {
		if num-1 != i {
			ans = append(ans, num)
		}
	}
	return
}
