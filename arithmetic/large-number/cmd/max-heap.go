package main

import (
	"fmt"
	"practice/arithmetic/large-number"
)

func main() {
	maxHeap := large_number.NewMaxHeap(8)
	maxHeap.BuildMaxHeap([]int{3, 5, 2, 8, 1, 11, 4, 10})
	fmt.Println(maxHeap.Heap())
	maxHeap.Add(8)
	fmt.Println(maxHeap.Heap())
}
