package main

import (
    "runtime"
    "time"
    "fmt"
)
type Garbage struct {
    a int
}

func notify(g *Garbage) {
    stats := &runtime.MemStats{}
    runtime.ReadMemStats(stats)
    fmt.Println(time.Unix(0, int64(stats.LastGC)))
}

func ProduceFinalizedGarbage(){
    g := &Garbage{}
    runtime.SetFinalizer(g, notify)
}

func main(){
    go ProduceFinalizedGarbage()

   for {
       runtime.GC()
       notify(nil)
       time.Sleep(30 * time.Second)

   }
}
