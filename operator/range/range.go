package main

import "fmt"

// 对range的测试, 在range的使用过程中需要注意range是将当前的值copy了一份，因此在改的时候需要注意
// 如果需要在range中修改需要使用下标

type Animal struct {
	name string
	legs int
}

func main() {
	zoo := []Animal{
		Animal{"dog", 4},
		Animal{"panda", 4},
		Animal{"peacock", 2},
	}
	Read(zoo)
	fmt.Printf("read: %v\n", zoo)
	Write(zoo)
	fmt.Printf("wrote: %v\n", zoo)
}

func Read(zoo []Animal) {
	for _, a := range zoo {
		a.legs = 100
	}
	fmt.Printf("%v\n", zoo)
}
func Write(zoo []Animal) {
	for i := range zoo {
		zoo[i].legs = 100
	}
	fmt.Printf("%v ✅\n", zoo)
}
