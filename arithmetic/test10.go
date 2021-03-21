package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		text := scanner.Text()
		tmp := strings.Split(text)
		nums := make([]int, len(tmp)-2)
		for i := 0; i < len(nums); i++ {
			n, _ := strconv.Atoi(tmp[i])
			nums[i] = n
		}
		lTime, _ := strconv.Atoi(tmp[len(tmp)-1])
		findMin(lTime, nums)
	}
}

// 3, 6, 11, 7 8
func findMin(lTime int, nums []int) int {
	if len(nums) > lTime {
		return -1
	}
	sort.Ints(nums)
	if len(nums) == lTime {
		return nums[len(nums)-1]
	}

}
