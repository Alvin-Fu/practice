package main

import "fmt"

func main(){
	var arrary = []int{1, 2, 3, 4, 5, 9, 0, 8, 7, 6, 9}
	tmp := 0
	for _, a := range arrary{
		tmp ^= a
	}
	fmt.Println(tmp)
	Xor()
}


func Xor(){
	var tmp = make([]int, 0)
	for i := 1; i <= 1000; i ++{
		tmp = append(tmp, i)
	}
	num := 0

	for _, t := range tmp{
		if t == 800 || t == 100{
			continue
		}
		num ^= t
	}
	fmt.Println(num)
	fmt.Println(num ^ 1000)
}