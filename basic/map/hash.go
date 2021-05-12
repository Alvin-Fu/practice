package main

import (
    "hash/fnv"
    "fmt"
)


func main(){
    fmt.Println(MakeSvrType("alloc"))
}
func MakeSvrType(service string) int64 {
    return HashStr2Int64(fmt.Sprintf("/%s/%s/%s", "hall", "vietnam", service))
}
// HashStr2Int64 calculate a hash value of a string into a int64.
func HashStr2Int64(input string) int64 {
    // avoid to overflow a int64, we hash string to uint32 first.
    h := fnv.New32a()
    h.Write([]byte(input))
    return int64(h.Sum32())
}