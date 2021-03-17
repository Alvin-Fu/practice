package main

import (
	"fmt"
	"sort"
)

type Fruits struct {
	chip   int
	profit int
	isUse  bool
}

func main() {
	//scanner := bufio.NewScanner(os.Stdin)
	//for scanner.Scan() {
	//	tmp := strings.Split(scanner.Text(), " ")
	//	fruits := make([]Fruits, len(tmp))
	//	for i, t := range tmp {
	//		n, _ := strconv.Atoi(t)
	//		fruits[i] = Fruits{
	//			chip:  n,
	//			isUse: false,
	//		}
	//	}
	//	for i, t := range strings.Split(scanner.Text(), " ") {
	//		n, _ := strconv.Atoi(t)
	//		fruits[i].profit = n - fruits[i].chip
	//	}
	//	capital, _ := strconv.Atoi(scanner.Text())
	//	fmt.Println(findMaxProfit(fruits, capital))
	//}
	fruits := make([]Fruits, 4)
	fruits[0] = Fruits{
		chip: 4, isUse: false, profit: 1,
	}
	fruits[1] = Fruits{
		chip: 2, isUse: false, profit: 1,
	}
	fruits[2] = Fruits{
		chip: 9, isUse: false, profit: 3,
	}
	fruits[3] = Fruits{
		chip: 20, isUse: false, profit: 9,
	}
	fmt.Println(findMaxProfit(fruits, 15))
}

func findMaxProfit(fruits []Fruits, capital int) int {
	sort.Slice(fruits, func(i, j int) bool {
		return fruits[i].profit > fruits[j].profit
	})
	temp := capital
	max := 0
	useCount := 0
	for isOver(fruits, temp) {
		for i := 0; i < len(fruits); i++ {
			if !fruits[i].isUse && temp-fruits[i].chip >= 0 {
				max += fruits[i].profit
				fruits[i].isUse = true
				temp -= fruits[i].chip
				useCount++
			}
		}
		temp = capital + max
	}
	return max + capital
}

func isOver(fruits []Fruits, capital int) bool {
	for i := 0; i < len(fruits); i++ {
		if fruits[i].isUse {
			continue
		}
		if fruits[i].chip <= capital {
			return true
		}
	}
	return false
}
