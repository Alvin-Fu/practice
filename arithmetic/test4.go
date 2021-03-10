package main

import (
	"bufio"
	"os"
	"strings"

	"fmt"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		str := scanner.Text()
		temp := strings.Split(str, ",")
		rue := ""
		for _, s := range temp {
			if len(s) == 0 {
				continue
			}
			data := make([]byte, 0)
			data = append(data, '/')
			i := 0
			for i < len(s) {
				if s[i] == '/' {
					i++
				} else {
					break
				}
			}
			j := len(s) - 1
			for j >= 0 {
				if s[j] == '/' {
					j--
				} else {
					break
				}
			}
			data = append(data, s[i:j]...)
			rue += string(data)
		}
		if len(rue) == 0 {
			fmt.Println("/")
		} else {
			fmt.Println(rue)
		}

	}
}
