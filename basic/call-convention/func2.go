package main

func main(){
    var l0 int = 0x3
    l1, l2 := myFunc2(0x7, 0x11)
    var l3 int = l1 + 0x5
    _ = l0 + l2 + l3

}

func myFunc2(a1 int, a2 int)(int, int){
    return a1, a2 + 0x13
}
