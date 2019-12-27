package copy_append

// 测试切片的append和copy的性能

func Copy(arr []int) []int {
	var rue = make([]int, len(arr))
	copy(rue, arr)
	rue[0] = 10
	return rue
}

func Append(arr []int) []int {
	var rue []int
	rue = append(rue, arr...)
	return rue
}
