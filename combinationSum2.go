package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []int{2,5,2,1,2}
	t := 5
	combinationSum2(s, t)
	fmt.Println(combinationSum2(s, t))
}

func combinationSum2(candidates []int, target int) [][]int {
	allResult := make([][]int, 0)
	sort.Ints(candidates)
	allResult = backtrack2(allResult, make([]int, 0), candidates, target)
	return allResult
}

func backtrack2(list [][]int, temp []int, nums []int, target int) [][]int {
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
				if i != 0 && nums[i-1] == nums[i] {
					continue
				}
				t := make([]int, len(temp))
				copy(t, temp)
				t = append(t, nums[i])
				list = backtrack2(list, t, nums[i+1:], target - nums[i])
			}
		}
	}
	return list
}


