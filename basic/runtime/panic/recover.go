package main

import (
	"fmt"
	"runtime"
)

type Test struct {
	Name string
}

func main() {
	a()
}

func a() {
	b()
}

func b() {
	c()
}

func c() {
	defer RecoverFromPanic("func c")
	var t *Test
	fmt.Println(t.Name)
}

func RecoverFromPanic(funcName string) {
	if e := recover(); e != nil {
		buf := make([]byte, 64<<10)
		buf = buf[:runtime.Stack(buf, false)]
		fmt.Printf(" %s: %v, stack: %s", funcName, e, string(buf))
	}
}
