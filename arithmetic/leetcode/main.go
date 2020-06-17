package main

import "fmt"

var tmp = make([][]int, 1000)
func init(){
	for i := 0; i < 1000; i++{
		tmp[i] = []int{1}
	}
}
func main() {
	fmt.Println(5 / 2)
	var n int
	n = '9' - '0'

	fmt.Println(n)

	fmt.Println(len(tmp), tmp)
}
