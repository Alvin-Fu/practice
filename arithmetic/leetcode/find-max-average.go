package main

import "fmt"

func main() {
	nums := []int{1, 12, -5, -6, 50, 3}
	fmt.Println(findMaxAverage2(nums, 4))
}

func findMaxAverage(nums []int, k int) float64 {
	if len(nums) < k {
		return 0
	}
	max := 0
	for i := 0; i < k; i++ {
		max += nums[i]
	}
	index := 1
	for len(nums)-index >= k {
		tmp := 0
		for i := 0; i+index < len(nums) && len(nums)-index >= k && i < k; i++ {
			tmp += nums[i+index]
		}
		if tmp > max {
			max = tmp
		}
		index++
	}
	return float64(max) / float64(k)
}

func findMaxAverage2(nums []int, k int) float64 {
	if len(nums) < k {
		return 0
	}
	max := 0
	for i := 0; i < k; i++ {
		max += nums[i]
	}
	sum := max
	for i := k; i < len(nums); i++ {
		sum = sum - nums[i-k] + nums[i]
		if sum > max {
			max = sum
		}
	}
	return float64(max) / float64(k)
}
