package linklist

import (
	"fmt"
	"testing"
)

func TestReverseLinklist(t *testing.T) {
	array := []int{1, 2, 3, 4, 5, 6, 7}
	head := CreateLinklist(array)
	head = head.ReverseLinklist()
	PrintLinklist(head)
}

func TestReverseLinklistBetween(t *testing.T) {
	array := []int{1, 2, 3, 4, 5, 6, 7}
	head := CreateLinklist(array)

	node := head.ReverseLinklistWithCertainDistance(2, 4)
	//node := head.ReverseLinklistWithCertainDistance2(4, 7)
	PrintLinklist(node)
}

func TestOddEvenList(t *testing.T) {
	array := []int{1, 2, 3, 4, 5, 6, 7}
	head := CreateLinklist(array)

	// head.OddEvenList()
	head.OddEvenList2()
	PrintLinklist(head)
}

func TestPalindromeLinklist(t *testing.T) {
	array := []int{1, 3, 3, 2, 3, 3, 1}
	head := CreateLinklist(array)
	result := head.PalindromeLinklist()
	fmt.Println(result)
}
