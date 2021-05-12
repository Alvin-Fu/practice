package main

import "fmt"

func main() {
	n := findMedianSortedArrays([]int{1, 3}, []int{2})
	fmt.Println(n)
}

func findMedianSortedArrays(nums1, nums2 []int) float64 {
	array := make([]int, 0)
	var n float64
	if len(nums1) >= len(nums2) {
		array = JoinArray(nums1, nums2)
	} else {
		array = JoinArray(nums2, nums1)
	}
	if len(array) == 0 {
		return 0
	}
	if len(array) == 1 {
		return float64(array[0])
	}
	if len(array)%2 == 0 {
		n := len(array) / 2
		return float64(array[n-1]+array[n]) / 2
	} else {
		n := len(array) / 2
		return float64(array[n])
	}
	return n
}

func JoinArray(nums1, nums2 []int) []int {
	if len(nums1) == 0 {
		return nums2
	}
	if len(nums2) == 0 {
		return nums1
	}
	array := make([]int, 0)
	index := 0
	for i := 0; i < len(nums1); {
		if index >= len(nums2) {
			array = append(array, nums1[i:]...)
			break
		}
		if nums1[i] < nums2[index] {
			array = append(array, nums1[i])
			i++
		} else {
			array = append(array, nums2[index])
			index++
		}
	}
	if index < len(nums2) {
		array = append(array, nums2[index:]...)
	}
	return array
}
