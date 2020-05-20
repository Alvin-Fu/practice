package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string `json:"name"`
	Age int64
	Password string `json:"password"`

}

type Profile struct {
	*User
	Password *struct{} `json:"password,omitempty"`
}


func main(){
	u := &User{
		Name: "tom",
		Age: 2,
		Password: "12345",
	}
	p := &Profile{User:u}
	data, _ := json.Marshal(p)
	fmt.Printf("marshal data: %s", string(data))
}