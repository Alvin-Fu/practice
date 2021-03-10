package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	target, _, _ := reader.ReadLine()
	source, _, _ := reader.ReadLine()
	fmt.Println(find(target, source))
}

func find(target, source []byte) int {
	if len(target) > len(source) || len(target) > 100 || len(source) > 500000 {
		return -1
	}
	index := len(source) - 1
	count := 0
	for i := len(target) - 1; i >= 0; i-- {
		for j := index; j >= 0; j-- {
			if target[i] == source[j] {
				index = j - 1
				count++
				break
			}
		}
	}
	if count == len(target) {
		return index + 1
	}
	return -1
}
