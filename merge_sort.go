package main

import "fmt"

func main() {

	array := []int{55, 94, 87, 1, 4, 32, 11, 77, 39, 42, 64, 53, 70, 12, 9}
	fmt.Println(array)
	array = MergeSort(array)
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
