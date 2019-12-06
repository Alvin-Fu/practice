package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"os"
	"time"
)

// 测试goroutine在释放后内存没释放的情况

func hello(w http.ResponseWriter, r *http.Request) {}
func main() {
	for i := 0; i < 1000000; i++ {
		go func() {
			time.Sleep(10 * time.Second)
		}()
	}
	http.HandleFunc("/", hello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
