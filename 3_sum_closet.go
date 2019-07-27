package main

import (
	"fmt"
	"sort"
)

/*
Given an array nums of n integers and an integer target, find three integers in nums such that the sum is closest to target.
Return the sum of the three integers. You may assume that each input would have exactly one solution.

Example:

Given array nums = [-1, 2, 1, -4], and target = 1.

The sum that is closest to the target is 2. (-1 + 2 + 1 = 2).
*/


/*

Runtime: 4 ms, faster than 98.91% of Go online submissions for 3Sum Closest.
Memory Usage: 2.7 MB, less than 83.78% of Go online submissions for 3Sum Closest.

*/

//这题做的还挺不错的
//https://leetcode.com/problems/3sum-closest/submissions/

func main()  {
	fmt.Println(threeSumClosest([]int{1,2,4,8,16,32,64,128}, 82))
}
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	all := nums[0] + nums[1] + nums[2]
	offset := Abs(nums[0] + nums[1] + nums[2] - target)
	for i:=0;i < len(nums) - 2; i ++ {
		if i==0 || nums[i] != nums[i - 1] {
			tail := len(nums) - 1
			cur := i + 1
			for cur < tail {
				sum := nums[i] + nums[tail] + nums[cur] - target
				if sum < 0 && Abs(sum) < offset {
					offset = Abs(sum)
					all = nums[i] + nums[tail] + nums[cur]
					cur++
				} else if sum == 0 {
					all = nums[i] + nums[tail] + nums[cur]
					return all
				} else if sum < 0 && Abs(sum) >= offset {
					cur++
				} else if sum > 0 && Abs(sum) < offset {
					offset =  Abs(sum)
					all = nums[i] + nums[tail] + nums[cur]
					tail--
				} else if sum > 0 && Abs(sum) >= offset {
					tail--
				}
			}
		}
	}
	return all
}

func Abs(a int) int  {
	if a < 0 {
		return -a
	}
	return a
}