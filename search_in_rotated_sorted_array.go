package main

import "fmt"

func main()  {
	fmt.Println(search([]int{1,2,3}, 1))
}



/*
Suppose an array sorted in ascending order is rotated at some pivot unknown to you beforehand.

(i.e., [0,1,2,4,5,6,7] might become [4,5,6,7,0,1,2]).

You are given a target value to search. If found in the array return its index, otherwise return -1.

You may assume no duplicate exists in the array.

Your algorithm's runtime complexity must be in the order of O(log n).

Example 1:

Input: nums = [4,5,6,7,0,1,2], target = 0
Output: 4
Example 2:

Input: nums = [4,5,6,7,0,1,2], target = 3
Output: -1

应该用二分法做的，不过暴力查其实也很快

*/
func search(nums []int, target int) int {

	for i,v := range nums {
		if v == target {
			return i
		}
	}
	return -1
}

