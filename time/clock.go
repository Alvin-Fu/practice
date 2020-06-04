package main

import (
    "time"
    "fmt"
)

func main(){
    t := time.Now().Unix()
    fmt.Println(time.Unix(t, 0).Format("15:04"))

}
