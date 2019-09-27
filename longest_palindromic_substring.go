package main

import (
	"fmt"
)


/*
示例 1：

输入: "babad"
输出: "bab"
注意: "aba" 也是一个有效答案。
示例 2：

输入: "cbbd"
输出: "bb"

*/


func main() {

	fmt.Println(longestPalindrome("bb"))
}

func longestPalindrome(s string) string {

	max := ""
	for i:=0;i<len(s);i++ {
		len1 := extendString(s, i, i)
		len2 := extendString(s, i, i+1)
		if len(len1) > len(max) {
			max = len1
		}
		if len(len2) > len(max) {
			max = len2
		}
	}
	return max
}


func extendString(s string, left, right int) string {
	temp := ""
	for left >=0 && right <len(s) && s[left] == s[right] {
		temp = s[left:right+1]
		left--
		right++
	}
	return temp
}





