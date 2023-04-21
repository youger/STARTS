package array

func QuickSort(nums []int) (res []int) {
	sorted(nums, 0, len(nums)-1)
	return res
}

func sorted(nums []int, left int, right int) {
	anchor := left
	for i, v := range nums[left:right] {
		if v < nums[anchor] {
			Swap(&v, &nums[anchor])
		}
		anchor = i
	}
}

func Swap(a *int, b *int) {

}
