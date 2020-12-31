package main

import (
	"context"
	"encoding/json"
	"fmt"
)

type valueCtx struct {
	Key, Val interface{}
}

func main() {
	ctx := context.WithValue(context.Background(), "a", "b")
	fmt.Println(ctx.Value("a"))
	data, _ := json.Marshal(ctx)
	fmt.Println(string(data))
	vs := &valueCtx{
		Key: 1,
		Val: 2,
	}
	data, _ = json.Marshal(vs)
	fmt.Println(string(data))
}
