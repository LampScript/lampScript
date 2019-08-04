package main

import "fmt"



/*
Runtime: 8 ms, faster than 85.63% of Go online submissions for Find First and Last Position of Element in Sorted Array.
Memory Usage: 4.1 MB, less than 27.27% of Go online submissions for Find First and Last Position of Element in Sorted Array.

这题做得还行

https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/submissions/
*/
func main()  {

	//fmt.Println(binSearch([]int{1,3,5}, 5))
	//fmt.Println(binSearch([]int{5,7,7,8,8,10}, 6))
	fmt.Println(searchRange([]int{0,0,1,1,1,4,5,5}, 2))
}




func searchRange(nums []int, target int) []int {

	a := binSearch(nums, target)

	left := a
	right := a
	result := []int{left, right}
	if a == -1 {
		return result
	}

	for left > 0 {
		if nums[left-1] == target {
			left--
			result[0] = left
		} else {
			break
		}
	}

	for right < len(nums) - 1 {
		if nums[right+1] == target {
			right++
			result[1] = right
		} else {
			break
		}
	}
	return result
}



func binSearch(nums []int, target int) int {
	if len(nums) == 1 {
		if nums[0] == target {
			return 0
		} else {
			return -1
		}
	}
	lo := 0
	hi := len(nums) - 1

	for lo < hi {
		mid := (lo + hi) / 2
		if target == nums[mid] {
			return mid
		}

		if hi - lo == 1 && len(nums) - 1 >= mid + 1 {
			if target == nums[mid+1] {
				return mid + 1
			} else {
				return -1
			}
		}

		if target < nums[mid] {
			hi = mid
		} else if target > nums[mid] {
			lo = mid
		}
	}
	return -1
}