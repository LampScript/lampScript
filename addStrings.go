package main

import (
	"fmt"
	"strconv"
)

func main() {

	fmt.Println(addStrings("1234", "9"))
}



func addStrings(num1 string, num2 string) string {

	c := 0
	result := ""

	i:=len(num1)-1
	j:=len(num2)-1

	for {
		if i < 0 && j < 0  {
			break
		}
		var a1 int
		var a2 int
		if i >= 0 {
			a1, _ = strconv.Atoi(string(num1[i]))
			i--
		}

		if j >= 0 {
			a2, _ = strconv.Atoi(string(num2[j]))
			j--
		}

		tmp := a1 + a2 + c
		c=0
		result =  strconv.Itoa(tmp%10) + result
		if tmp > 9 {
			c = 1
		}
	}

	if c > 0 {
		result="1"+result
	}

	return result
}
