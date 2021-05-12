package main

import (
    "go.uber.org/ratelimit"
    "time"
    "fmt"
    //"github.com/uber-go/ratelimit"
    "sync"
)

func main()  {
    rl := ratelimit.New(1000) // per second

    prev := time.Now()
    wg := sync.WaitGroup{}
    wg.Add(10000)
    for i := 0; i < 10000; i++ {
        go func (i int, prev time.Time){
            now := rl.Take()
            if i > 0 {
                fmt.Println(i, now.Sub(prev))
            }
            prev = now
            wg.Done()
        }(i, prev)
    }
    wg.Wait()
}
