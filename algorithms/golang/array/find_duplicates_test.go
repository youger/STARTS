package array

import (
	"fmt"
	"testing"
)

func TestFindDuplicate(t *testing.T) {

	nums := findDuplicate([]int{4, 3, 2, 7, 8, 2, 3, 1})
	fmt.Println(nums)
}
