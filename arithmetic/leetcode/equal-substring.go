package main

import (
	"fmt"
)

/*
给你两个长度相同的字符串，s 和 t。

将 s 中的第 i 个字符变到 t 中的第 i 个字符需要 |s[i] - t[i]| 的开销（开销可能为 0），也就是两个字符的 ASCII 码值的差的绝对值。

用于变更字符串的最大预算是 maxCost。在转化字符串时，总开销应当小于等于该预算，这也意味着字符串的转化可能是不完全的。

如果你可以将 s 的子字符串转化为它在 t 中对应的子字符串，则返回可以转化的最大长度。

如果 s 中没有子字符串可以转化成 t 中对应的子字符串，则返回 0。

输入：s = "abcd", t = "bcdf", cost = 3
输出：3
解释：s 中的 "abc" 可以变为 "bcd"。开销为 3，所以最大长度为 3。

输入：s = "abcd", t = "acde", cost = 0
输出：1
解释：你无法作出任何改动，所以最大长度为 1。
"krrgw"
"zjxss"
19
*/

func main() {
	fmt.Println(equalSubstring("abcd", "cdef", 1))
}

func equalSubstring(s string, t string, maxCost int) int {
	bs := []byte(s)
	bt := []byte(t)
	r := 0
	tmp := make([]int, len(s))
	maxLen := 0
	for r < len(s) {
		n := absolute(int(bs[r]) - int(bt[r]))
		tmp[r] = n
		r++
	}
	j := 0
	nums := 0
	for i, t := range tmp {
		nums += t
		for nums > maxCost {
			nums -= tmp[j]
			j++
		}
		if maxLen < i-j+1 {
			maxLen = i - j + 1
		}
	}

	return maxLen
}

func absolute(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
