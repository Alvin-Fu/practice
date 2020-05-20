package main

import (
    "golang.org/x/text/message"
        "golang.org/x/text/language"
    "os"
    "fmt"
)
// NOTE: change -lang option to add more languages support.
//go:generate gotext -srclang=en update -out=catalog.go -lang=en,zh-hans,zh-hant
func main(){
    printer := make(map[string]*message.Printer)
    tag := language.Make("zh_hans")
    dp := message.NewPrinter(tag)
    if dp == nil {
        os.Exit(1)
    }
    printer["zh_hans"] = dp
    fmt.Println(printer["zh_hans"])
    fmt.Println(message.DefaultCatalog.Languages())
}
