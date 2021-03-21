package main

import "fmt"

func main() {
	num := 0
	fmt.Scan(&num)
	nums := make([]int, 0)
	for num > 0 {
		if num%26 != 0 {
			nums = append(nums, num%26)
			num /= 26
		} else {
			nums = append(nums, 26)
			num = num/26 - 1
		}
	}
	rue := make([]byte, 0)
	for i := len(nums) - 1; i >= 0; i-- {
		a := uint8(nums[i]) + 'a' - 1
		rue = append(rue, a)
	}
	fmt.Println(string(rue))
}
