package main

import (
	"fmt"
	"sync"
)

func Get() {

}

type ArrayT struct {
	mtx   sync.Mutex
	array []int
}

func (at *ArrayT) Del(b int) {
	at.mtx.Lock()
	defer at.mtx.Unlock()
	for i, a := range at.array {
		if a == b {
			at.array = append(at.array[:i], at.array[i+1:]...)
			//fmt.Println("arr1", at.array)
			break
		}
	}
}

func main() {
	//for j := 0; j < 10000000; j++ {
	//	var arr = make([]int, 0)
	//	for i := 0; i < 10; i++ {
	//		arr = append(arr, i+1)
	//	}
	//	arrayT := &ArrayT{
	//		array: arr,
	//	}
	//	del := 0
	//	//fmt.Println(arr)
	//	wg := sync.WaitGroup{}
	//	wg.Add(3)
	//	go func() {
	//		arrayT.Del(2)
	//		wg.Done()
	//	}()
	//	go func() {
	//		arrayT.Del(3)
	//		wg.Done()
	//	}()
	//	go func() {
	//		//arrayT.mtx.Lock()
	//		var arr2 []int
	//		arr2 = append(arr2, arrayT.array...)
	//		//arrayT.mtx.Unlock()
	//		flage := false
	//		for _, a := range arr2 {
	//			//fmt.Println(a)
	//			if a == 7 {
	//				del = a
	//				flage = true
	//				break
	//			}
	//		}
	//		if !flage {
	//			fmt.Println("arr2: ", arr2)
	//			fmt.Println("array: ", arrayT.array)
	//		}
	//		wg.Done()
	//	}()
	//
	//	wg.Wait()
	//	arrayT.Del(del)
	//	if len(arrayT.array) != 7 {
	//		fmt.Println(arrayT.array)
	//	}
	//}


	arr := []int{2, 3, 1, 6, 8}
	for i, a := range arr{
		fmt.Printf("index: %d, value: %d\t", i, a)
	}
	fmt.Printf("\n")
}