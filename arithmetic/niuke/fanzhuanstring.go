package main

import "fmt"

func main() {
	str := ""
	fmt.Scanf("%s", &str)
	i := 0
	j := len(str) - 1
	tmp := []byte(str)
	for i < j {
		if tmp[i] > 'z' || tmp[i] < 'a' || tmp[j] > 'z' || tmp[j] < 'a' {
			return
		}
		tmp[i], tmp[j] = tmp[j], tmp[i]
		i++
		j--
	}
	fmt.Println(string(tmp))
}

/*
题目描述
接受一个只包含小写字母的字符串，然后输出该字符串反转后的字符串。（字符串长度不超过1000）

输入描述:
输入一行，为一个只包含小写字母的字符串。

输出描述:
输出该字符串反转后的字符串。

示例1
输入
abcd
输出
dcba
*/
