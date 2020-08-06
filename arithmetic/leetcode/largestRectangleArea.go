package main

import "fmt"

func main() {
	fmt.Print(largestRectangleArea([]int{2, 1, 5, 6, 2, 3}))
}

func largestRectangleArea(heights []int) int {
	if len(heights) <= 0 {
		return 0
	}
	max := 0
	stack := make([]int, 0)
	for i := 0; i <= len(heights); i++ {
		cur := 0
		if i == len(heights) {
			cur = 0
		} else {
			cur = heights[i]
		}
		for len(stack) > 0 && cur <= heights[stack[len(stack)-1]] {
			pop := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			h := heights[pop]
			w := i
			if len(stack) != 0 {
				peek := stack[len(stack)-1]
				w = i - peek - 1
			}
			if max < h*w {
				max = h * w
			}
		}
		stack = append(stack, i)
	}
	return max
}
