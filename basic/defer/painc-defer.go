package main

import "fmt"

func main() {
	pt()
}

func pt() {
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
	panic("h")
}
