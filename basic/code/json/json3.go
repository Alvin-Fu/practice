package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string `json:"name"`
	Age int64
	Password string `json:"password,omitempty"`
	*Profile  `json:"profile,omitempty"`
}

type Profile struct {
	Website string `json:"site"`
	Slogan string `json:"slogan"`
}

func main(){
	u := &User{
		Name: "tom",
		Age: 2,
		Password: "12345",
	}
	data, _ := json.Marshal(u)
	fmt.Printf("marshal data: %s", string(data))
}
