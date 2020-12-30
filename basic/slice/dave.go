package main

import "fmt"

func f(s []string, level int) {
	if level > 8 {
		return
	}
	s = append(s, fmt.Sprint(level))
	f(s, level+1)
	fmt.Println("level:", level, "slice:", s, cap(s))
	if level == 4 {
		fmt.Println(s[:cap(s)], len(s))
	}
}

func main() {

	f(nil, 0)
}
