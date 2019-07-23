package main

import (
	"fmt"
	"strings"
)

/*
The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this: (you may want to display this pattern in a fixed font for better legibility)

P   A   H   N
A P L S I I G
Y   I   R
And then read line by line: "PAHNAPLSIIGYIR"

Write the code that will take a string and make this conversion given a number of rows:

string convert(string s, int numRows);
Example 1:

Input: s = "PAYPALISHIRING", numRows = 3
Output: "PAHNAPLSIIGYIR"
Example 2:

Input: s = "PAYPALISHIRING", numRows = 4
Output: "PINALSIGYAHRPI"
Explanation:

P     I    N
A   L S  I G
Y A   H R
P     I

*/

//https://leetcode.com/problems/zigzag-conversion/solution/
func main()  {
	fmt.Println(convert("PAYPALISHIRING", 3))

}


func convert(s string, numRows int) string {
	square := make([][]string, 0, 0)
	array := strings.Split(s, "")
	if len(array) < 2 {
		return s
	}
	for i:=0; i< len(array); i++ {
		c, r := getLocation(i, numRows)
		if len(square) < c + 1 {
			square = append(square, []string{})
		}

		list := square[c]
		if list == nil || len(list) == 0 {
			list = make([]string, numRows)
		}
		list[r] = array[i]
		square[c] = list
	}

	result := ""
	for i:=0;i< numRows; i++  {
		for j:=0; j < len(square); j++  {
			temp := square[j]
			result = result + temp[i]
		}
	}

	return result
}

func getLocation(index, numRows int) (col, row int) {
	if numRows < 2 {
		return index, 0
	}
	offset := index / (2 * (numRows - 1)) * (numRows - 1)
	index = index % (2 * (numRows - 1))
	if index < numRows {
		row = index
		col = 0
	} else {
		row = numRows - 1 - (index - numRows + 1)
		col = 0 + (index - numRows + 1)
	}
	col = col + offset
	return
}