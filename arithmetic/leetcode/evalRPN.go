package main

import (
	"strconv"
)

func main() {
	evalRPN([]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"})
}

func evalRPN(tokens []string) int {
	nums := make([]int, 0)
	operators := make([]string, 0)
	for _, t := range tokens {
		if t == "/" || t == "-" || t == "+" || t == "*" {
			if len(nums) < 2 {
				operators = append(operators, t)
			} else {
				l := len(nums)
				num := calculate(nums[l-2], nums[l-1], t)
				nums = nums[:l-2]
				nums = append(nums, num)
			}
		} else {
			n, _ := strconv.Atoi(t)
			nums = append(nums, n)
		}
	}
	return nums[0]
}

func calculate(a, b int, operate string) int {
	switch operate {
	case "/":
		return a / b
	case "*":
		return a * b
	case "+":
		return a + b
	case "-":
		return a - b
	default:
		return 0
	}
}
