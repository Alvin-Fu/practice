package probability

import "math/rand"

func rand7() int {
	n := rand.Intn(7)
	if n == 0 {
		return 7
	}
	return n
}

// 独立事件，取两次rand7的组合有49种，由于是独立事件，每一种的可能都是1/49
// 详解https://blog.csdn.net/shandianling/article/details/8091449
func rand10() int {
	var a int
	a = (rand7()-1)*7 + rand7()
	for a > 40 {
		a = (rand7()-1)*7 + rand7()
	}
	return a%10 + 1
}
