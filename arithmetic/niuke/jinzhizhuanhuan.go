package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

/*
题目描述
写出一个程序，接受一个十六进制的数，输出该数值的十进制表示。

输入描述:
输入一个十六进制的数值字符串。注意：一个用例会同时有多组输入数据，请参考帖子https://www.nowcoder.com/discuss/276处理多组输入的问题。

输出描述:
输出该数值的十进制字符串。不同组的测试用例用\n隔开。
示例1
输入
0xA
0xAA
输出
10
170
*/

func main() {
	str := "0x76E"
	fmt.Println(hexadecimalToDecimalism(str[2:]))

	reader := bufio.NewScanner(os.Stdin)
	rue := make([]int, 0)
	for reader.Scan() {
		s := reader.Text()
		s = s[2:]
		num := 0
		j := 0
		for i := len(s) - 1; i >= 0; i-- {

			num += byteToInt(s[i]) * int(math.Pow(16, float64(j)))
			j++
		}
		rue = append(rue, num)
	}
	for _, r := range rue {
		fmt.Printf("%d\n", r)
	}

}

func hexadecimalToDecimalism(tmp string) int {
	num := 0
	j := 0
	for i := len(tmp) - 1; i >= 0; i-- {
		fmt.Println(byteToInt(tmp[i]), j)
		num += byteToInt(tmp[i]) * int(math.Pow(16, float64(j)))
		j++
	}
	return num
}

func byteToInt(a byte) int {
	switch a {
	case 'A', 'a':
		return 10
	case 'B', 'b':
		return 11
	case 'C', 'c':
		return 12
	case 'D', 'd':
		return 13
	case 'E', 'e':
		return 14
	case 'F', 'f':
		return 15
	default:
		return int(a) - 48
	}
}
