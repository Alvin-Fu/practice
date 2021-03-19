package main

import "fmt"

func main() {
	fmt.Println(bitCount(-2))
}

func bitCount(n int) int {
	count := 0
	for n != 0 {
		n = n & (n - 1)
		count++
	}
	return count
}
