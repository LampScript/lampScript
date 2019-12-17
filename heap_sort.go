package main

import (
	"fmt"
)

//
//func HeapSort(array []int) {
//	m := len(array)
//	s := m/2
//	for i := s; i > -1; i-- {
//		heap(array, i, m-1)
//	}
//	fmt.Println(array)
//	for i := m-1; i > 0; i-- {
//		array[i], array[0] = array[0], array[i]
//		heap(array, 0, i-1)
//	}
//}
//
//func heap(array []int, i, end int){
//	l := 2*i+1
//	if l > end {
//		return
//	}
//	n := l
//	r := 2*i+2
//	if r <= end && array[r]>array[l]{
//		n = r
//	}
//	if array[i] > array[n]{
//		return
//	}
//	array[n], array[i] = array[i], array[n]
//	heap(array, n, end)
//}

func main() {
	array := []int{
		-1, 2, 8, 4, 5, 9, 1, 7, 66,
	}
	HeapSort(array)
	fmt.Println(array)
}

func HeapSort(nums []int) {
	if len(nums) <= 1 {
		return
	}
	n := len(nums) / 2 - 1

	maxIndex := len(nums) - 1

	for n > -1 {
		left := (n + 1) * 2 - 1
		if left > maxIndex {
			n--
			continue
		}
		if nums[n] > nums[left] {
			nums[n], nums[left] = nums[left], nums[n]
		}

		right := (n + 1) * 2
		if right > maxIndex {
			n--
			continue
		}

		if nums[n] > nums[right] {
			nums[n], nums[right] = nums[right], nums[n]
		}
		n--
	}
	HeapSort(nums[1:])
}



