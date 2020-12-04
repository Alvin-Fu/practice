package main

import "fmt"

func main() {
	fmt.Println(canPlaceFlowers([]int{1, 0, 0, 0, 0, 1}, 2))
}

func canPlaceFlowers(flowerbed []int, n int) bool {
	if len(flowerbed) == 1 && n == 1 && flowerbed[0] == 0 {
		return true
	}
	for i := 0; i < len(flowerbed)-1; i++ {
		if n == 0 {
			return true
		}
		if i == 0 {
			if flowerbed[i] == 0 && flowerbed[i+1] == 0 {
				flowerbed[i] = 1
				n--
			}
			continue
		}
		if i == len(flowerbed)-2 {
			if flowerbed[i+1] == 0 && flowerbed[i] == 0 {
				flowerbed[i+1] = 1
				n--
				continue
			}
		}
		if flowerbed[i] == 0 && flowerbed[i+1] == 0 && flowerbed[i-1] == 0 {
			n--
			flowerbed[i] = 1
			continue
		}
	}
	if n <= 0 {
		return true
	}
	return false
}
