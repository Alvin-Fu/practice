package main

import (
	"fmt"
)

func main() {
	//var l, s, t, m int
	//fmt.Scan(&l, &s, &t, &m)
	//stone := make([]int, m)
	//for i := 0; i < m; i++ {
	//	fmt.Scan(&stone[i])
	//}

	findWay(500, 4, 5, []int{2, 3, 5, 6, 7, 8, 9, 10, 30, 100})
}

func findWay(lenth, stepL, stepR int, stone []int) int {
	ans := 0
	if stepL == stepR {
		for _, s := range stone {
			if s%stepR == 0 {
				ans++
			}
		}
		return ans
	}
	ans = 100
	step := stepL * stepR
	ed := 0
	deep := make([]int, lenth)
	for i := 1; i < len(stone); i++ {
		if stone[i]-stone[i-1] >= step {
			ed += step
		} else {
			ed = ed + stone[i] - stone[i-1]
		}
		deep[ed] = 1
	}
	deep2 := make([]int, lenth)
	for i := ed; i >= 0; i-- {
		deep2[i] = len(stone)
		for j := stepL; j <= stepR; j++ {
			deep2[i] = min(deep2[i], deep[i]+deep2[i+j])
		}
	}
	for i := stepL; i <= step; i++ {
		ans = min(ans, deep2[i])
	}
	fmt.Println(ans)
	return ans
}

//func findWay2(lenth, stepL, stepR int, stone []int) int {
//	ans := 0
//	if stepL == stepR {
//		for _, s := range stone {
//			if s%stepR == 0 {
//				ans++
//			}
//		}
//		return ans
//	}
//	sort.Slice(stone, func(i, j int) bool {
//		return stone[i] < stone[j]
//	})
//	que := make([]int, 0)
//	for i := 0; i < stepR; i ++{
//	    que = append(que, 0)
//    }
//	stepTotal, stepStone := 0, -1
//	for {
//	    if stepTotal - stepR >=
//    }
//}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
