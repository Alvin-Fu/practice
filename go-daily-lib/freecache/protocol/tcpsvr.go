package protocol

import (
	"net"
	"practice/go-daily-lib/freecache/uitl"
		"runtime"

	"log"
)

type TCPHandel interface {
	Handle(net.Conn)
}

func TCPSever(listener net.Listener, handler TCPHandel)error{
	waitGroup := uitl.WaitGroupWrapper{}
	for {
		conn, err := listener.Accept()
		if err != nil {
			if nerr, ok := err.(net.Error); ok && nerr.Temporary() {
				runtime.Gosched()
				continue
			}
			// TODO:

			break

		}
		// 处理放在Handle中
		waitGroup.Wrap(func() {
			handler.Handle(conn)
		})
	}
	log.Printf("tcp sever closing")
	waitGroup.Wait()
	log.Printf("tcp sever quit")
	return nil
}