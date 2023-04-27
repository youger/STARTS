package array

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestHIndex(t *testing.T) {
	index := hIndex([]int{4, 3, 2, 7})
	fmt.Println(index)
}

func TestQuickSort(t *testing.T) {
	//[81 87 47 59 81 18 25 40 56 0]
	nums := make([]int, 0)
	for i := 0; i < 10; i++ {
		nums = append(nums, rand.Intn(100))
	}
	fmt.Println(nums)
	quickSort(nums)
	fmt.Println("afert sorted:", nums)
}

func TestFindDuplicate(t *testing.T) {

	nums := findDuplicates([]int{4, 3, 2, 7, 8, 2, 3, 1})
	fmt.Println(nums)
}

func TestMinMoves(t *testing.T) {
	step := minMoves([]int{4, 3, 2, 7})
	fmt.Println(step)
}

func TestRotateArray(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8}
	rotate(nums, 3)
	fmt.Println(nums)
}
