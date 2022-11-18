package main

import (
	"algorithms/linklist"
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
	var head *linklist.ListNode = nil
	pre := head
	array := []int{1, 3, 3, 2, 3, 3, 1}
	for _, v := range array {
		node := linklist.ListNode{Data: v, Next: nil}
		if head == nil {
			head = &node
			pre = head
		}
		pre.Next = &node
		pre = &node
	}
	// reverseLinklistWithLog(&head)
	// head.OddEvenList2()
	printLinklist(head)
	// node := head.ReverseLinklistWithCertainDistance2(4, 7)
	// printLinklist(node)
	result := head.PalindromeLinklist()
	fmt.Println(result)
}

func maxNumbersWithLog(numbers []int, k int) {
	fmt.Println("Get ", numbers, "in", k)
	ret := max_number_in_slide_window(numbers, k)
	fmt.Println(ret)
	fmt.Println("======================")
}

func threeNumbersSumWithLog(numbers []int, target int) {
	fmt.Println("Get ", numbers)
	ret := threeNumbersSum(numbers, 0)
	fmt.Println(ret)
	fmt.Println("======================")
}

func sortStringWithLog(s string) {
	fmt.Println("Sort " + s + ":")
	result := sorted(s)
	fmt.Println(result)
	fmt.Println("======================")
}

func reverseLinklistWithLog(head *linklist.ListNode) {
	printLinklist(head)
	head = head.ReverseKGroup(2)
	printLinklist(head)
}

func printLinklist(head *linklist.ListNode) {
	for head != nil {
		fmt.Print(head.Data, "->")
		head = head.Next
	}
	fmt.Println("nil")
}
