package main

import (
	"bytes"
	"fmt"
	"sync"
)

func main() {
	arr := []int{4: 1, 4, 1: 45, 6}
	fmt.Println(arr, cap(arr))
	var x = []int{4: 44, 55, 66, 1: 77, 88, 2, 7: 6}
	println(len(x), x[2])
	//arr := []int{1, 2, 3}
	//rue := copy_append.Copy(arr)
	//rue[0] = 10
	//fmt.Println(arr, rue)
	//app()
}
func app() {
	printData := func(wg *sync.WaitGroup, data []byte) {
		defer wg.Done()
		var buff bytes.Buffer
		for _, b := range data {
			fmt.Fprintf(&buff, "%c", b)
		}

		fmt.Println(buff.String())
		fmt.Println(data)
	}
	var wg sync.WaitGroup
	wg.Add(5)
	data := []byte("golang")
	go printData(&wg, data[:3])
	go printData(&wg, data[3:])
	go printData(&wg, data[3:])
	go printData(&wg, data[3:])
	go printData(&wg, data[3:])
	wg.Wait()
	fmt.Println(string(data))
}
