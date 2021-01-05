package main

import (
	"fmt"
	"github.com/gogo/protobuf/proto"
	"practice/basic/code/pb/proto"
)

func main() {
	t1 := &pbTest.Test1{
		A: proto.Int32(200000),
	}
	data, _ := t1.Marshal()
	fmt.Println(data)
	t1.Unmarshal(data)
}
