package main

import "fmt"

func main() {

	array := make([]int, 3, 10)
	array[0] = 1
	array[1] = 2
	array[2] = 3
	a1 := array[:2]
	fmt.Println(a1)
	a1 = append(a1, 4,5,6)
	a1[0]= 123123
	fmt.Println(a1)
	fmt.Println(array)
}


func MergeSort(array []int) []int{
	n := len(array)
	if n < 2 {
		return array
	}
	key := n / 2
	left := MergeSort(array[0:key])
	right := MergeSort(array[key:])
	return merge(left, right)
}

func merge(left []int, right []int) []int{
	newArr := make([]int, len(left)+len(right))
	i, j, index :=0,0,0
	for {
		if left[i] > right[j] {
			newArr[index] = right[j]
			index++
			j++
			if j == len(right) {
				copy(newArr[index:], left[i:])
				break
			}

		}else{
			newArr[index] = left[i]
			index++
			i++
			if i == len(left) {
				copy(newArr[index:], right[j:])
				break
			}
		}
	}
	return newArr
}
