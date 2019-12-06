package main

import "fmt"

func main()  {

	fmt.Println(maxSubArray([]int{-2,1,-3,4,-1,2,1,-5,4}))

}

//todo 结果不对，还有边界条件

func maxSubArray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	} else if len(nums) == 2 {
		if nums[0] > nums[1] && nums[0] > nums[0] + nums[1] {
			return nums[0]
		} else if nums[1] > nums[0] && nums[1] > nums[0] + nums[1] {
			return nums[1]
		} else {
			return nums[0]+nums[1]
		}
	} else if len(nums) == 0 {
		fmt.Println("len(nums) == 0")
		return 0
	}

	left := maxSubArray(nums[:len(nums)/2])
	right := maxSubArray(nums[len(nums)/2:])

	mid := getMid(nums)

	if left >= right && left >= mid {
		return left
	}
	if left <= right && right >= mid {
		return right
	}
	if left <= mid && right <= mid {
		return mid
	}
	return 0
}

func getMid(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	maxLeft := nums[len(nums)/2-1]
	for i:=len(nums)/2-2;i>=0 ;i-- {
		tmp := maxLeft + nums[i]
		if tmp > maxLeft {
			maxLeft = tmp
		}
	}
	maxRight := nums[len(nums)/2]
	for i:=len(nums)/2 + 1;i<len(nums) ;i++ {
		tmp := maxRight + nums[i]
		if tmp > maxRight {
			maxRight = tmp
		}
	}
	return maxLeft + maxRight
}