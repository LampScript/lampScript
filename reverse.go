package main

import (
	"fmt"
	"math"
)

func main() {
	//fmt.Println(math.MaxInt32 > reverse(1534236469))
	fmt.Println(reverse(-2147483648))
	//fmt.Println(reverse(-2147483412))
}

func reverse(x int) int {
	var r int
	for true{
		r = r + x % 10
		x/=10
		if x == 0 {
			break
		}
		if  math.MaxInt32/ 10 < r {
			return 0
		}
		r *= 10
	}
	return r
}