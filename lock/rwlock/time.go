package main

import (
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"time"
)

var t time.Duration = 60

//#####PTHE
func CpuProfile() {
	f, err := os.OpenFile("cpu.prof", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	log.Println("CPU Profile started")
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()
	time.Sleep(t * time.Second)
	fmt.Println("CPU Profile stopped")
}

func main() {
	go func() {
		CpuProfile()
	}()
	ticker := time.NewTicker(time.Millisecond)
	defer ticker.Stop()
	for range ticker.C {
		//time.Sleep(1 * time.Second)
	}
}
