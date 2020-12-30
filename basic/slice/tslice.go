package main

import "fmt"

func main() {
	a := make([]int64, 0)
	for i := 0; i < 10; i++ {
		a = append(a, int64(i+1))
	}
	a = Arr(a)
	fmt.Println(len(a), cap(a), a)
}

func Arr(tmp []int64) []int64 {
	for i := 0; i < 5; i++ {
		tmp[i] = 1
	}
	tmp = tmp[:5]
	return tmp
}
