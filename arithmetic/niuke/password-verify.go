package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		if len(text) <= 8 {
			fmt.Println("NG")
			continue
		}
		tmp := make(map[int]struct{})
		for i, t := range text {
			if len(tmp) >= 3 {
				break
			}
			if t >= '0' && t <= '9' {
				tmp[0] = struct{}{}
			} else if t >= 'a' && t <= 'z' {
				tmp[1] = struct{}{}
			} else if t >= 'A' && t <= 'Z' {
				tmp[2] = struct{}{}
			} else {
				tmp[3] = struct{}{}
			}
		}
		if len(tmp) <= 3 {
			fmt.Println("NG")
		}
		for i := 1; i < len(tmp); {
			if tmp[i] != tmp[i-1] {

			}
		}
	}

}

/*
题目描述
密码要求:

1.长度超过8位
2.包括大小写字母.数字.其它符号,以上四种至少三种
3.不能有相同长度大于2的子串重复

输入描述:
一组或多组长度超过2的子符串。每组占一行

输出描述:
如果符合要求输出：OK，否则输出NG

示例1
输入
021Abc9000
021Abc9Abc1
021ABC9000
021$bc9000
输出
OK
NG
NG
OK
*/
