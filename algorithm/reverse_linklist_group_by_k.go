package main

type Node struct {
	data int
	next *Node
}

func (head *Node) reverseLinklist() (node *Node) {
	node = head
	if head.next != nil {
		node = head.next.reverseLinklist()
		head.next.next = head
		head.next = nil
	}
	return node
}

func (head *Node) reverseLinklistToNode(tail *Node) {
	if head != tail {
		head.next.reverseLinklistToNode(tail)
		head.next.next = head
		head.next = nil
	}
}

func (head *Node) reverseKGroup(k uint) (newHead *Node) {
	if k <= 1 {
		return head
	}
	node := head
	count := uint(1)
	var lastTail *Node = nil
	for node != nil {
		count++
		node = node.next
		if count%k == 0 && node != nil {
			next := node.next
			head.reverseLinklistToNode(node)
			if lastTail != nil {
				lastTail.next = node
			}
			lastTail = head
			head.next = next
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
