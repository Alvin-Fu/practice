package main

import (
	"fmt"
	"strings"
)

func main() {
	str := ""
	fmt.Scanf("%s", &str)
	tmp := strings.Split(str, ";")
	temp := make(map[byte]int, 4)
	for _, t := range tmp {
		if len(t) < 2 || len(t) > 3 {
			continue
		}
		if !(t[0] == 'A' || t[0] == 'S' || t[0] == 'W' || t[0] == 'D') {
			continue
		}
		if t[1] < 49 || t[1] > 57 {
			continue
		}
		if len(t[1:]) == 1 {
			temp[t[0]] += int(t[1] - 48)
		} else if len(t[1:]) == 2 {
			if t[2] < 48 || t[2] > 57 {
				continue
			}
			temp[t[0]] += int(t[1]-48)*10 + int(t[2]-48)
		}
	}
	a := temp['A']
	d := temp['D']
	w := temp['W']
	s := temp['S']
	fmt.Printf("%d,%d\n", d-a, w-s)
}

/*
题目描述
开发一个坐标计算工具， A表示向左移动，D表示向右移动，W表示向上移动，S表示向下移动。从（0,0）点开始移动，
从输入字符串里面读取一些坐标，并将最终输入结果输出到输出文件里面。

输入：

合法坐标为A(或者D或者W或者S) + 数字（两位以内）

坐标之间以;分隔。

非法坐标点需要进行丢弃。如AA10;  A1A;  $%$;  YAD; 等。

下面是一个简单的例子 如：

A10;S20;W10;D30;X;A1A;B10A11;;A10;

处理过程：

起点（0,0）

+   A10   =  （-10,0）

+   S20   =  (-10,-20)

+   W10  =  (-10,-10)

+   D30  =  (20,-10)

+   x    =  无效

+   A1A   =  无效

+   B10A11   =  无效

+  一个空 不影响

+   A10  =  (10,-10)

结果 （10， -10）

注意请处理多组输入输出

输入描述:
一行字符串

输出描述:
最终坐标，以逗号分隔

示例1
输入
A10;S20;W10;D30;X;A1A;B10A11;;A10;
输出
10,-10
*/
