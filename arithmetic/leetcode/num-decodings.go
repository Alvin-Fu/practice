package main

func main() {
	numDecodings("00019")
}
func numDecodings(s string) int {
	if len(s) <= 0 {
		return 0
	}
	var tmp string
	for i, t := range s {
		if t == '0' {
			continue
		} else {
			tmp = s[i:]
			break
		}
	}
	if len(tmp) == 0 {
		return 0
	}
	num := 1
	for _, t := range tmp {
		//for
	}
	return num
}
