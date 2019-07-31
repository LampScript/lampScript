package main

import "fmt"

func main()  {
	fmt.Println(strStr("a", "a"))
}

/*
这题好像没什么难的，实际上考的是 KMP 算法
https://leetcode.com/problems/implement-strstr/submissions/
Runtime: 0 ms, faster than 100.00% of Go online submissions for Implement strStr().
Memory Usage: 2.3 MB, less than 98.51% of Go online submissions for Implement strStr().
*/
func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}

	if len(haystack) < len(needle) {
		return -1
	}

	for i:=0;i< len(haystack) - len(needle) + 1;i++  {
		if haystack[i:i+len(needle)] == needle {
			return i
		}
	}

	return -1
}