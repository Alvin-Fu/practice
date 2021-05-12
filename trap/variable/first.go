package main

import "fmt"

func main(){
	fmt.Println(watShadowedDefer(50))
}
// 注意go中变量的作用域，以及参数传递时的值传递和指针传递的区别
func watShadowedDefer(i int)(ret int){
	ret = i * 2
	if ret > 10{
		ret := 10
		defer func() {
			ret = ret + 1
		}()
	}
	add(ret)
	return
}

func add(i int){
	i = i + 2
}