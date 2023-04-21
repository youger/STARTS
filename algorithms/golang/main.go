package main

import (
	"algorithms/array"
	"fmt"
)

func main() {
	// sortStringWithLog("ABC")
	// sortStringWithLog("ABABC")
	// sortStringWithLog("ABDBC")

	// threeNumbersSumWithLog([]int{-1, 0, 1, 2, -1, -4}, 0)
	// threeNumbersSumWithLog([]int{}, 0)
	// threeNumbersSumWithLog([]int{0}, 0)

	// maxNumbersWithLog([]int{}, 2)
	// maxNumbersWithLog([]int{1, 3, 3, 2, -1, 5}, 3)
	// maxNumbersWithLog([]int{1, 3, 3, 2, -1, 5}, 2)
	// maxNumbersWithLog([]int{1, 3, 3, 2, -1, 5}, 9)
	// maxNumbersWithLog([]int{1, 3, 3, 2, -1, 5}, 1)
	// maxNumbersWithLog([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3)
}

func maxNumbersWithLog(numbers []int, k int) {
	fmt.Println("Get ", numbers, "in", k)
	ret := max_number_in_slide_window(numbers, k)
	fmt.Println(ret)
	fmt.Println("======================")
}

func threeNumbersSumWithLog(numbers []int, target int) {
	fmt.Println("Get ", numbers)
	ret := array.ThreeNumbersSum(numbers, 0)
	fmt.Println(ret)
	fmt.Println("======================")
}

func sortStringWithLog(s string) {
	fmt.Println("Sort " + s + ":")
	result := sorted(s)
	fmt.Println(result)
	fmt.Println("======================")
}
