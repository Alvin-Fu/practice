package main

var letterCombination = make(map[byte][]byte)

func init() {
	letterCombination['2'] = []byte{'a', 'b', 'c'}
	letterCombination['3'] = []byte{'d', 'e', 'f'}
	letterCombination['4'] = []byte{'g', 'h', 'i'}
	letterCombination['5'] = []byte{'j', 'k', 'l'}
	letterCombination['6'] = []byte{'m', 'n', 'o'}
	letterCombination['7'] = []byte{'p', 'q', 'r', 's'}
	letterCombination['8'] = []byte{'t', 'u', 'v'}
	letterCombination['9'] = []byte{'w', 'x', 'y', 'z'}
}

func main() {
}

func letterCombinations(digits string) []string {
	if len(digits) <= 0 {
		return nil
	}
	rue := make([][]byte, 0)
	for i, d := range []byte(digits) {
		for j, c := range letterCombination[d] {

		}
	}
	return nil
}
