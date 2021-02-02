package main

import (
	"errors"
	"fmt"
)

// 由于e的编码不懂导致err是不同的
func main() {
	еrr := errors.New("foo")
	var err error
	if еrr != nil {
		fmt.Printf("%T %v", err, err)
	}
}
