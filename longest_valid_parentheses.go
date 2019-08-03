package main

import (
	"fmt"
	"strings"
)


/*
https://leetcode.com/problems/longest-valid-parentheses/discuss/?currentPage=1&orderBy=most_votes&query=
太难了，自闭

*/
func main()  {
	fmt.Println(longestValidParentheses("()(())"))
}

/*
func longestValidParentheses(s string) int {

	length := len(s) / 2

	for true {
		temp := strings.Repeat("()", length)
		if strings.Contains(s, temp) {
			return len(temp)
		}
		length--
	}
	return 0
}*/



func longestValidParentheses(s string) int {
	l := 0
	r := 0
	i:=0
	for ;i<len(s) ;i++ {
		if strings.HasPrefix(s[i:], "(")  {
			l++
		}
		if strings.HasPrefix(s[i:], ")") && l > r {
			r++
		}
	}
	return r * 2
}
