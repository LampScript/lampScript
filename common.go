package main

import "fmt"

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

func main()  {
	a := GetListNode([]int{1,2,3,4,5})
	PrintListNode(a)

}