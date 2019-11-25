package main

import "fmt"

func main()  {

}


func addTwoNumbers3(l1 *ListNode, l2 *ListNode) *ListNode {
	head := new(ListNode)
	cursor := head
	extra := 0

	for true {
		v := extra
		extra = 0
		if l1 != nil {
			v+=l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			v+=l2.Val
			l2=l2.Next
		}
		if v >= 10 {
			v-=10
			extra++
		}
		cursor.Val = v
		if l1 == nil && l2 == nil && extra == 0 {
			break
		}
		tmp := new(ListNode)
		cursor.Next = tmp
		cursor = tmp
	}

	return head
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	add := 0
	array := make([]*ListNode, 0)

	for true {
		if l1 == nil && l2 == nil && add == 0 {
			break
		} else if l1 == nil && l2 == nil && add > 0 {
			l3 := &ListNode{Val:add}
			array = append(array, l3)
			break
		}
		sum := 0
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		sum += add
		add = 0
		val := sum
		if sum >= 10 {
			val -= 10
			add = 1
		}
		l3 := &ListNode{Val:val}
		array = append(array, l3)
	}
	fmt.Println(len(array))
	for i:=0;i < len(array) - 1 ;i++  {
		array[i].Next = array[i+1]
	}
	if len(array) > 0 {
		return array[0]
	} else {
		return nil
	}
}

func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}
	current := result
	each_sums := 0

	for l1 != nil || l2 != nil || each_sums != 0{
		if l1 != nil{
			each_sums += l1.Val
			l1 = l1.Next
		}
		if l2 != nil{
			each_sums += l2.Val
			l2 = l2.Next
		}
		current.Next = &ListNode{Val: each_sums % 10}
		each_sums /= 10
		current = current.Next
	}

	return result.Next
}