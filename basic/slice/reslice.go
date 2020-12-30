package main

import "fmt"

func a() []int {
	a1 := []int{3}
	fmt.Println(len(a1), cap(a1))
	a2 := a1[1:]
	return a2
}

func main() {
	fmt.Println(a())
}
