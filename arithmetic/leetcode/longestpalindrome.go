package main

import (
    "strings"
    "unsafe"
    "reflect"
    "fmt"
)

func longestPalindrome(s string) string {
    if len(s) > 1000 {
        return ""
    }
    data := StringToBytesHeader(s)
    tmp := string(data[:1])
    for i, d := range data {
        index :=strings.LastIndexByte(s, d)
        if index != i{
            index1 := i + 1
            index2 := index - 1
            if len(tmp) > index - i + 1{
                continue
            }
            for  {
                if index1 == index2 || index1 > index2 {
                   tmp = string(data[i:index + 1])
                    break
                }
                if data[index1] != data[index2]{
                    break
                }
                index1 ++
                index2 --
            }
        }

    }
    return tmp
}
func StringToBytesHeader(str string)[]byte{
    sh := (*reflect.StringHeader)(unsafe.Pointer(&str))
    bh := reflect.SliceHeader{
        Data: sh.Data,
        Len: sh.Len,
        Cap: sh.Len,
    }
    return *(*[]byte)(unsafe.Pointer(&bh))
}
func main() {
    fmt.Println(longestPalindrome("aaabaaaa"))
}

