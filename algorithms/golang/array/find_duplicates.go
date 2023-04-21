package array

func findDuplicate(nums []int) (res []int) {
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
