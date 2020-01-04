package main

import "fmt"

func main() {
	arr := []int{1, 2, 3}
	for i:=0;i<len(arr) ; i++ {
		fmt.Println(len(arr))
		arr = append(arr, len(arr))
	}
	fmt.Println(arr)
}




func multiply(num1, num2 string) string {


	return ""
}