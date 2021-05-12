package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

type Option func(str string) int

func compileTime() time.Time {
	info, err := os.Stat(os.Args[0])
	if err != nil {
		return time.Now()
	}
	return info.ModTime()
}

func main() {
	fmt.Println(compileTime().String())
	fmt.Println(time.Now().String())
	n := strToInt("123", Atoi)
	fmt.Println(n)
}
func Atoi(str string) int {
	n, _ := strconv.Atoi(str)
	return n
}
func strToInt(str string, option Option) int {
	return option(str)
}
