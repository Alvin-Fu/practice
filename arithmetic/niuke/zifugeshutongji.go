package main

import "fmt"

func main() {
	str := ""
	fmt.Scanf("%s", &str)
	tmp := make(map[byte]bool)
	for i := 0; i < len(str); i++ {
		if str[i] > 0 && str[i] <= 127 {
			tmp[str[i]] = true
		}
	}
	fmt.Println(len(tmp))
}

/*
题目描述
编写一个函数，计算字符串中含有的不同字符的个数。字符在ACSII码范围内(0~127)，换行表示结束符，不算在字符里。不在范围内的不作统计。
多个相同的字符只计算一次
例如，对于字符串abaca而言，有a、b、c三种不同的字符，因此输出3。
输入描述:
输入一行没有空格的字符串。

输出描述:
输出范围在(0~127)字符的个数。

示例1
输入
abc
输出
3
*/
