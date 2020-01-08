package main

func main() {
	lengthOfLongestSubString("hello")
}

func lengthOfLongestSubString(str string) int {
	n := 0
	bytes := []byte(str)
	tmp := make([]byte, 0)
	tmp = append(tmp, bytes[0])
	for i, b := range bytes {
		flag := false
		for _, t := range tmp {
			if b == t {

			}
		}
	}
	return n
}
