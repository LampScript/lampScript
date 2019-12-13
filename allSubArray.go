package main

import "fmt"

func main() {
	fmt.Println(getSubArray([]int{1,2,3,4}))
}

func getSubArray(num []int) [][]int  {
	all := make([][]int, 0)
	for i:=1;i<len(num);i++{
		tmpAll := make([][]int, 0)
		tmpAll, _ = getAllSubArray(tmpAll, num, make([]int, 0), i)
		all = append(all, tmpAll...)
	}
	return all
}

func getAllSubArray(all [][]int, num, current []int, length int) ([][]int, []int) {
	if length <= 1 {
		for _,v := range num {
			tmp := make([]int, len(current), len(current)+1)
			copy(tmp, current)
			tmp = append(tmp, v)
			all = append(all, tmp)
		}
		return all, current
	}

	l := len(num)
	for i:=0;i< l;i++ {
		current = append(current, num[i])
		num[0],num[i] = num[i], num[0]
		all, current = getAllSubArray(all, num[i+1:], current, length-1)
		current = current[:len(current)-1]
		num[0],num[i] = num[i], num[0]
	}
	return all, current
}