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
		if len(text) > 20 {
			continue
		}
		fmt.Println(delete([]byte(text)))
	}
}

func delete(data []byte) string {
	temp := make(map[byte]int)
	for i := 0; i < 26; i++ {
		temp[i]++
	}
	min := 20
	for _, n := range temp {
		if min > n {
			min = n
		}
	}
	rue := make([]byte, 0)
	for _, d := range data {
		if temp[d] > min {
			rue = append(rue, d)
		}
	}
	return string(rue)
}

/*
题目描述
实现删除字符串中出现次数最少的字符，若多个字符出现次数一样，则都删除。
输出删除这些单词后的字符串，字符串中其它字符保持原来的顺序。
注意每个输入文件有多组输入，即多个字符串用回车隔开
输入描述:
字符串只包含小写英文字母, 不考虑非法输入，输入的字符串长度小于等于20个字节。

输出描述:
删除字符串中出现次数最少的字符后的字符串。

示例1
输入
abcdd
aabcddd
输出
dd
aaddd
*/
