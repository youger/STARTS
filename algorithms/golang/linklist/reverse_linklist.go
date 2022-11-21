package linklist

func (head *ListNode) ReverseLinklist() (node *ListNode) {
	node = head
	if head.Next != nil {
		node = head.Next.ReverseLinklist()
		head.Next.Next = head
		head.Next = nil
	}
	return node
}

func (head *ListNode) reverseLinklistToNode(tail *ListNode) {
	if head != tail {
		head.Next.reverseLinklistToNode(tail)
		head.Next.Next = head
		head.Next = nil
	}
}

func (head *ListNode) ReverseKGroup(k uint) (newHead *ListNode) {
	if k <= 1 {
		return head
	}
	node := head
	count := uint(1)
	var lastTail *ListNode = nil
	for node != nil {
		count++
		node = node.Next
		if count%k == 0 && node != nil {
			next := node.Next
			head.reverseLinklistToNode(node)
			if lastTail != nil {
				lastTail.Next = node
			}
			lastTail = head
			head.Next = next
			head = next
			if count == k {
				newHead = node
			}
			node = next
			count++
		}
	}
	return newHead
}

func (head *ListNode) ReverseLinklistWithCertainDistance(left, right uint) *ListNode {
	if left >= right {
		return head
	}
	var count uint = 1
	node := head
	var preStart *ListNode = nil
	var endNext *ListNode = nil
	startNode := head
	endNode := head
	for node != nil {
		if count == left-1 {
			preStart = node
		}
		if left == count {
			startNode = node
		}
		if right == count {
			endNode = node
			endNext = node.Next
			break
		}
		node = node.Next
		count++
	}
	startNode.reverseLinklistToNode(endNode)
	startNode.Next = endNext
	if preStart != nil {
		preStart.Next = endNode
		return head
	} else {
		return endNode
	}
}

func (head *ListNode) ReverseLinklistWithCertainDistance2(left, right uint) *ListNode {
	if left >= right {
		return head
	}
	var count uint = 1
	node := head
	var preStart *ListNode = nil
	startNode := head
	endNode := head
	for node != nil {
		if count == left-1 {
			preStart = node
		}
		if left == count {
			startNode = node
			endNode = node
		}
		if count > left && count <= right {
			if preStart != nil {
				preStart.Next = node
			}
			endNode.Next = node.Next
			node.Next = startNode
			startNode = node
			if count == right {
				break
			}
			node = endNode
		}
		node = node.Next
		count++
	}
	if preStart != nil {
		return head
	} else {
		return startNode
	}
}
