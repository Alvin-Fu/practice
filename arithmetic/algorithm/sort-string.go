package main

import (
	"fmt"
	"strings"
)

func main() {
	//scanner := bufio.NewScanner(os.Stdin)
	//for scanner.Scan() {
	//	text := scanner.Text()
	//	temp := strings.Split(text, " ")
	//}

	sortString(strings.Split("I am an 20-years out–standing @*** * -stu- dent", " "))
}

func sortString(temp []string) []string {
	rue := make([]string, 0)
	isFirst := false
	for j := len(temp) - 1; j >= 0; j-- {
		t := temp[j]
		if strings.Count(t, "-") == 1 && !isFirst {
			isFirst = true
			tmp := make([]byte, 0)
			for i := 0; i < len(t); i++ {
				if isSave(t[i]) || t[i] == '-' {
					tmp = append(tmp, t[i])
				} else {
					if len(tmp) > 0 {
						rue = append(rue, string(tmp))
						tmp = make([]byte, 0)
					}
				}
			}
			if len(tmp) > 0 {
				rue = append(rue, string(tmp))
				tmp = make([]byte, 0)
			}
		} else {
			tmp := make([]byte, 0)
			for i := 0; i < len(t); i++ {
				if isSave(t[i]) {
					tmp = append(tmp, t[i])
				} else {
					if len(tmp) > 0 {
						rue = append(rue, string(tmp))
						tmp = make([]byte, 0)
					}
				}
			}
			if len(tmp) > 0 {
				rue = append(rue, string(tmp))
				tmp = make([]byte, 0)
			}
		}
	}
	fmt.Println(rue)
	return rue
}

func isSave(b byte) bool {
	if (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z') || (b >= '0' && b <= '9') {
		return true
	}
	return false
}

/*
题目
对数字，字符，数字串，字符串，以及数字与字符串组合进行倒序排列
字符范围：由a到z，A到Z，数字范围：由0到9
符号”的定义
（1）”做为连接符使用时作为字符串的一部分，例如“20-years”作为一个整体字符串呈现；
（2）连续出现2个’及以上时视为字符串间隔符，如“out–standing"中的“-“视为间隔符，是2个独立整体字符串"out"和"standing"；
除了1，2里面定义的字符以外其他的所有字符，都是非法字符，作为字符串的间隔符处理，倒序后间隔符作为空格处理；
要求倒排后的单词间隔符以一个空格表示；如果有多个间隔符时，倒排转换后也只介许出现一个字格间隔符：
输入示例
I am an 20-years out–standing @ * -stu- dent

输出示例
dent stu standing out 20-years an am I
*/
