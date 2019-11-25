package main

import (
	"fmt"
)

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

type Animal interface {
	Eat(s string)
}

type Monster interface {
	Eat(s string)
}

type Monkey struct {

}

func (m Monkey) Eat(s string)  {
	fmt.Println(s)
}

func main()  {
	WhoIsThat(Monkey{})
	WhoIsThat2(Monkey{})
}

func WhoIsThat(a Animal)  {
	a.Eat("mmm")
}
func WhoIsThat2(a Monster)  {
	a.Eat("mmm")
}

func myaaa(a []int)  {
	a = append(a,1)
	fmt.Println(cap(a))
}


func sort222(a []int)  {
	if a == nil || len(a) <=1 {
		return
	}
	head :=1
	tail := len(a) - 1
	i := 0
	cur := a[0]

	for tail != i {
		if cur < a[head] {
			a[i],a[head] = a[head], a[i]
			i++
			head++
		} else {
			a[tail], a[head] = a[head], a[tail]
			tail--
		}
	}
	sort222(a[:i+1])
	sort222(a[i+1:])
}

func reverseListNode(node *ListNode)*ListNode  {
	if node == nil || node.Next == nil {
		return node
	}
	var last  *ListNode
	var next  *ListNode
	for node != nil {
		next = node.Next
		node.Next = last
		last = node
		node = next
	}
	return last
}