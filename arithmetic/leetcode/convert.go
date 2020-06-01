package main

import "fmt"

func main() {
	str := convert("LEETCODEISHIRING", 3)
	fmt.Println(str)
}

func convert(s string, numRows int) string {
	str := make([]byte, 0)
	data := []byte(s)
	for j := 0; j < numRows; {

	}
	for i := 0; i < len(s); i += (numRows - 1) * 2 {
		str = append(str, data[i])
	}
	for i := 1; i < len(s); i += (numRows - 2) * 2 {
		str = append(str, data[i])
	}
	for i := 2; i < len(s); i += (numRows - 1) * 2 {
		str = append(str, data[i])
	}
	return string(str)
}
