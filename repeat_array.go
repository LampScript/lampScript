package main

import "fmt"

func main()  {
	fmt.Println(removeDuplicates([]int{1,1,2,3,6}))
}



func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	curr := 0
	for i:=0;i<len(nums);i++ {
		if nums[curr] == nums[i] && curr == i {
			curr++
			continue
		} else if nums[curr] == nums[i] && curr != i {
			continue
		} else if nums[curr] != nums[i] && curr == i {
			curr++
			continue
		} else if nums[curr] != nums[i] && curr != i {
			curr++
			continue
		}
	}
	return curr
}