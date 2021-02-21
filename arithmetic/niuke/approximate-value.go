package main

import (
	"fmt"
	"strconv"
)

func main() {
	value := ""
	fmt.Scanf("%s", &value)
	for i, v := range value {
		if v == '.' {
			n, _ := strconv.Atoi(value[:i])
			if value[i+1]-48 >= 5 {
				n += 1
			}
			fmt.Println(n)
			break
		}
	}
}

/*
题目描述
写出一个程序，接受一个正浮点数值，输出该数值的近似整数值。如果小数点后数值大于等于5,向上取整；小于5，则向下取整。

输入描述:
输入一个正浮点数值

输出描述:
输出该数值的近似整数值

示例1
输入
5.5
输出
6
*/
