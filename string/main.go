package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "/a/b/c/d/a"
	//n := LastIndex(str, "/")
	//str = string([]byte(str)[:n])
	//fmt.Println(str)
	//strings.Repeat()
	//fmt.Println(strings.Replace("../../t/", "../../", "", 1))
	//
	//fmt.Println("Hello, 世界")
	//var s string
	//s = "333,"
	fmt.Println(strings.TrimRight(str, "/"))
	//fmt.Println(s)
	//s = strings.TrimRight(s, ",")
	//fmt.Println(s)
}

func LastIndex(str, sep string) int {
	return strings.LastIndex(str, sep)
}
