package main

import "fmt"

type article struct {
	price   int
	weight  int
	id, wp  int
	p1, wp1 int
	p2, wp2 int
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func main() {
	for {
		var maxPrice, maxNum int
		fmt.Scanf("%d %d", &maxPrice, &maxNum)

		if maxPrice <= 0 || maxPrice%10 != 0 || maxNum <= 0 {
			return
		}

		maxPrice /= 10

		goods := make([]article, maxNum+1)

		for i := 1; i <= maxNum; i++ {
			var w, p, q int
			fmt.Scanf("%d %d %d", &p, &w, &q)
			if w <= 0 || p <= 0 || p%10 != 0 {
				return
			}
			p /= 10
			goods[i] = article{
				price:  p,
				weight: w,
				id:     q,
				wp:     p * w,
			}
			if q > 0 {
				if goods[q].p1 == 0 {
					goods[q].p1 = goods[i].price
					goods[q].wp1 = goods[i].wp
				} else if goods[q].p2 == 0 {
					goods[q].p2 = goods[i].price
					goods[q].wp2 = goods[i].wp
				} else {
					fmt.Println("wrong input data")
					return
				}
			}
		}

		totalValue := make([]int, maxPrice+1)

		for i := 1; i <= maxNum; i++ {
			item := goods[i]
			if item.id != 0 {
				continue
			}
			fmt.Println(item, totalValue)
			for j := maxPrice; j >= item.price; j-- {
				fmt.Println(j, totalValue)
				totalValue[j] = max(totalValue[j], totalValue[j-item.price]+item.wp)

				if item.p1 > 0 && j >= (item.price+item.p1) {
					totalValue[j] = max(totalValue[j], totalValue[j-item.price-item.p1]+item.wp+item.wp1)
				}

				if item.p2 > 0 && j >= (item.price+item.p2) {
					totalValue[j] = max(totalValue[j], totalValue[j-item.price-item.p2]+item.wp+item.wp2)
				}

				if item.p1 > 0 && item.p2 > 0 && j >= (item.price+item.p1+item.p2) {
					totalValue[j] = max(totalValue[j], totalValue[j-item.price-item.p1-item.p2]+item.wp+item.wp1+item.wp2)
				}
			}
		}

		fmt.Printf("%d\n", totalValue[maxPrice]*10)
	}
}
