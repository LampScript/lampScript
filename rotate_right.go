package main

import "fmt"

func main() {
	a := GetListNode([]int{})

	c := rotateRight(a, 1)

	PrintListNode(c)

}

func PrintListNode(l *ListNode)  {
	t := l
	for true {
		if t == nil {
			fmt.Println()
			break
		}
		fmt.Printf("%d", t.Val)
		if t.Next != nil {
			fmt.Print("->")
		}
		t= t.Next
	}
}
type ListNode struct {
	Val  int
	Next *ListNode
}

func GetListNode(array []int) *ListNode {
	if len(array) == 0 {
		return nil
	}
	l := &ListNode{
		Val:array[0],
	}
	head := l
	for i, v:= range array {
		if i == 0 {
			continue
		}
		t := &ListNode{Val:v}
		l.Next = t
		l = l.Next
	}
	return head
}


func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	length := 1
	tmp := head
	for true {
		if tmp.Next != nil {
			length++
			tmp = tmp.Next
		} else {
			k = k % length
			if k == 0 {
				return head
			}
			tmp.Next = head
			newHead := head
			var last *ListNode
			for i := length - k ;i > 0 ;i--  {
				last = newHead
				newHead = newHead.Next
			}
			last.Next = nil
			return newHead
		}
	}
	return head
}