package main

import "fmt"

func main() {
	t := maxArea2([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})
	fmt.Println(t)
}

func maxArea(height []int) int {
	tmp := 0
	for i, h := range height {
		if i+1 >= len(height) {
			break
		}
		for j, l := range height[i+1:] {
			tmp1 := 0
			if l <= h {
				tmp1 = l * (j + 1)
			} else {
				tmp1 = h * (j + 1)
			}
			if tmp1 > tmp {
				tmp = tmp1
			}
		}
	}
	return tmp
}
func maxArea2(height []int) int {
	tmp := 0
	for i, j := 0, len(height)-1; i != j; {
		tmp1 := 0
		if height[i] > height[j] {
			tmp1 = height[j] * (j - i)
			j--
		} else {
			tmp1 = height[i] * (j - i)
			i++
		}
		if tmp1 > tmp {
			tmp = tmp1
		}
	}
	return tmp
}
