package linklist

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	slow := head
	fast := head.Next
	for {
		if slow == fast {
			break
		}
		slow = slow.Next
		if fast == nil || fast.Next == nil {
			return nil
		}
		fast = fast.Next.Next
	}
	node := head
	for node != slow {
		node = node.Next
		slow = slow.Next
	}
	return node
}
