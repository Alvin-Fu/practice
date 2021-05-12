package main

import (
	"fmt"
	"reflect"
)

type TestA struct {
	ta int
}

func (*TestA) sub()   {}
func (t *TestA) Add() {}

func main() {
	var ta TestA
	typ := reflect.TypeOf(&ta)
	for i := 0; i < typ.NumMethod(); i++ {
		fmt.Println(typ.Method(i).Name)
	}
}
