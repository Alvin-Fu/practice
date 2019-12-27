package main

import (
	"fmt"
	"practice/slice/copy-append"
)

func main() {
	arr := []int{1, 2, 3}
	rue := copy_append.Copy(arr)
	rue[0] = 10
	fmt.Println(arr, rue)
}
