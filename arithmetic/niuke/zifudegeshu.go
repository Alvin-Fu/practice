package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var str = make([]string, 0)
	reader := bufio.NewScanner(os.Stdin)
	for reader.Scan() {
		str = append(str, reader.Text())
	}
	if len(str) != 2 {
		return
	}
	var a, b byte
	a = str[1][0]
	if str[1][0] >= 65 && str[1][0] <= 90 {
		b = a + 32
	} else {
		b = a - 32
	}
	num := 0
	for _, s := range str[0] {
		if byte(s) == a || byte(s) == b {
			num++
		}
	}
	fmt.Println(num)
}

/*
题目描述
写出一个程序，接受一个由字母、数字和空格组成的字符串，和一个字母，然后输出输入字符串中该字母的出现次数。不区分大小写。

输入描述:
第一行输入一个由字母和数字以及空格组成的字符串，第二行输入一个字母。

输出描述:
输出输入字符串中含有该字符的个数。
示例1
输入
ABCabc A
输出
2
*/
