package main

import "fmt"

func main() {
	res := reversBits(1)
	fmt.Println(res)
	fmt.Println(5 >> 1)
	fmt.Println(rangeBitwiseAnd(10, 18))
}

// 颠倒给定的无符号32位整数的二进制位
func reversBits(num uint32) uint32 {
	var res uint32
	var pow = 31
	for num != 0 {
		res += (num & 1) << pow
		num >>= 1
		pow--
	}
	return res
}

func rangeBitwiseAnd(m, n int) int {
	var count int
	for m != n {
		m >>= 1
		n >>= 1
		count++
	}
	return m << count
}
