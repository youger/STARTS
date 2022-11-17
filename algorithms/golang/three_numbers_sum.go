package main

func twoNumbersSum(numbers []int, target int) (ret [][]int) {
	count := len(numbers)
	ret = make([][]int, 0)
	if count < 2 {
		return
	}
	hashTable := make(map[int]bool, count)
	for _, val := range numbers {
		hashTable[val] = true
	}
	flags := make(map[int]bool, count)
	for _, val := range numbers {
		if flags[val] {
			continue
		}
		flags[val] = true
		val2 := target - val
		if hashTable[val2] && flags[val2] == false {
			ret = append(ret, []int{val, val2})
			flags[val2] = true
		}
	}
	return
}

func threeNumbersSum(numbers []int, target int) (ret [][]int) {
	count := len(numbers)
	ret = make([][]int, 0)
	if count < 3 {
		return
	}
	flags := make(map[int]bool, count-1)
	for i, val := range numbers {
		if flags[val] {
			continue
		}
		flags[val] = true
		target2 := target - val
		ret2 := twoNumbersSum(remove(numbers, i), target2)
		for _, array := range ret2 {
			ret = append(ret, append(array, val))
		}
	}
	return
}

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}
