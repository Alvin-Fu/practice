package main

import "fmt"

func main() {
	fmt.Println(Multiply("999", "999"))

}
func Multiply(num1, num2 string) string {
	rue := ""
	if num1 == "" || num1 == "1" {
		return num2
	}
	if num2 == "" || num2 == "1" {
		return num1
	}
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	len1, len2 := len(num1), len(num2)
	tmp := make(map[int]int)
	bytes := make([]byte, len1+len2)
	//nums := make([]int, len1+len2)
	for i := len1; i > 0; i-- {
		for j := len2; j > 0; j-- {
			index := i + j - 1
			if index <= 0 {
				index = 0
			}
			num := byteToInt(num1[i-1]) * byteToInt(num2[j-1])
			tmp[index] += num
			//tmp[index-1] += num / 10
		}
	}
	fmt.Println(tmp)
	for i := len1 + len2 - 1; i > 0; i-- {
		tmp[i-1] += tmp[i] / 10
		tmp[i] = tmp[i] % 10
	}
	rue = string(bytes)
	fmt.Println(string(bytes[:]))
	fmt.Println(tmp)
	return rue
}

func byteToInt(b byte) int {
	switch b {
	case 48:
		return 0
	case 49:
		return 1
	case 50:
		return 2
	case 51:
		return 3
	case 52:
		return 4
	case 53:
		return 5
	case 54:
		return 6
	case 55:
		return 7
	case 56:
		return 8
	case 57:
		return 9
	default:
		return 0
	}
}

func getBitsTen(num int) (int, int) {
	return num / 10, num % 10
}
