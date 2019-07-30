package main

import "fmt"

func main()  {
	a := []int{4,4}
	length := removeElement(a, 4)
	fmt.Println(length)
	fmt.Println(a[:length])
}

/*

https://leetcode.com/problems/remove-element/submissions/

这题也做得不错
Runtime: 0 ms, faster than 100.00% of Go online submissions for Remove Element.
Memory Usage: 2.3 MB, less than 97.18% of Go online submissions for Remove Element.

*/

func removeElement(nums []int, val int) int {
	if len(nums) < 1 {
		return 0
	}

	if len(nums) == 1 && nums[0] == val {
		nums = nil
		return 0
	}

	fmt.Println(len(nums))
	tail := len(nums) - 1

	for i:=0;i<=tail; {
		if nums[i] == val {
			nums[i], nums[tail] = nums[tail], nums[i]
			tail--
		} else if i == tail && nums[i] == nums[tail] {
			i++
		} else {
			i++
		}
	}
	return tail+1
}