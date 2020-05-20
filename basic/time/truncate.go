package main

import (
	"fmt"
	"time"
)

type People struct {
	name string `json:"name"`
}

func main() {
	t := time.Now()
	//for i := 0; i <= 24; i++{
	//	fmt.Println(i, t.Truncate(time.Hour * time.Duration(i)))
	//}
	fmt.Println(t.Format("23:59"))
	fmt.Println(t.Round(time.Hour).String())
}
