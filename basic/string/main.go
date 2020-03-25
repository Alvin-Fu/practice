package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	str := "hello"
	//n := LastIndex(str, "/")
	//str = string([]byte(str)[:n])
	//fmt.Println(str)
	//strings.Repeat()
	fmt.Println(utf8.DecodeLastRuneInString(str))
	fmt.Println(strings.Replace(str, "l", "23", 3))
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
