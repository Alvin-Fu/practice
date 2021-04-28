package main

import "fmt"

func main() {
	num := 3929
	nums := make([]int, 0)
	for num > 0 {
		nums = append(nums, num%10)
		num = num / 10
	}
	fmt.Println(nums)
	fmt.Println(intToRoman(500))
}

func intToRoman(num int) string {
	if num < 1 || num > 4000 {
		return ""
	}
	tmp := ""
	for i := 1000; i > 0; i = i / 10 {
		n := num / i
		if n == 0 {
			continue
		}
		num = num % i
		tmp += Roman(n, i)
	}

	return tmp
}

func Roman(n, index int) string {
	str := ""
	switch index {
	case 1000:
		for i := 0; i < n; i++ {
			str += "M"
		}
	case 100:
		if n < 4 {
			for i := 0; i < n; i++ {
				str += "C"
			}
		} else if n == 4 {
			str += "CD"
		} else if n == 9 {
			str += "CM"
		} else {
			str += "D"
			for i := 0; i < n-5; i++ {
				str += "C"
			}
		}
	case 10:
		if n < 4 {
			for i := 0; i < n; i++ {
				str += "X"
			}
		} else if n == 4 {
			str += "XL"
		} else if n == 9 {
			str += "XC"
		} else {
			str += "L"
			for i := 0; i < n-5; i++ {
				str += "X"
			}
		}
	default:
		if n < 4 {
			for i := 0; i < n; i++ {
				str += "I"
			}
		} else if n == 4 {
			str += "IV"
		} else if n == 9 {
			str += "IX"
		} else {
			str += "V"
			for i := 0; i < n-5; i++ {
				str += "I"
			}
		}
	}
	return str
}
