package main

import (
		"time"
	"unsafe"
	"fmt"
)
type MySelect struct {
	Num int
}

func Chann(ch chan int, stopCh chan bool) {
	var i int
	i = 10
	for j := 0; j < 10; j++ {
		ch <- i
		time.Sleep(time.Second)
	}
	stopCh <- true
}


func main() {
	sle := make([]MySelect, 2)
	sle[0] = MySelect{Num: 1}
	sle[1] = MySelect{Num: 2}
	sle = append(sle, MySelect{Num: 3}, MySelect{Num:4})
	order := make([]uint16, len(sle)/2)
	testSelect(&sle[0], &order[0],len(sle))

//	ch := make(chan int)
//	c := 0
//	stopCh := make(chan bool)
//
//	go Chann(ch, stopCh)
//
//	for {
//		select {
//		case c = <-ch:
//			fmt.Println("Recvice 1", c)
//			fmt.Println("channel")
//		case s := <-ch:
//			fmt.Println("Receive 2 ", s)
//		case _ = <-stopCh:
//			goto end
//		}
//	}
//end:
}

func testSelect(s *MySelect, m *uint16, n int){
	t := (*[1<<16]MySelect)((unsafe.Pointer(s)))
	f := (*[1<<17]uint64)(unsafe.Pointer(m))
	fmt.Println(t[:n])
	fmt.Println(f[:n:n])
	fmt.Println(f[n:][n:n])
}
