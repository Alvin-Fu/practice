package main

import (
    "math/rand"
    "encoding/binary"
    crand "crypto/rand"
    "fmt"
)

func main(){
    //var randomSeed = make([]int, 52)
    //for n := 0; n < 52; n++ {
    //    randomSeed[n] = n
    //}
    //rand.Shuffle(52, func(i, j int) {
    //    randomSeed[i], randomSeed[j] = randomSeed[j], randomSeed[i]
    //})
    for i := 0; i < 100; i++ {
        fmt.Println(float64(GetOneRand(35)) / 10)
    }

}


//根据scope获取一个随机数，返回是(0, scope]
func GetOneRand(scope int64) int64 {
    if scope <= 0 {
        return 1
    }
    var n uint64
    binary.Read(crand.Reader, binary.LittleEndian, &n)
    rand.Seed(int64(n))
    return rand.Int63n(scope)
}