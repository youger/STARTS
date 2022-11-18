package linklist

func (head *ListNode) PalindromeLinklist() (result bool) {
	if head != nil {
		_, _, result = recursivelyCheck(head, head)
	}
	return result
}

func recursivelyCheck(head *ListNode, cur *ListNode) (*ListNode, *ListNode, bool) {
	if cur.Next != nil {
		headNext, tailNext, result := recursivelyCheck(head, cur.Next)
		if result == false {
			return nil, nil, false
		}
		if headNext == tailNext {
			return headNext, tailNext, true
		} else if headNext.Val == tailNext.Val {
			return headNext.Next, cur, true
		} else {
			return nil, nil, false
		}
	} else {
		return head, cur, true
	}
}
