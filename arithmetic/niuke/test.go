package main

import "fmt"

func main() {
	subnetMasks := []int{255, 255, 255, 255}
	var b, j uint32
	for i := 3; i >= 0; i-- {
		b = b + (uint32(subnetMasks[i]) << uint(j*8))

		j++
	}
	b = ^b + 1
	var a uint32 = 2
	a = ^a + 1
	fmt.Println(b, b&(b-1), a&(a-1) == 0)
	if (b & (b - 1)) != 0 {
		fmt.Println("hello")
	}
}
