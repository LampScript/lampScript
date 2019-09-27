package main


func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	l3 := new(ListNode)
	initNode := l3
	for true {
		if l1 == nil && l2 == nil {
			break
		}
		if l1 == nil && l2 != nil {
			l3.Val = l2.Val
			l3.Next = l2.Next
			break
		}

		if l2 == nil && l1 != nil {
			l3.Val = l1.Val
			l3.Next = l1.Next
			break
		}

		if l1.Val < l2.Val {
			l3.Val = l1.Val
			l3.Next = new(ListNode)
			l1 = l1.Next
			l3 = l3.Next
		} else {
			l3.Val = l2.Val
			l3.Next = new(ListNode)
			l2 = l2.Next
			l3 = l3.Next
		}
	}
	return initNode
}
