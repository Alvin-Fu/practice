package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	line, _, _ := reader.ReadLine()
	n, err := strconv.Atoi(string(line))
	if err != nil {
		fmt.Println(-1)
		return
	}
	if n > 50 {
		fmt.Println(-1)
		return
	}
	var weight = make([][]int, n)
	for i := 0; i < n; i++ {
		line, _, _ := reader.ReadLine()
		for j, s := range strings.Split(string(line), " ") {
			n, err := strconv.Atoi(string(line))
			if err != nil {
				fmt.Println(-1)
				return
			}
			weight[i] = append(weight[i], n)
		}
	}
}
