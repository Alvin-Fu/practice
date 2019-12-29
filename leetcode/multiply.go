package main

import "fmt"

func main() {
	fmt.Println(Multiply("99", "99"))
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
	//bytes1, bytes2 := []byte(num1), []byte(num2)
	bytes := make([]byte, len1+len2)
	for i := 0; i < len1; i++ {
		for j := 0; j < len2; j++ {
			num := byteToInt(num1[i]) * byteToInt(num2[j])
			n, m := getBitsTen(num)
			if m != 0 {
				t := byte(byteToInt(bytes[i+j]) + m)
				n1, m1 := getBitsTen(t)

			}
			bytes[i+j+1] = byte(byteToInt(bytes[i+j+1]) + m)
		}
	}
	fmt.Println(bytes)
	rue = string(bytes)
	fmt.Println(string(bytes[:]))
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
