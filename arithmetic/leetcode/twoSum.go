package main



func TwoSum(nums []int, target int) []int {
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
