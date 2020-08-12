package main

import "fmt"

func main() {
	nums := []int{2, 4, 1, 8, 5, 3, 0, 9, 7, 10, 2, 3}
	//result := mergeSortR(nums)
	//fmt.Println(result)
	result := mergeSortC(nums)
	fmt.Println(result)
}

// 递归版本
func mergeSortR(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	mid := len(nums) / 2
	// 先进行分
	left := mergeSortR(nums[:mid])
	right := mergeSortR(nums[mid:])
	result := mergeR(left, right)
	return result
}

func mergeR(left, right []int) []int {
	// 治理
	result := make([]int, 0)
	l, r := 0, 0
	for l < len(left) && r < len(right) {
		if left[l] < right[r] {
			result = append(result, left[l])
			l++
		} else {
			result = append(result, right[r])
			r++
		}
	}
	result = append(result, left[l:]...)
	result = append(result, right[r:]...)
	return result
}

// 使用循环
func mergeSortC(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	step := 2
	length := len(nums)
    i := 0
	for step <= length {
		i = 0
		temp := make([]int, 0)
		for i+step <= length {
			temp = append(temp, mergeR(nums[i:i + step/2], nums[i + step/2:i + step])...)
			i += step
		}
		if i + step / 2 <= length{
            temp = append(temp, mergeR(nums[i:i + step / 2], nums[i + step/2:length ])...)
        }
		tmp := nums[len(temp):]
		nums = temp
		nums = append(nums, tmp...)
		step = step * 2
	}
	return mergeR(nums[0:step/2], nums[step/2:])
}

