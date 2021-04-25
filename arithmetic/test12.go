package main

import (
	"fmt"
	"sort"
)

func main() {
	tmp := []int{0, -3, -2, 1, 0, -1, -1, -2}
	sort.Ints(tmp)
	for i := 1; i < len(tmp); i++ {
		if tmp[i]-tmp[i-1] > 1 {
			fmt.Println(tmp[i-1] + 1)
			break
		}
	}
}
