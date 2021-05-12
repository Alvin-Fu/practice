package main

import "fmt"

type Tmp struct {

}

func main(){
    tmp := Tmp{}
    Test(tmp)

}

func Test(tmp interface{}){
    switch c := tmp.(type) {
    // TODO(LeoQin): make business notify and send to client with protoNotify
    case struct{}:
        fmt.Println("struct")
    default:
       fmt.Println(c)
    }
}