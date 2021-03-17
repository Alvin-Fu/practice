package main

import "fmt"

func main() {
	str := ""
	fmt.Scanf("%s", &str)
	fmt.Println(findMaxNumber(str))
}

func findMaxNumber(str string) string {
	var max []byte
	data := []byte(str)
	for i := 0; i < len(data); i++ {
		if data[i] >= '0' && data[i] <= '9' {
			j := i
			for ; j < len(data); j++ {
				if data[j] > '9' || data[j] < '0' {
					break
				}
			}
			if len(max) < len(data[i:j]) {
				max = data[i:j]
			}
			i = j
		}
	}
	return string(max)
}
