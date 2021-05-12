package main

import (
    "golang.org/x/text/message"
    "golang.org/x/text/language"
    _"golang.org/x/text/message/catalog"
    "fmt"
)

// NOTE: change -lang option to add more languages support.
//go:generate gotext -srclang=en update -out=catalog.go -lang=en,zh-hans,zh-hant



var acceptPrinter = make(map[string]*message.Printer)
func main(){
    for _, tag := range message.DefaultCatalog.Languages() {
        fmt.Println(tag.String())
       acceptPrinter[tag.String()] = message.NewPrinter(tag)
    }
    //p := GetPrinter(language.English.String())
    //p.Printf("login multiple times")
    //p.Println()
    //p.Printf("client request error")
    //p.Println()
    ph := GetPrinter("zh-Hant")
    ph.Printf("login multiple times")

}
// GetPrinter return a printer by language string, if printer not found, return a default printer.
func GetPrinter(lang string) *message.Printer {
    if len(lang) == 0 || lang == "default" {
        fmt.Println(lang)
        return acceptPrinter[ language.English.String()]
    }

    p, ok := acceptPrinter[lang]
    if !ok {
        fmt.Println(lang)
        return acceptPrinter[ language.English.String()]
    }
    return p
}
