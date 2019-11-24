package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main(){
	fmt.Println(threeSum([]int{0,0,0,0,1}))
}
func threeSum(nums []int) [][]int {
	if len(nums) < 3{
		return nil
	}
	sort.Ints(nums)
	if nums[0] > 0 {
		return nil
	}
	var results = make([][]int,0)
	checkResult := make(map[string]bool)
	for i, n := range nums{
		if n > 0 {
			break
		}
		if i > 0 {
			if n == nums[i - 1]{
				continue
			}
		}
		indexs := TwoSum(nums, -n, i)
		if indexs != nil {
			for _, index := range indexs{
				key := getAscii([]int{nums[i], nums[index[0]], nums[index[1]]})
				fmt.Println(key)
				if _, ok := checkResult[key]; !ok {
					checkResult[key] = true
					results = append(results, []int{nums[i], nums[index[0]], nums[index[1]]})
				}
			}
		}
	}
	return results
}

func TwoSum(nums []int, target int, index int) [][]int {
	var numMap = make(map[int]int, len(nums))
	var results = make([][]int, 0)
	for i, n := range nums{
		if i == index {
			continue
		}
		if j, ok := numMap[n]; ok {
			results = append( results, []int{j, i})
		}
		if _, ok := numMap[target - n]; !ok {
			numMap[target - n] =  i
		}
	}
	return results
}

func getAscii(arr  []int)string{
	sort.Ints(arr)
	return strconv.Itoa(arr[0]) + strconv.Itoa(arr[1]) + strconv.Itoa(arr[2])
}