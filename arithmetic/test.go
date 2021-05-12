package main

import "fmt"

func main() {
	fmt.Println(A, B, C, D, E)
	//s := []int{1, 2}
	//f(s)
	//fmt.Println(s, len(s), cap(s))
}
func f(s []int) {
	s = append(s, 3)
}

const (
	A = iota
	B
	C = 1
	D
	E
)
