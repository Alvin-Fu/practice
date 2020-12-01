package main

import (
    "go/token"
    "go/scanner"
    "fmt"
)

func main(){
    var src = []byte(`println("hello world")`)
    var fset = token.NewFileSet()
    var file = fset.AddFile("hello.go", fset.Base(), len(src))
    var s scanner.Scanner
    s.Init(file, src, nil, scanner.ScanComments)
    for {
        pos, tok, lit := s.Scan()
        if tok == token.EOF{
            break
        }
        fmt.Printf("%s\t%s\t%q\n", fset.Position(pos), tok, lit)
    }
    a := token.Position{Filename: "hello.go", Line: 1, Column: 2}
    b := token.Position{Filename: "hello.go", Line: 1}
    c := token.Position{Filename: "hello.go"}

    d := token.Position{Line: 1, Column: 2}
    e := token.Position{Line: 1}
    f := token.Position{Column: 2}

    fmt.Println(a.String())
    fmt.Println(b.String())
    fmt.Println(c.String())
    fmt.Println(d.String())
    fmt.Println(e.String())
    fmt.Println(f.String())
}
