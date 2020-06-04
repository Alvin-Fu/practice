package main

import "fmt"

func main(){
	tmp := make(map[int]string)
	tmp[1] = "hello"
	tmp[2] = "world"
	tmp[3] = "tom"
	fmt.Println(len(tmp))
	delete(tmp, 2)
	fmt.Println(len(tmp), tmp)
	seq := make(map[int]int, 3)
	for i := 0; i < 5; i ++{
		seq[i] = i
	}
	fmt.Println(seq)
}
