package main

import (
	"fmt"
)

var rue = make([]int, 0)

func main() {
	fmt.Println(findPid([]int{1, 3, 10, 5, 2, 6, 7}, []int{3, 0, 5, 3, 5, 2, 6}, 6))
}

func findPid(pid, ppid []int, id int) []int {
	if len(pid) != len(ppid) {
		return nil
	}
	temp := make(map[int][]int)
	for i := 0; i < len(ppid); i++ {
		_, ok := temp[ppid[i]]
		if !ok {
			temp[ppid[i]] = make([]int, 0)
		}
		temp[ppid[i]] = append(temp[ppid[i]], pid[i])
	}

	getRue(temp, id)
	return append(rue, id)
}

func getRue(temp map[int][]int, id int) {
	t, ok := temp[id]
	if !ok {
		return
	}
	rue = append(rue, t...)
	for _, i := range t {
		getRue(temp, i)
	}

}

func findPid2(pid, ppid []int, id int) []int {
	if len(pid) != len(ppid) {
		return nil
	}
	return nil
}
