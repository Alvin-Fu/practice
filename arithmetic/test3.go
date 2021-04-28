package main

import "fmt"

func main() {
	fmt.Println(f([]byte("hello   sdjk sdjls")))
	fmt.Println(test("mn", "abmny"))
}

func f(arr []byte) string {
	index := 0
	count := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] == ' ' {
			if arr[index] != ' ' {
				index = i
			}
			count++
		} else {
			if arr[index] == ' ' {
				m := 0
				for j := i; j < len(arr); j++ {
					if arr[j] == ' ' {
						m = j
						index = j
						break
					} else {
						arr[index] = arr[j]
						index++
					}
				}
				i = m
			}
		}
	}
	arr = arr[:len(arr)-count]
	return string(arr)
}

func test(s1, s2 string) bool {
	if len(s1) > len(s2) || len(s1) <= 0 || len(s2) <= 0 {
		return false
	}
	index := 0
	isAdd := true
	i, j := 0, 0
	for j < len(s2) && j >= 0 {
		if s1[i] == s2[j] {
			i++
			index = j
		} else if s1[i] != s2[j] && i > 0 {
			if j > 0 {
				isAdd = false
				j = index
			} else {
				i = 0
				j++
				isAdd = true
				continue
			}
		}
		if i == len(s1) {
			return true
		}
		if !isAdd {
			j--
		} else {
			j++
		}
	}
	return false
}

func te(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
