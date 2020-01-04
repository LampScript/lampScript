package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []int{7,3,2}
	t := 18
	combinationSum(s, t)
	fmt.Println(combinationSum(s, t))
}

func combinationSum(candidates []int, target int) [][]int {
	allResult := make([][]int, 0)
	sort.Ints(candidates)
	allResult = backtrack(allResult, make([]int, 0), candidates, target)
	return allResult
}

func backtrack(list [][]int, temp []int, nums []int, target int) [][]int {
	if target == 0 {
		list = append(list, temp)
	} else if target > 0 && len(nums) > 0 {
		for i:= 0; i< len(nums) ;i++  {
			if nums[i] == target {
				temp = append(temp, nums[i])
				list = append(list, temp)
				return list
			} else if nums[i]  > target {
				return list
			} else {
				t := make([]int, len(temp))
				copy(t, temp)
				t = append(t, nums[i])
				list = backtrack(list, t, nums[i:], target - nums[i])
			}
		}
	}
	return list
}


