package main

import (
	"fmt"
	"math"
)

/*
 * 统计所有小于非负整数 n 的质数的数量。
 * 输入：n = 10
 * 输出：4
 * 解释：小于 10 的质数一共有 4 个, 它们是 2, 3, 5, 7 。
 */

func main() {
	fmt.Println(countPrimes(40))
}

func countPrimes(n int) int {
	if n <= 1 {
		return 0
	}
	if n > 5*int(math.Pow10(6)) {
		return 0
	}
	if n < 5 {
		if n < 3 {
			return 0
		} else {
			return 2
		}
	}
	count := 2
	for i := 5; i < n; i++ {
		if isPrimes(i) {
			count++
		}
	}
	return count
}

func isPrimes(n int) bool {
	if !(n%6 == 1 || n%6 == 5) {
		return false
	}
	temp := int(math.Sqrt(float64(n)))
	for i := 5; i <= temp; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}
