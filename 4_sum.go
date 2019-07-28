package main

import (
	"fmt"
	"sort"
)



/*

Runtime: 12 ms, faster than 61.33% of Go online submissions for 4Sum.
Memory Usage: 3.5 MB, less than 23.68% of Go online submissions for 4Sum.

https://leetcode.com/problems/4sum/submissions/

这题做得还行
*/
func main()  {
	fmt.Println(fourSum([]int{-4,-1,-1,0,1,2}, -1))
}


var dupMap map[string]bool

func fourSum(nums []int, target int) [][]int {
	dupMap = make(map[string]bool)
	sort.Ints(nums)
	result := make([][]int, 0)
	for i:=0;i < len(nums) - 3 ; i ++ {
		for j:=len(nums)-1;j > i + 2; j--  {
			tempTarget := target - (nums[i] + nums[j])
			num := nums[i+1: j]
			r := twoSum(num, tempTarget, nums[i], nums[j])
			if r != nil {
				result = append(result, r...)
			}
		}
	}
	return result
}




func twoSum(num []int, target, oldHead, oldTail int) [][]int {
	if len(num) < 2 {
		return nil
	}
	result := make([][]int, 0)
	head := 0
	tail := len(num) - 1
	for head < tail  {
		if head !=0 && num[head] == num[head - 1] {
			head++
			continue
		}
		temp := num[head] + num[tail] - target
		if temp == 0 {
			a := []int{oldHead, num[head], num[tail], oldTail}
			_, ok := dupMap[fmt.Sprintf("%v", a)];if !ok {
				dupMap[fmt.Sprintf("%v", a)] = true
				result = append(result, a)
			}
			head++
			tail--
		} else if temp > 0 {
			tail --
		} else if temp < 0 {
			head++
		}
	}
	return result
}