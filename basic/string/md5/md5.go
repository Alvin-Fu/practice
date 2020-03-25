package main

import (
	"crypto/md5"
	"fmt"
)

func main() {
	fmt.Println(fmt.Sprintf("%x", md5.Sum([]byte("hello"))))
	hash := md5.New()
	fmt.Println(string(hash.Sum([]byte("hello"))))
}
