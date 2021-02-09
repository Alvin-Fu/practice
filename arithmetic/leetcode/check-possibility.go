package main

import (
	"fmt"
)

/*
给你一个长度为 n 的整数数组，请你判断在 最多 改变 1 个元素的情况下，该数组能否变成一个非递减数列。

我们是这样定义一个非递减数列的： 对于数组中所有的 i (0 <= i <= n-2)，总满足 nums[i] <= nums[i + 1]。

输入: nums = [4,2,3]
输出: true
解释: 你可以通过把第一个4变成1来使得它成为一个非递减数列。
*/

func main() {
	nums := []int{3, 4, 2, 3}
	fmt.Println(checkPossibility2(nums))
}

func checkPossibility(nums []int) bool {
	if len(nums) <= 1 {
		return false
	}
	index := 0
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			tmp := nums[i+1]
			nums[i+1] = nums[i]
			flag := false
			for j := i + 1; j < len(nums)-1; j++ {
				if nums[j] > nums[j+1] {
					flag = true
					break
				}
			}
			fmt.Println(i, flag)
			if !flag {
				return true
			}
			nums[i+1] = tmp
			nums[i] = tmp
			for j := i; j >= 1; j-- {
				if nums[j] < nums[j-1] {
					return false
				}
			}
			for j := i + 1; j < len(nums)-1; j++ {
				if nums[j] > nums[j+1] {
					return false
				}
			}
			index++
		}
		if index >= 2 {
			return false
		}
	}
	return true
}

func checkPossibility2(nums []int) bool {
	if len(nums) <= 1 {
		return false
	}
	index := 0
	for i := 0; i < len(nums)-1; i++ {
		x, y := nums[i], nums[i+1]
		if x > y {
			index++
			if i > 0 && y < nums[i-1] {
				nums[i+1] = x
			}
			if index > 1 {
				return false
			}
		}
	}
	return true
}
