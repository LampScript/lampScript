package main

import "fmt"

func main() {

	b := []int{1,2,3}
	t := append(b[:1], b[2:]...)
	getSea(t)
	fmt.Println(b)
	fmt.Println(t)
}

func getSea(n []int)  {
	fmt.Println(n)
}