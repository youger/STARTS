package array

import (
	"fmt"
	"math/rand"
	"testing"
)

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
