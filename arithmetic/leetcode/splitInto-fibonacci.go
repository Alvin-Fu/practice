package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	splitIntoFibonacci("123456579")
}

func splitIntoFibonacci(S string) []int {
	if len(S) < 3 {
		return nil
	}
	rue := make([]int, 0)
	backtrack([]byte(S), rue, 0)
	fmt.Println(rue)
	return rue
}

func backtrack(data []byte, rue []int, index int) bool {
	if len(data) < index && len(rue) >= 3 {
		return true
	}
	for i := index; i < len(data); i++ {
		if data[index] == '0' && i > index {
			break
		}

		n, _ := strconv.Atoi(string(data[index:i]))
		if n > math.MaxInt32 {
			break
		}
		rueLen := len(rue)
		if rueLen >= 2 && n > rue[rueLen-1]+rue[rueLen-2] {
			break
		}
		if rueLen <= 1 || (n == rue[rueLen-1]+rue[rueLen-2]) {
			rue = append(rue, n)
			if backtrack(data, rue, i+1) {
				return true
			}
			rue = rue[:len(rue)-1]
		}
	}
	return false
}
