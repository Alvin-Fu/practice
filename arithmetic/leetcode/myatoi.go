package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(strconv.Atoi("012"))
	fmt.Println(2 << 32)
	fmt.Println(myAtoi("-004.2nnn77"))
}

func myAtoi(str string) int {
	data := []byte(str)
	tmp := make([]byte, 0)
	isNegative := false
	for _, d := range data {
		if d == ' ' {
			continue
		}
		if d == '-' && len(tmp) == 0 {
			tmp = append(tmp, d)
			isNegative = true
			continue
		}
		if d <= '0' && d >= '9' {
			if len(tmp) == 0 {
				return 0
			}
			continue
		}
		if d >= '0' && d <= '9' {
			if len(tmp) == 0 && d == '0' {
				continue
			}
			if len(tmp) == 1 && tmp[0] == '-' && d == '0' {
				continue
			}
			tmp = append(tmp, d)
		} else if len(tmp) != 0 {
			break
		}

	}
	n := 0
	if isNegative {
		n, _ = strconv.Atoi(string(tmp[1:]))
		n = -1 * n
	} else {
		n, _ = strconv.Atoi(string(tmp))
	}
	if n >= 2147483647 {
		return 2147483646
	}
	if n <= -2147483648 {
		return -2147483647
	}
	return n
}
