package main

import "fmt"

func main() {
	str := convert("LEAD", 2)
	fmt.Println(str)
}

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	str := make([]byte, 0)
	data := []byte(s)
	first := (numRows - 1) * 2
	for j := 0; j < numRows; j++ {
		count := 1
		m := j + 1
		//fmt.Println(m)
		t := (numRows - m) * 2
		for i := j; i < len(s); {
			str = append(str, data[i])
			//fmt.Println(string(str), i, t)
			if t == 0 {
				i += first
			} else if count%2 == 0 && first-t != 0 {
				i += first - t
			} else {
				i += t
			}
			count++
		}
	}
	return string(str)
}
