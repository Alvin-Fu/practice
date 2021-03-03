package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		data := scanner.Text()
		fmt.Println(sortWord([]byte(data)))
	}

}

func sortWord(data []byte) string {
	temp := make(map[byte][]byte)
	for _, d := range data {
		t := d
		if d >= 'a' && d <= 'z' {
			t = d - 32
		}
		if t >= 'A' && t <= 'Z' {
			if _, ok := temp[t]; !ok {
				temp[t] = make([]byte, 0)
			}
			temp[t] = append(temp[t], d)
		}
	}
	rue := make([]byte, len(data))
	var t byte = 'A'
	for i, d := range data {
		t1 := d
		if d >= 'a' && d <= 'z' {
			t1 = d - 32
		}
		if t1 < 'A' || t1 > 'Z' {
			rue[i] = d
			continue
		}
		for t < 'z'+1 {
			rt, ok := temp[t]
			if ok && len(rt) > 0 {
				rue[i] = rt[0]
				if len(rt) == 1 {
					delete(temp, t)
					t++
				} else {
					temp[t] = rt[1:]
				}
				break
			}
			t++
		}
	}
	return string(rue)
}

/*
题目描述
编写一个程序，将输入字符串中的字符按如下规则排序。

规则 1 ：英文字母从 A 到 Z 排列，不区分大小写。
如，输入： Type 输出： epTy

规则 2 ：同一个英文字母的大小写同时存在时，按照输入顺序排列。
如，输入： BabA 输出： aABb

规则 3 ：非英文字母的其它字符保持原来的位置。
如，输入： By?e 输出： Be?y

注意有多组测试数据，即输入有多行，每一行单独处理（换行符隔开的表示不同行）

输入描述:
输入字符串
输出描述:
输出字符串
示例1
输入
复制
A Famous Saying: Much Ado About Nothing (2012/8).
输出
复制
A aaAAbc dFgghh: iimM nNn oooos Sttuuuy (2012/8).
*/
