package main

import "fmt"

func main() {
	tmp := []int{2, 2, 2, 2, 3, 3, 3, 3, 8}
	fmt.Println(singleNumber(tmp))
}

func singleNumber(nums []int) int {
	var result int
	for i := 0; i < 64; i++ {
		sum := 0
		for j := 0; j < len(nums); j++ {
			// 统计1的个数
			sum += (nums[j] >> i) & 1
		}
		// 还原位00^10=10 或者用| 也可以
		result ^= (sum % 4) << i
	}
	return result
}
