package main

import "fmt"

func main() {

}

func optimalDivision(nums []int) string {
	if len(nums) == 1 {
		return fmt.Sprintf("%d", nums[0])
	}
	if len(nums) == 2 {
		return fmt.Sprintf("%d/%d", nums[0], nums[1])
	}
	rue := fmt.Sprintf("%d/(%d", nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		rue += fmt.Sprintf("/%d", nums[i])
	}
	rue += ")"
	return rue
}
