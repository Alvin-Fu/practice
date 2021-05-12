package main

import "fmt"

func main() {
	CapLen()
}
func CapLen() {
	s := []int{1, 2}
	fmt.Printf("len=%d, cap=%d\n", len(s), cap(s))
	s = append(s, 1, 2, 3)
	fmt.Printf("len=%d, cap=%d\n", len(s), cap(s))
}
