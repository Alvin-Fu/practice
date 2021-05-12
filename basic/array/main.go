package main

import "fmt"

type Test struct {
	name string
}

func main() {
	a1 := [5]int{1, 2, 3, 4}
	a2 := [5]int{1, 2, 3, 4, 5}
	if a1 == a2 {
		fmt.Println("hello")
	}
	as1 := [1]int{1}
	as2 := [...]int{}
	if as1 == as2 {
		fmt.Println("as")
	}
	t1 := new(Test)
	t2 := new(Test)
	if t1 == t2 {
		fmt.Println("hello")
	}
	tn1 := Test{
		name: "",
	}
	tn2 := Test{
		name: "",
	}
	if tn1 == tn2 {
		fmt.Println("world")
	}
	s1 := []int{}
	//s2 := []int{}
	//if s1 == s2 {
	//	fmt.Println("s1")
	//}
	if s1 == nil {
		fmt.Println("nil")
	}

}
