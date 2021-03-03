package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(encrypt(scanner.Text()))
		fmt.Println(decrypt(scanner.Text()))
	}
}

func encrypt(str string) string {
	data := []byte(str)
	for i := 0; i < len(data); i++ {
		d := data[i]
		t := d
		if d >= 'a' && d < 'z' {
			t = d - 31
		} else if d == 'z' {
			t = 'A'
		} else if d == 'Z' {
			t = 'a'
		} else if d >= 'A' && d < 'Z' {
			t = d + 33
		} else if d >= '0' && d < '9' {
			t = d + 1
		} else if d == '9' {
			t = '0'
		}
		data[i] = t
	}
	return string(data)

}

func decrypt(str string) string {
	data := []byte(str)
	for i := 0; i < len(data); i++ {
		d := data[i]
		t := d
		if d > 'a' && d <= 'z' {
			t = d - 33
		} else if d == 'a' {
			t = 'Z'
		} else if d == 'A' {
			t = 'z'
		} else if d > 'A' && d <= 'Z' {
			t = d + 31
		} else if d > '0' && d <= '9' {
			t = d - 1
		} else if d == '0' {
			t = '9'
		}
		data[i] = t
	}
	return string(data)
}

/*
题目描述
1、对输入的字符串进行加解密，并输出。

2、加密方法为：

当内容是英文字母时则用该英文字母的后一个字母替换，同时字母变换大小写,如字母a时则替换为B；字母Z时则替换为a；

当内容是数字时则把该数字加1，如0替换1，1替换2，9替换0；

其他字符不做变化。

3、解密方法为加密的逆过程。


本题含有多组样例输入。
输入描述:
输入说明
输入一串要加密的密码
输入一串加过密的密码

输出描述:
输出说明
输出加密后的字符
输出解密后的字符

示例1
输入
abcdefg
BCDEFGH
输出
BCDEFGH
abcdefg
*/
