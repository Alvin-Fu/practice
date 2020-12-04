package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

/*
给你一个字符串 s 和一个整数数组 cost ，其中 cost[i] 是从 s 中删除字符 i 的代价。

返回使字符串任意相邻两个字母不相同的最小删除成本。

请注意，删除一个字符后，删除其他字符的成本不会改变。



示例 1：

输入：s = "abaac", cost = [1,2,3,4,5]
输出：3
解释：删除字母 "a" 的成本为 3，然后得到 "abac"（字符串中相邻两个字母不相同）。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/minimum-deletion-cost-to-avoid-repeating-letters
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

//"aaabbbabbbb"
//[3,5,10,7,5,3,5,5,4,8,1]

//"cddcdcae"
//[4,8,8,4,4,5,4,2]
func main() {
	fmt.Println(minCost("cddcdcae", []int{4, 8, 8, 4, 4, 5, 4, 2}))
	fmt.Println(minCost2("cddcdcae", []int{4, 8, 8, 4, 4, 5, 4, 2}))
}

func minCost(s string, cost []int) int {
	data := StringToBytesHeader(s)
	min := 0
	for i := 0; i < len(data)-1; i++ {
		if data[i] == data[i+1] {
			index := i + 1
			for ; index < len(data)-1; index++ {
				if data[index] != data[index+1] {
					break
				}
			}
			max := cost[i]
			for x := i; x <= index; x++ {
				if cost[x] > max {
					max = cost[x]
				}
				min += cost[x]
			}
			min -= max
			i = index
		}
	}
	return min
}

func minCost2(s string, cost []int) int {
	data := StringToBytesHeader(s)
	min := 0
	for i := 0; i < len(data); {
		ch := data[i]
		max := 0
		for i < len(data) && data[i] == ch {
			if max < cost[i] {
				max = cost[i]
			}
			min += cost[i]
			i++
		}
		min -= max
	}
	return min
}

func StringToBytesHeader(str string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&str))
	bh := reflect.SliceHeader{
		Data: sh.Data,
		Len:  sh.Len,
		Cap:  sh.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&bh))
}
