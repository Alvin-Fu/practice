package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewScanner(os.Stdin)
	reader.Text()
	tmp := make(map[uint8][]string)
	for i := 0; i < 26+26; i++ {
		tmp[uint8(i)] = make([]string, 0)
	}
	for reader.Scan() {
		s := reader.Text()
		if s[0] > 'Z' {
			tmp[s[0]-'a'+26] = append(tmp[s[0]-'a'+26], s)
		} else {
			tmp[s[0]-'A'] = append(tmp[s[0]-'A'], s)
		}
	}
	for i := 0; i < 26+26; i++ {
		str, ok := tmp[uint8(i)]
		if ok {
			str = compare(str)
			for _, s := range str {
				fmt.Println(s)
			}
		}
	}
}

func compare(str []string) []string {
	for i := 0; i < len(str); i++ {
		for j := 0; j < len(str)-i-1; j++ {
			count := len(str[j])
			if len(str[j]) > len(str[j+1]) {
				count = len(str[j+1])
			}
			if count == 1 {
				if len(str[j+1]) == 1 {
					str[j], str[j+1] = str[j+1], str[j]
				}
				continue
			}
			for k := 0; k < count; k++ {
				if str[j][k] > str[j+1][k] {
					str[j], str[j+1] = str[j+1], str[j]
					break
				} else if str[j][k] < str[j+1][k] {
					break
				}
			}

		}
	}
	return str
}

/*
题目描述
给定n个字符串，请对n个字符串按照字典序排列。
输入描述:
输入第一行为一个正整数n(1≤n≤1000),下面n行为n个字符串(字符串长度≤100),字符串中只含有大小写字母。
输出描述:
数据输出n行，输出结果为按照字典序排列的字符串。
示例1
输入

9
cap
to
cat
card
two
too
up
boat
boot
输出

boat
boot
cap
card
cat
to
too
two
up
*/
