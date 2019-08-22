package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []int{1,2,3}
	t := 3
	combinationSum(s, t)
	fmt.Println(combinationSum(s, t))
}

/*func main() {
	nums:=[]int{1,2,3,4,5}
	fmt.Println("  len cap   address")
	fmt.Print("111---",len(nums),cap(nums))
	fmt.Printf("    %p\n",nums)//0xc4200181e0
	a:=nums[:3]
	fmt.Print("222---",len(a),cap(a))
	fmt.Printf("    %p\n",a)//0xc4200181e0 一样
	//output: 222--- 3 5

	b:= make([]int, 3) //第二个冒号 设置cap的
	n:=copy(b,nums[:3:3]) //第二个冒号 设置cap的
	fmt.Print("333---",len(b),cap(b))
	fmt.Printf("    %p\n",b)//0xc4200181e0 一样
	//output: 333--- 3 3
	fmt.Println(n,b)
	nums[0]=55
	fmt.Println(nums,a,b)
}*/

var allResult = make([][]int, 0)

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	backtrack(allResult, make([]int, 0), candidates, target, 0)
	return allResult
}

func backtrack(list [][]int, temp []int, nums []int, remain int, start int )  {
	fmt.Printf("list addr is %p \n", list)
	if remain < 0 {
		return
	} else if remain == 0 {
		r := make([]int, len(temp))
		copy(r, temp)
		list = append(list, r)
		//fmt.Println(list)
	} else {
		for i:= start; i< len(nums) ;i++  {
			temp = append(temp, nums[i])
			backtrack(list, temp, nums, remain - nums[i], i)
			temp = temp[len(temp)-1:]
		}
	}
}
