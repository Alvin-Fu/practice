package main

import (
	"encoding/json"
	"fmt"
		"rpcx/log"
)

type Number struct {
	ID int64 `json:"id,string"`
	Chip float64 `json:"chip,string"`
}

func main(){
	jsonStr := `{"id":"1234", "chip":"87.9"}`
	n := new(Number)
	err := json.Unmarshal([]byte(jsonStr), n)
	if err != nil {
		log.Fatalf("unmarshal err: %v", err)
	}
	fmt.Printf("unmarshal number: %+v", n)
}
