package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var reader = bufio.NewReader(os.Stdin)
	for {
		var buff []byte
		var err error
		if buff, _, err = reader.ReadLine(); err != nil {
			return
		}
		var slices = strings.Split(string(buff), " ")
		var words = slices[1 : len(slices)-2]
		var word = slices[len(slices)-2]
		k, _ := strconv.Atoi(slices[len(slices)-1])
		fmt.Scanf("%s%d", &word, &k)
		count, str := findBrotherWord(words, word, k)
		fmt.Println(count)
		fmt.Println(str)
	}

}

func findBrotherWord(words []string, word string, k int) (int, string) {
	count := 0
	temp := make([]string, 0)
	temp2 := make(map[string]struct{}, 0)
	for _, w := range words {
		if len(w) != len(word) || w == word {
			continue
		}
		if _, ok := temp2[w]; ok {
			count++
			continue
		}
		var tmp [26]int
		isBrother := true
		for i := 0; i < len(w); i++ {
			tmp[word[i]-'a']++
			tmp[w[i]-'a']--
		}
		for i := 0; i < 26; i++ {
			if tmp[i] != 0 {
				isBrother = false
				break
			}
		}
		if isBrother {
			count++
			temp = append(temp, w)
			temp2[w] = struct{}{}
		}
	}
	if k <= len(temp) {
		sort.Strings(temp)
		return count, temp[k-1]
	}
	return count, ""
}

/*
题目描述
定义一个单词的“兄弟单词”为：交换该单词字母顺序，而不添加、删除、修改原有的字母就能生成的单词。
兄弟单词要求和原来的单词不同。例如：ab和ba是兄弟单词。ab和ab则不是兄弟单词。
现在给定你n个单词，另外再给你一个单词str，让你寻找str的兄弟单词里，字典序第k大的那个单词是什么？
注意：字典中可能有重复单词。本题含有多组输入数据。
输入描述:
先输入单词的个数n，再输入n个单词。
再输入一个单词，为待查找的单词x
最后输入数字k
输出描述:
输出查找到x的兄弟单词的个数m
然后输出查找到的按照字典顺序排序后的第k个兄弟单词，没有符合第k个的话则不用输出。
示例1
输入
3 abc bca cab abc 1
输出
2
bca
*/
