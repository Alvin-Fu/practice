package main

import (
	"fmt"
)

var N, maxs = 0, [2]int{0, 0}

func main() {
	for i := 0; i < 2; i++ {
		fmt.Scan(&N)
		maxs[i] = loopInput(N)
	}
	for _, v := range maxs {
		fmt.Println(v)
	}
}

func loopInput(N int) int {
	hights, inhights := make([]int, N), make([]int, N)
	asc, des := make([]int, N), make([]int, N)
	max := 0
	for i := 0; i < N; i++ {
		fmt.Scan(&hights[i])
	}
	asc = getMax(N, hights)
	for i := range hights {
		inhights[i] = hights[N-1-i]
	}
	des = getMax(N, inhights)
	for i, v := range asc {
		if max < v+des[N-1-i]-1 {
			max = v + des[N-1-i] - 1
		}
	}
	return N - max
}

func getMax(N int, hights []int) []int {
	ace, ends := make([]int, N), make([]int, N)
	right, l, r, m := 0, 0, 0, 0
	ends[0], ace[0] = hights[0], 1
	for i := 1; i < len(hights); i++ {
		l = 0
		r = right
		for l <= r {
			m = (l + r) / 2
			if hights[i] > ends[m] {
				l = m + 1
			} else {
				r = m - 1
			}
		}
		if right < l {
			right = l
		}
		ends[l] = hights[i]
		ace[i] = l + 1
	}
	return ace
}

/*
题目描述
计算最少出列多少位同学，使得剩下的同学排成合唱队形
说明：
N位同学站成一排，音乐老师要请其中的(N-K)位同学出列，使得剩下的K位同学排成合唱队形。
合唱队形是指这样的一种队形：设K位同学从左到右依次编号为1，2…，K，
他们的身高分别为T1，T2，…，TK，
则他们的身高满足存在i（1<=i<=K）使得T1<T2<......<Ti-1<Ti>Ti+1>......>TK。

你的任务是，已知所有N位同学的身高，计算最少需要几位同学出列，可以使得剩下的同学排成合唱队形。

注意不允许改变队列元素的先后顺序
请注意处理多组输入输出！

输入描述:
整数N

输出描述:
最少需要几位同学出列

示例1
输入
8
186 186 150 200 160 130 197 200
输出
4
*/
