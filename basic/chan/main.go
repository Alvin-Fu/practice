package main

import "fmt"

func main() {
	msgChan := make(chan int, 10)
	fmt.Println(len(msgChan))
	for i := 0; i < 8; i++{
		msgChan <- i
	}
	fmt.Println(len(msgChan))
	//for i := 0; i < len(msgChan); i++{
	//	j := <-msgChan
	//	//go func() {
	//	//	fmt.Println(<- msgChan)
	//	//}()
	//	fmt.Println(i, j)
	//}

	for len(msgChan) > 0{
		<-msgChan
		go func() {
			<- msgChan
		}()

	}
	fmt.Println(len(msgChan))
	//var ch = make(chan struct{})
	////var c = make(chan struct{})
	//go func() {
	//	time.Sleep(1 * time.Second)
	//	ch <- struct{}{}
	//}()
	//c := fanIn(ch)
	//<-c
	//fmt.Println("bye")
//	wg := sync.WaitGroup{}
//	wg.Add(1)
//	go func() {
//		time.Sleep(5 * time.Second)
//		close(ch)
//		wg.Done()
//	}()
//	for {
//		select {
//		case <-ch:
//			goto exit
//		}
//	}
//exit:
//	fmt.Println("bye")
}

func fanIn(ch chan struct{})chan struct{}{
	c := make(chan struct{})
	go func() {
		c <- <- ch
	}()
	return c
}