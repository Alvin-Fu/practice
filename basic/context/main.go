package main

import (
    "fmt"
    "context"
    "time"
)



func f(ctx context.Context) {
    context.WithValue(ctx, "foo", -6)
    time.Sleep(10 * time.Second)
}

func main() {
    ctx := context.TODO()
    f(ctx)
    fmt.Println(ctx.Value("foo"))

}
