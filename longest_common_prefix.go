package main

import "fmt"

func main()  {


	fmt.Println(longestCommonPrefix([]string{"a"}))
}

//https://leetcode.com/problems/longest-common-prefix/
func longestCommonPrefix(strs []string) string {
	if len(strs) < 1 {
		return ""
	}

	first := strs[0]
	common := ""
	for i:=0;i < len(first) + 1 ; i++ {
		temp := first[:i]

		for j:=0; j < len(strs) ; j++ {
			if len(strs[j]) < i {
				return common
			}
			temp2 := strs[j][:i]
			if temp2 != temp {
				return common
			}
		}
		common = temp
	}

	return common
}