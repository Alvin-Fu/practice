package main

import (
	"fmt"
	"time"
)

func main(){
	t := time.NewTimer(1 * time.Second)
	now := time.Now()
	fmt.Println(now)
	t.Reset(2 * time.Second)
	fmt.Println(<-t.C)
	true := false
	fmt.Println(true)
}


