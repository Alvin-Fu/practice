package main

import (
	"encoding/json"
	"fmt"
	"bytes"
	"rpcx/log"
)

func demo1(){
	var tmp1 = make(map[string]interface{})
	tmp1["count"] = 1  // int
	m1, _ := json.Marshal(tmp1)
	fmt.Printf("marshal: %s\n", string(m1))
	var tmp2 = make(map[string]interface{})
	json.Unmarshal(m1, &tmp2)
	fmt.Printf("value: %v\n", tmp2["count"])
	fmt.Printf("type: %T\n", tmp2["count"])
}

func decoderDemo(){
	var tmp1 = make(map[string]interface{})
	tmp1["count"] = 1  // int
	m1, _ := json.Marshal(tmp1)
	fmt.Printf("marshal: %s\n", string(m1))
	var tmp2 = make(map[string]interface{})
	decoder := json.NewDecoder(bytes.NewReader(m1))
	decoder.UseNumber()
	decoder.Decode(&tmp2)

	fmt.Printf("value: %v\n", tmp2["count"])
	fmt.Printf("type: %T\n", tmp2["count"])
	num, err  := tmp2["count"].(json.Number).Int64()
	if err != nil {
		log.Fatalf("err: %v", err)
	}
	fmt.Printf("type: %T\n", num)
}

func main(){
	decoderDemo()
}
