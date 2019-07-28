package main

func main()  {

}


/**
Given a linked list, remove the n-th node from the end of list and return its head.

Example:

Given linked list: 1->2->3->4->5, and n = 2.

After removing the second node from the end, the linked list becomes 1->2->3->5.
Note:

Given n will always be valid.

Follow up:

Could you do this in one pass?


Runtime: 0 ms, faster than 100.00% of Go online submissions for Remove Nth Node From End of List.
Memory Usage: 2.2 MB, less than 100.00% of Go online submissions for Remove Nth Node From End of List.

这题比较简单，就懒得看其他回答了

 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	length := getListLength(head)

	previous := getNodeByIndex(head, length - n - 1)
	next := getNodeByIndex(head, length - n + 1)

	if previous == nil && next == nil {
		return nil
	}
	if previous != nil {
		previous.Next = next
		return head
	}

	return next
}



func getListLength(list *ListNode) int  {
	l := list
	i := 0
	for true  {
		if l != nil {
			i++
			l = l.Next
		} else {
			break
		}
	}
	return i
}

func getNodeByIndex(list *ListNode, i int) *ListNode  {
	if i < 0 {
		return nil
	}
	l := list
	for true  {
		if l != nil && i > 0 {
			i--
			l = l.Next
		} else {
			break
		}
	}
	return l
}
