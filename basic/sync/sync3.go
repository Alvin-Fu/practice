package main

import "fmt"

func main() {
	n := 3
	i := Fibo1(n)
	fmt.Printf("%v ", i)
}

//递归实现
func Fibo1(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else if n > 1 {
		return Fibo1(n-1) + Fibo1(n-2)
	} else {
		return -1
	}
}

