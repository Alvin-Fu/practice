package main

import (
	"fmt"
)

func main() {
	str := ""
	fmt.Scanf("%s", &str)
	fmt.Println(getLangStr(str))
}

func getLangStr(str string) int {
	max := 0
	stack := make([]int, 0)
	stack = append(stack, -1)
	for i := 0; i < len(str); i++ {
		if str[i] == '{' {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				max = maxInt(max, i-stack[len(stack)-1])
			}
		}
	}
	return max
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
