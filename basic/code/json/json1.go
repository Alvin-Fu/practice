package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main(){
	var tmp1 []int
	tmp2 := make([]int, 0)
	data1, err := json.Marshal(tmp1)
	if err != nil {
		log.Fatalf("marshal err: %v", err)
		return
	}
	fmt.Println(string(data1))
	fmt.Println(data1)
	data2, err := json.Marshal(tmp2)
	if err != nil {
		log.Fatalf("marshal err: %v", err)
		return
	}
	fmt.Println(string(data2))
	fmt.Println(data2)
}

