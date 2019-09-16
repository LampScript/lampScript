package main

import "fmt"

func main() {
	fmt.Println(findMedianSortedArrays([]int{1,2}, []int{3,4}))
}


func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums3 := make([]int, 0, len(nums1) + len(nums2))

	i:=0
	j:=0
	for true {
		if i >= len(nums1) && j >= len(nums2) {
			break
		}
		temp := 0
		hasValue := false
		if i < len(nums1) {
			temp = nums1[i]
			hasValue = true
		}
		if j < len(nums2) && hasValue && nums2[j] < temp {
			nums3 = append(nums3, nums2[j])
			j++
		} else if j < len(nums2) && !hasValue {
			nums3 = append(nums3, nums2[j])
			j++
		} else {
			nums3 = append(nums3, temp)
			i++
		}
	}
	if len(nums3) % 2 == 0 {
		a := nums3[len(nums3)/2] + nums3[len(nums3)/2 - 1]
		return float64(a)/2
	} else {
		return float64(nums3[len(nums3)/2])
	}
}