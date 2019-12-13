package main

import "fmt"

func main() {
	fmt.Println(canJump([]int{2,3,1,1,4}))
}

func canJump(nums []int) bool {
	if len(nums) < 1 {
		return false
	}
	memo := make([]int, len(nums))

	return tryJump(nums, 0, memo)
}

func tryJump(nums []int, pos int, memo []int) bool  {
	if len(nums) == pos + 1 {
		memo[pos] = 1
		return true
	}

	if memo[pos] != 0 {
		return memo[pos] == 1
	}

	furthestJump := pos + nums[pos]
	if furthestJump > len(nums) - 1 {
		furthestJump = len(nums) - 1
	}

	for i:=pos + 1;i<=furthestJump;i++ {
		if tryJump(nums, i, memo) {
			memo[pos] = 1
			return true
		}
	}
	memo[pos] = 2
	return false
}