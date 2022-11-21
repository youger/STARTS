package linklist

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func CreateLinklist(array []int) (head *ListNode) {
	pre := head
	for _, v := range array {
		node := ListNode{Val: v, Next: nil}
		if head == nil {
			head = &node
			pre = head
		}
		pre.Next = &node
		pre = &node
	}
	return
}

func PrintLinklist(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val, "->")
		head = head.Next
	}
	fmt.Println("nil")
}
