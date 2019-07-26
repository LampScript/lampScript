package main

import (
	"fmt"
	"strings"
)

/*

Roman numerals are represented by seven different symbols: I, V, X, L, C, D and M.
Symbol       Value
I             1
V             5
X             10
L             50
C             100
D             500
M             1000
For example, two is written as II in Roman numeral, just two one's added together.
Twelve is written as, XII, which is simply X + II. The number twenty seven is written as XXVII, which is XX + V + II.

Roman numerals are usually written largest to smallest from left to right.
However, the numeral for four is not IIII. Instead, the number four is written as IV.
Because the one is before the five we subtract it making four.
The same principle applies to the number nine, which is written as IX.
There are six instances where subtraction is used:

I can be placed before V (5) and X (10) to make 4 and 9.
X can be placed before L (50) and C (100) to make 40 and 90.
C can be placed before D (500) and M (1000) to make 400 and 900.
Given an integer, convert it to a roman numeral. Input is guaranteed to be within the range from 1 to 3999.
*/

var levelMap map[int]string
var subMap map[int]int
var nextMap map[int]int

func main()  {
	levelMap = make(map[int]string, 7)
	levelMap[1000] = "M"
	levelMap[500] = "D"
	levelMap[100] = "C"
	levelMap[50] = "L"
	levelMap[10] = "X"
	levelMap[5] = "V"
	levelMap[1] = "I"

	subMap = make(map[int]int, 4)
	subMap[1000] = 100
	subMap[500] = 100
	subMap[100] = 10
	subMap[50] = 10
	subMap[10] = 1
	subMap[5] = 1

	nextMap = make(map[int]int, 4)
	nextMap[1000] = 500
	nextMap[500] = 100
	nextMap[100] = 50
	nextMap[50] = 10
	nextMap[10] = 5
	nextMap[5] = 1

	fmt.Println(intToRoman(1))
}

//这个答案太秀了
//https://leetcode.com/problems/integer-to-roman/discuss/?currentPage=1&orderBy=most_votes&query=
func intToRoman(num int) string {
	level := 1000
	s := ""
	for num > 0 {
		temp, left := isM(num, level)
		if left < level - subMap[level] {
			level = nextMap[level]
		}
		s = s + temp
		num = left
	}

	return s
}

func isM(num, level int) (string, int)  {
	if num >= level {
		s := strings.Repeat(levelMap[level], num/ level)
		return s, num - (num / level) * level
	}


	next, ok := subMap[level];if ok {
		if num >= level - next {
			return levelMap[next] + levelMap[level], num + next - level
		}
	}

	return "", num
}

