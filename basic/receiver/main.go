package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Uid     int64
	ClassId int64
	Score   int64
}

func (s Student) Notify() {
	s.Score = 1
	fmt.Println("notify")
}

func (s *Student) Broadcast(id int64) {
	s.ClassId = 1
	fmt.Println("broadcast")
}

func main() {
	s := &Student{
		Score:   2,
		ClassId: 2,
	}
	t := reflect.TypeOf(s)
	fmt.Printf("function count: %d\n", t.NumMethod())
	m := t.Method(1)
	mt := m.Type
	fmt.Printf("%s function %d args", m.Name, mt.NumIn())

}
