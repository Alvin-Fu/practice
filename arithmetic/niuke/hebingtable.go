package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Srange struct {
	Key   int
	Value int
}

func main() {
	var nums = make([]int, 0)
	reader := bufio.NewScanner(os.Stdin)
	reader.Text()
	tmp := make(map[int]int)
	for reader.Scan() {
		sp := strings.Split(reader.Text(), " ")
		key, _ := strconv.Atoi(sp[0])
		value, _ := strconv.Atoi(sp[1])
		tmp[key] += value
	}
	rue := make([]Srange, len(tmp))
	i := 0
	for k, v := range tmp {
		rue[i] = Srange{
			Key:   k,
			Value: v,
		}
		i++
	}
	sort.Slice(rue, func(i, j int) bool {
		return rue[i].Key < rue[j].Key
	})
	for _, v := range rue {
		fmt.Println(v.Key, " ", v.Value)
	}
}

/*
题目描述
数据表记录包含表索引和数值（int范围的正整数），请对表索引相同的记录进行合并，即将相同索引的数值进行求和运算，输出按照key值升序进行输出。

输入描述:
先输入键值对的个数
然后输入成对的index和value值，以空格隔开

输出描述:
输出合并后的键值对（多行）
*/
