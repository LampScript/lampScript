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
For example, two is written as II in Roman numeral, just two one's added together. Twelve is written as, XII, which is simply X + II. The number twenty seven is written as XXVII, which is XX + V + II.

Roman numerals are usually written largest to smallest from left to right. However, the numeral for four is not IIII. Instead, the number four is written as IV. Because the one is before the five we subtract it making four. The same principle applies to the number nine, which is written as IX. There are six instances where subtraction is used:

I can be placed before V (5) and X (10) to make 4 and 9.
X can be placed before L (50) and C (100) to make 40 and 90.
C can be placed before D (500) and M (1000) to make 400 and 900.
Given a roman numeral, convert it to an integer. Input is guaranteed to be within the range from 1 to 3999.

Example 1:

Input: "III"
Output: 3
Example 2:

Input: "IV"
Output: 4
Example 3:

Input: "IX"
Output: 9
Example 4:

Input: "LVIII"
Output: 58
Explanation: L = 50, V= 5, III = 3.
Example 5:

Input: "MCMXCIV"
Output: 1994
Explanation: M = 1000, CM = 900, XC = 90 and IV = 4.

*/

func main()  {

	fmt.Println(romanToInt("MCMXCIV"))
}

//https://leetcode.com/problems/roman-to-integer/discuss/6529/My-solution-for-this-question-but-I-don't-know-is-there-any-easier-way

func romanToInt(s string) int {
	num := 0


	for true {
		f := strings.HasPrefix(s, "M")
		if f {
			num = num + 1000
			s = strings.Replace(s, "M", "", 1)
		} else {
			break
		}
	}

	for true {
		f := strings.HasPrefix(s, "CM")
		if f {
			num = num + 900
			s = strings.Replace(s, "CM", "", 1)
		} else {
			break
		}
	}

	for true {
		f := strings.HasPrefix(s, "D")
		if f {
			num = num + 500
			s = strings.Replace(s, "D", "", 1)
		} else {
			break
		}
	}

	for true {
		f := strings.HasPrefix(s, "CD")
		if f {
			num = num + 400
			s = strings.Replace(s, "CD", "", 1)
		} else {
			break
		}
	}
	for true {
		f := strings.HasPrefix(s, "C")
		if f {
			num = num + 100
			s = strings.Replace(s, "C", "", 1)
		} else {
			break
		}
	}

	for true {
		f := strings.HasPrefix(s, "XC")
		if f {
			num = num + 90
			s = strings.Replace(s, "XC", "", 1)
		} else {
			break
		}
	}

	for true {
		f := strings.HasPrefix(s, "L")
		if f {
			num = num + 50
			s = strings.Replace(s, "L", "", 1)
		} else {
			break
		}
	}

	for true {
		f := strings.HasPrefix(s, "XL")
		if f {
			num = num + 40
			s = strings.Replace(s, "XL", "", 1)
		} else {
			break
		}
	}

	for true {
		f := strings.HasPrefix(s, "X")
		if f {
			num = num + 10
			s = strings.Replace(s, "X", "", 1)
		} else {
			break
		}
	}
	for true {
		f := strings.HasPrefix(s, "IX")
		if f {
			num = num + 9
			s = strings.Replace(s, "IX", "", 1)
		} else {
			break
		}
	}

	for true {
		f := strings.HasPrefix(s, "V")
		if f {
			num = num + 5
			s = strings.Replace(s, "V", "", 1)
		} else {
			break
		}
	}

	for true {
		f := strings.HasPrefix(s, "IV")
		if f {
			num = num + 4
			s = strings.Replace(s, "IV", "", 1)
		} else {
			break
		}
	}

	for true {
		f := strings.HasPrefix(s, "I")
		if f {
			num = num + 1
			s = strings.Replace(s, "I", "", 1)
		} else {
			break
		}
	}



	return num
}
