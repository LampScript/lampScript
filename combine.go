package main

import (
	"fmt"
)

var allCombine [][]int

func main()  {
	fmt.Println(permute([]int{1,2,3}))
}

func getAllNum(input []int)  {

	usedMap := make(map[int]struct{})

	for _,v := range input  {
		usedMap[v] = struct{}{}
	}

	for i:=1;i<len(input)+1 ;i++  {
		combine(usedMap, make([]int, i), i)
	}
}

func combine(data map[int]struct{}, result []int, target int) {
	if target == 1 {
		for i := range data  {
			temp := make([]int, len(result))
			copy(temp, result)
			temp[len(temp)-1] = i
			allCombine = append(allCombine, temp)
		}
		return
	}
	for i := range data {
		delete(data, i)
		result[len(result) - target] = i
		combine(data, result, target-1)
		data[i]= struct{}{}
	}
}

func permute(input []int) [][]int  {

	out := make([][]int, 0)

	out = backTrack(len(input), 0, out, input)
	return out
}

func backTrack(n, first int, out [][]int, nums []int) [][]int  {
	if n == first {
		s := make([]int, len(nums))
		copy(s, nums)
		out = append(out, s)
	}

	for i:=first;i<len(nums);i++  {
		nums[first], nums[i] = nums[i], nums[first]
		out = backTrack(n, first+1, out, nums)
		nums[first], nums[i] = nums[i], nums[first]
	}
	return out
}


