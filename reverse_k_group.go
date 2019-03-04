package main


func reverseKGroup(head *ListNode, k int) *ListNode {
	length := 0
	s := make([]*ListNode, 0, 10)
	temp := head
	for {
		if temp != nil {
			length++
			s = append(s, temp)
			temp = temp.Next
			continue
		}
		break
	}

	if k > length {
		return head
	}

	nS := make([]*ListNode, 0, len(s))
	n := length / k
	for i:= 0;i < n;i++  {
		begin := i * k
		end := i * k + k
		for i := end - 1; i >= begin; i--  {
			nS = append(nS, s[i])
		}
	}

	for i:= n*k; i < length ;i++  {
		nS = append(nS, s[i])
	}

	for i:=0;i < length ;i++  {
		if i == length -1 {
			nS[i].Next = nil
			continue
		}
		nS[i].Next = nS[i+1]
	}
	return nS[0]
}
