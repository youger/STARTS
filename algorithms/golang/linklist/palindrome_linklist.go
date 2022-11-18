package linklist

func (head *ListNode) PalindromeLinklist() (result bool) {
	if head != nil {
		_, _, result = recursion(head, head)
	}
	return result
}

func recursion(head *ListNode, cur *ListNode) (*ListNode, *ListNode, bool) {
	if cur.Next != nil {
		headNext, tailNext, result := recursion(head, cur.Next)
		if result == false {
			return nil, nil, false
		}
		if headNext == tailNext {
			return headNext, tailNext, true
		} else if headNext.Data == tailNext.Data {
			return headNext.Next, cur, true
		} else {
			return nil, nil, false
		}
	} else {
		return head, cur, true
	}
}
