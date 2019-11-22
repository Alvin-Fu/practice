package main

import "fmt"

type MyIFace interface {
	Print()
}

type MySt struct {
}

func (m *MySt) Print() {

}

func main() {
	x := 1
	var y interface{} = x
	var m *MySt
	var t MyIFace = m
	fmt.Println(y, t)
	var a int = 1
	var arr = make([]int, 0)
	arr = append(arr, 1, 2, 34, 5, 6, 7)
	fmt.Println(arr[:a:a])
}
