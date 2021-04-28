package main

import (
	"encoding/json"
	"os"
	"fmt"
)

type People struct {
	Name string `json:"name"`
	Age *int64
	Password string `json:"password,omitempty"`
}

func main(){
	p := &People{
		Name: "tom",
		//Age: 2,
	}
	data, err := json.Marshal(p)
	if err != nil {
		os.Exit(-1)
	}
	fmt.Printf("marshal data: %s", string(data))
}


