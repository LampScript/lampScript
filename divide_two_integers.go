package main

import (
	"fmt"
	"math"
)

func main()  {
	fmt.Println(divide(-2147483648, -1))
}


/*
86题
https://leetcode.com/problems/divide-two-integers/submissions/
Runtime: 2488 ms, faster than 19.01% of Go online submissions for Divide Two Integers.
Memory Usage: 2.4 MB, less than 45.45% of Go online submissions for Divide Two Integers.

其实不知道这题考点是啥。。。
*/
func divide(dividend int, divisor int) int {

	if dividend == 0  {
		return 0
	}

	positive := false
	if (dividend > 0 && divisor > 0) || (divisor < 0 && dividend < 0 ) {
		positive = true
	}

	if dividend < 0 {
		dividend = -dividend
	}
	if divisor < 0 {
		divisor = -divisor
	}

	i := 0
	for true {
		if dividend < divisor {
			break
		}
		dividend = dividend - divisor
		i++
	}

	if !positive {
		i = -i
		if i < -math.MaxInt32 - 1 {
			i = -math.MaxInt32 - 1
		}
		return i
	}
	
	if i > math.MaxInt32 {
		i = math.MaxInt32
	}

	return i
}