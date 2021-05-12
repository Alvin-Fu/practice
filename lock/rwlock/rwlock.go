package main

import (
	"fmt"
	"os"
	"runtime/trace"
	"sync"
	"time"
)

// 加读锁的时候超时超过了100ms

var mlock sync.RWMutex
var wg sync.WaitGroup

func main() {
	trace.Start(os.Stderr)
	defer trace.Stop()
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go gets()
	}
	wg.Wait()
}

func gets() {
	for i := 0; i < 100000; i++ {
		get(i)
	}
	wg.Done()
}

func get(i int) {
	beginTime := time.Now()
	mlock.RLock()
	tmp := time.Since(beginTime).Nanoseconds() / 1000000
	if tmp > 100 {
		fmt.Println("fuck here")
	}
	mlock.RUnlock()
}

//相关博客https://xargin.com/a-rlock-story/
