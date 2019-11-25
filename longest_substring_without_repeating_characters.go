package main

import (
	"fmt"
)

func main() {
	fmt.Println(lengthOfLongestSubstringSlideWindow("dvdf"))

}


func lengthOfLongestSubstringSlideWindow(s string) int {
	if len(s) < 2 {
		return 1
	}
	max := 0
	left := 0
	right := 1
	for ;right<len(s);right++{
		v := string(s[right])
		tmp := Unique(s[left:right], v)
		if tmp >= 0 {
			left += tmp+1
		}
		if len(s[left:right+1]) > max {
			max = right - left + 1
		}
	}
	return max
}

func Unique(r1 string, r string) int  {
	for i, v := range r1{
		if string(v) == r {
			return i
		}
	}
	return -1
}


func lengthOfLongestSubstring(s string) int {

	longest := ""
	cursor := ""
	cursorMap := make(map[string]int)
	for i:=0;i<len(s); {
		v := string(s[i:i+1])
		if index, ok := cursorMap[v]; !ok {
			cursor += v
			if len(cursor) > len(longest) {
				longest = cursor
			}
			cursorMap[v]=i
			i++
		} else {
			i = index+1
			cursor = ""
			cursorMap = make(map[string]int)
		}
	}

	fmt.Println(longest)
	return len(longest)
}

func lengthOfLongestSubstringJava(s string) int {
	max := 0
	cursorMap := make(map[string]int)
	for i, j := 0, 0;i<len(s);i++ {
		v := string(s[i:i+1])
		if index, ok := cursorMap[v]; ok {
			if index + 1 > j {
				j = index + 1
			}
		}
		cursorMap[v] = i
		if i-j+1 > max {
			max = i-j+1
		}
	}

	return max
}