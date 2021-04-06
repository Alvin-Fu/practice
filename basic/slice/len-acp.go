package main

import "fmt"

func main() {

	s := make([]int, 10, 10)
	for i := 1; i <= 20; i++ {
		s = append(s, i)
	}
	fmt.Println(s, len(s), cap(s))

	s2 := make([]int, 0)
	for i := 1; i <= 20; i++ {
		s2 = append(s2, i)
	}
	s3 := make([]int, 10, 10)
	s4 := append(s3, s2...)
	fmt.Println(s4, len(s4), cap(s4))

	s5 := make([]int, 10, 10)
	for _, s := range s2 {
		s6 := append(s5, s)
		fmt.Printf("len: %d cap: %d ", len(s6), cap(s6))
	}
	fmt.Println(s5, len(s5), cap(s5), cap(s2))

}
