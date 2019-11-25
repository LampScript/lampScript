package main

import (
	"fmt"
)

func main()  {
	fmt.Println(twoSumHash([]int{3,5,0,4}, 7))
}

func twoSumHash(nums []int, target int) []int {
	if nums == nil || len(nums) == 0 {
		return nil
	}

	m := make(map[int]int, len(nums))

	for i,v := range nums {
		left := target - v
		i2, ok := m[left]; if ok {
			return []int{i, i2}
		}
		m[v]=i
	}
	return nil
}

func twoSumBrute(nums []int, target int) []int {

	if nums == nil || len(nums) == 0 {
		return nums
	}

	for i, v := range nums {
		left := target - v
		if i == len(nums)-1 {
			continue
		}

		for i2, j := range nums[i+1:] {
			if j == left {
				return []int{i,i2+i+1}
			}
		}
	}
	return nil
}

func twoSum2(nums []int, target int) []int {

	if nums == nil {
		return nil
	}
	for i:=0;i<len(nums);i++ {
		temp:= target - nums[i]
		if target <= 0 {
			continue
		}
		for j:=i+1;j<len(nums);j++{
			if nums[j] - temp == 0 {
				return []int{i, j}
			}
		}
	}

	return nil
}

func twoSum3(nums []int, target int) []int {

	if nums == nil {
		return nil
	}

	sumMap := make(map[int]int)
	for i:=0;i<len(nums);i++ {
		temp:= target - nums[i]
		if index, ok := sumMap[temp]; ok {
			return []int{i, index}
		} else {
			sumMap[nums[i]] = i
		}
	}

	return nil
}