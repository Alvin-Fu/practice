package main

import "fmt"

// Topk问题，即给定一个数组获取最大的k个数

func main() {
	k := 5
	temp := []int{2, 3, 4, 5, 1, 10, 2, 3, 8, 5}
	topk1(temp, 0, len(temp)-1, k)
	fmt.Println(temp[:5])
}

// 使用简单随机的办法，分而治之
func topk1(arr []int, low, high, k int) int {
	if low >= high {
		return low
	}
	i := partition(arr, low, high)
	if i >= k {
		return topk1(arr, low, i-1, k)
	} else {
		return topk1(arr, i+1, high, k-i)
	}
}
func partition(nums []int, low, high int) int {
	mid := nums[low]
	for low < high {
		for low < high && mid >= nums[high] {
			high--
		}
		nums[low] = nums[high]
		for low < high && mid <= nums[low] {
			low++
		}
		nums[high] = nums[low]
	}
	nums[low] = mid
	return low
}
