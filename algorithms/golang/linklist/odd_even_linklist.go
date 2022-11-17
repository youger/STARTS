package linklist

func (head *ListNode) OddEvenList() {
	if head == nil {
		return
	}
	node := head
	queue := make([]*ListNode, 0)
	for node != nil {
		if node.Next != nil {
			queue = append(queue, node.Next)
			node.Next = node.Next.Next
		} else {
			break
		}
		if node.Next != nil {
			node = node.Next
		} else {
			break
		}
	}

	for _, v := range queue {
		node.Next = v
		node.Next.Next = nil
		node = node.Next
	}
}

func (head *ListNode) OddEvenList2() {
	if head == nil {
		return
	}
	odd := head
	even := head.Next
	for even != nil && even.Next != nil {
		odd.Next = even.Next
		odd = odd.Next
		even.Next = odd.Next
		even = even.Next
	}
}
