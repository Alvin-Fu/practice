package main

import (
	"fmt"
	"practice/arithmetic/large-number"
)

func main() {
	maxHeap := large_number.NewMaxHeap(6)
	maxHeap.BuildMaxHeap([]int{3, 5, 2, 8, 11, 4})
	fmt.Println(maxHeap.Heap())
}
