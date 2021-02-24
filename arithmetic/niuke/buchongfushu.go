package main

import (
	"fmt"
)

func main() {
	num := 0
	fmt.Scanf("%d", &num)
	tmp := make(map[int]bool)
	rue := 0
	for num > 0 {
		n := num % 10
		num /= 10
		if tmp[n] {
			continue
		}
		tmp[n] = true
		rue = rue*10 + n
	}
	fmt.Println(rue)
}

/*
题目描述
输入一个int型整数，按照从右向左的阅读顺序，返回一个不含重复数字的新的整数。
保证输入的整数最后一位不是0。
输入描述:
输入一个int型整数

输出描述:
按照从右向左的阅读顺序，返回一个不含重复数字的新的整数

示例1
输入
9876673
输出
37689
*/
