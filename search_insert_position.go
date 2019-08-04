package main

import "fmt"

func main() {

	fmt.Println(searchInsert([]int{1,3,5,6}, 7))
}

/*
Runtime: 4 ms, faster than 90.72% of Go online submissions for Search Insert Position.
Memory Usage: 3.1 MB, less than 97.67% of Go online submissions for Search Insert Position.

这题很简单
https://leetcode.com/problems/search-insert-position/submissions/

*/


func searchInsert(nums []int, target int) int {

	for i,v := range nums {
		if v == target {
			return i
		}
		if v > target && i != 0 {
			return i
		} else if v > target && i == 0 {
			return 0
		}
	}

	return len(nums)
}


