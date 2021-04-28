package main

import (
	"fmt"
	"strings"
)

/*
题目描述
•连续输入字符串，请按长度为8拆分每个字符串后输出到新的字符串数组；
•长度不是8整数倍的字符串请在后面补数字0，空字符串不处理。

输入描述:
连续输入字符串(输入多次,每个字符串长度小于100)

输出描述:
输出到长度为8的新字符串数组
示例1
输入
abc
123456789
输出
abc00000
12345678
90000000
*/

func main() {
	//reader := bufio.NewScanner(os.Stdin)
	//rue := make([]string, 0)
	//for reader.Scan() {
	//	s := reader.Text()
	//	r := []byte{48, 48, 48, 48, 48, 48, 48, 48}
	//	j := 0
	//	for i := 0; i < len(s); {
	//		if j >= 8 {
	//			rue = append(rue, string(r))
	//			j = 0
	//			r = []byte{48, 48, 48, 48, 48, 48, 48, 48}
	//		}
	//		r[j] = s[i]
	//		i++
	//		j++
	//	}
	//	if j < 8 {
	//		rue = append(rue, string(r))
	//	}
	//}
	//fmt.Println(rue)
	var str string
	str = `12345678
1234
12345666
232332323
123456
12345678
24141111
6161616161
102345678`
	tmp := strings.Split(str, "\n")
	fmt.Println(len(tmp))
	rue := make([]string, 0)
	for _, s := range tmp {
		r := []byte{48, 48, 48, 48, 48, 48, 48, 48}
		j := 0
		for i := 0; i < len(s); {
			if j >= 8 {
				rue = append(rue, string(r))
				j = 0
				r = []byte{48, 48, 48, 48, 48, 48, 48, 48}
			}
			r[j] = s[i]
			i++
			j++
		}
		if j <= 8 {
			rue = append(rue, string(r))
		}
	}
	fmt.Println(rue)
}
