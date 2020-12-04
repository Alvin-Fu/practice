package main

import (
	"fmt"
)

/*
给你一个按升序排序的整数数组 num（可能包含重复数字），请你将它们分割成一个或多个子序列，
其中每个子序列都由连续整数组成且长度至少为 3 。

如果可以完成上述分割，则返回 true ；否则，返回 false 。

示例 1：

输入: [1,2,3,3,4,5]
输出: True
解释:
你可以分割出这样两个连续子序列 :
1, 2, 3
3, 4, 5

*/

func main() {
	isPossible([]int{1, 2, 3, 3, 4, 5})
}

func isPossible(nums []int) bool {
	if len(nums) < 3 || len(nums) > 10000 {
		return false
	}
	temp := make(map[int]int, len(nums))
	tmp := make(map[int]int)
	for _, n := range nums {
		temp[n]++
	}
	for _, n := range nums {
		if temp[n] == 0 {
			continue
		}
		temp[n]--
		if tmp[n-1] > 0 {
			tmp[n-1]--
			tmp[n]++
		} else if temp[n+1] > 0 && temp[n+2] > 0 {
			//从当前的位置找前面的两个位置是否存在，要是存在则设置结尾元素
			temp[n+1]--
			temp[n+2]--
			tmp[n+2]++
		} else {
			return false
		}
	}
	fmt.Println(tmp)
	return true
}
