package main

import (
	"fmt"
	"strconv"
)

func main() {
	rue := decodeString("2[ab4[d]]3[ef]g")
	fmt.Println(rue)
}

func decodeString(s string) string {
	if len(s) < 1 {
		return s
	}
	stack := make([]byte, 0)
	data := []byte(s)
	for i := 0; i < len(data); i++ {
		switch data[i] {
		case ']':
			temp := make([]byte, 0)
			for len(stack) != 0 && stack[len(stack)-1] != '[' {
				v := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				temp = append(temp, v)
			}
			stack = stack[:len(stack)-1]
			idx := 1
			for len(stack) >= idx && (stack[len(stack)-idx] >= '0' && stack[len(stack)-idx] <= '9') {
				idx++
			}
			num, _ := strconv.Atoi(string(stack[len(stack)-idx+1:]))
			stack = stack[:len(stack)-idx+1]
			for j := 0; j < num; j++ {
				for k := len(temp) - 1; k >= 0; k-- {
					stack = append(stack, temp[k])
				}
			}
		default:
			stack = append(stack, data[i])
		}
	}

	return string(stack)
}
