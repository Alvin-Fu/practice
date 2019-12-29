package main

import "fmt"

func main(){
	b1, b2 := byte(9), byte(9)
	b := b1 * b2
	fmt.Println(b / 10)
	fmt.Println(b %10)
}
