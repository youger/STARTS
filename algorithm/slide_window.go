package main

func max_number_in_slide_window(array []int, k int) []int {
	count := len(array)
	if count <= 0 {
		return []int{}
	}
	step := count - k
	if count < k {
		k = count
		step = 0
	}
	var res = make([]int, 0)
	for i := 0; i <= step; i++ {

		subArray := array[i : i+k]
		max := subArray[0]
		for j := 0; j < k; j++ {
			if max < subArray[j] {
				max = subArray[j]
			}
		}
		res = append(res, max)
	}
	return res
}

// func insertSort(array []int, insert int) []int {
// 	count := len(array)
// 	for i := 0; i < count; i++ {
// 		if insert > array[i] {

// 		}
// 	}
// }
