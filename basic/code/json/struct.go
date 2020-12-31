package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Stu struct {
	ID int64
}

func (s Stu) MarshalJSON() ([]byte, error) {
	fmt.Println("json")
	return json.Marshal(s)
}

func (s Stu) MarshalText() (text []byte, err error) {
	fmt.Println("text")
	return
}

type People struct {
	Stu
	Name string
}

func (p People) MarshalJSON() ([]byte, error) {
	fmt.Println("p json")
	return json.Marshal(p)
}

func (p People) MarshalText() (text []byte, err error) {
	fmt.Println("p text")
	return json.Marshal(p)
}

func main() {
	struct2()
	//js := `{"name":"11"}`
	//var p People
	//err := json.Unmarshal([]byte(js), &p)
	//if err != nil {
	//	fmt.Println("err: ", err)
	//	return
	//}
	//fmt.Println("people: ", p, err)
}

func struct2() {
	//i := time.Date(2020, 12, 20, 0, 0, 0, 0, time.UTC)
	//var t interface{}
	t := struct {
		time.Time
		People
		N int
	}{
		time.Date(2020, 12, 20, 0, 0, 0, 0, time.UTC),
		People{
			Stu{
				1,
			},
			"h",
		},
		5,
	}
	m, _ := json.Marshal(t)
	fmt.Printf("m: %s\n", m)
	//if tmj, ok := t.(encoding.TextMarshaler); ok {
	//	b, _ := tmj.MarshalText()
	//	fmt.Printf("type is identified as TextMarshaler, output is `%s`\n", b)
	//} else {
	//	fmt.Println("t doesn't has MarshalText function")
	//}
	//if tmj, ok := t.(json.Marshaler); ok {
	//	b, _ := tmj.MarshalJSON()
	//	fmt.Printf("type is identified as TextMarshaler, output is `%s`\n", b)
	//} else {
	//	fmt.Println("t doesn't has MarshalText function")
	//}
}
