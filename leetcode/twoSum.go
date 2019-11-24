package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{-4,-2,-2,-2,0,1,2,2,2,3,3,4,4,6,6}, 4))
}

func twoSum(nums []int, target int) []int {
	for i, n := range nums {
		tmp := target - n
		newNums := nums[i+1:]
		//fmt.Println(newNums)
		for j, m := range newNums {
			if tmp == m {
				return []int{i, j + i+1}
			}
		}
	}
	return nil
}
