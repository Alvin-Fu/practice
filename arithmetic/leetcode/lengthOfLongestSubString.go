package main

import "fmt"

func main() {
	n := lengthOfLongestSubString("dvdf")
	fmt.Println(n)
}

func lengthOfLongestSubString(str string) int {
	n := 0
	datas := []byte(str)
	tmp := make([]byte, 0)
	//tmp = append(tmp, datas[0])
	for _, b := range datas {
		flag := false
		index := 0
		for i, t := range tmp {
			if b == t {
				flag = true
				index = i + 1
				break
			}
		}
		if !flag {
			tmp = append(tmp, b)
		} else {
			if n < len(tmp) {
				n = len(tmp)
			}
			tmp = tmp[index:]
			tmp = append(tmp, b)
		}

	}
	if n < len(tmp) {
		n = len(tmp)
	}
	return n
}

//Todo: 使用map去判断是否有重复的
