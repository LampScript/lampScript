package main

import "fmt"

func main()  {
	a := []int{0,0,1,1,1,2,2,3,3,4}
	length := removeDuplicates(a)
	fmt.Println(length)
	fmt.Println(a[:length])
}

/*
https://leetcode.com/problems/remove-duplicates-from-sorted-array/submissions/

Runtime: 40 ms, faster than 98.17% of Go online submissions for Remove Duplicates from Sorted Array.
Memory Usage: 7.7 MB, less than 24.03% of Go online submissions for Remove Duplicates from Sorted Array.
这题做得比较好，当然题目也简单
*/

func removeDuplicates(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}

	length := 1
	head := 0
	for i:=1;i<len(nums);i++ {
		if nums[i] == nums[head] {
			continue
		} else {
			nums[head+1] = nums[i]
			head ++
			length++
		}

	}
	return length
}