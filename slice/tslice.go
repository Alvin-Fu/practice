package main

import (
	"fmt"
)

func main() {
	var arr = make([]int, 126028934144)
	fmt.Println(cap(arr), len(arr))
	//for i := 0; i < 10; i++ {
	//	arr = append(arr, i+1)
	//}
	//wg := sync.WaitGroup{}
	//var tmp = make([]int, 0)
	//wg.Add(2)
	//go func() {
	//	arr = arr[:5]
	//	wg.Done()
	//}()
	//go func() {
	//	tmp = append(tmp, arr...)
	//	wg.Done()
	//}()
	//wg.Wait()
	//fmt.Println("tmp", tmp)
	//fmt.Println("arr", arr)
}
