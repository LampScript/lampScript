package main

import (
	"fmt"
)

func main() {
	fmt.Println(lengthOfLongestSubstringJava("abcabc"))
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