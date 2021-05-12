package main

import (
		"strings"
	"strconv"
		"unsafe"
	"reflect"
	"fmt"
)

func main() {
	num := "06:30"
	//fmt.Println(strconv.Atoi(num))
	//t := time.Now()
	//fmt.Println(t.Hour() * 100 + t.Minute())
	//num = string([]byte(num)[:2])
	//fmt.Println(num)
	fmt.Println(StringToInt(num))

	//str := "hello"
	////n := LastIndex(str, "/")
	////str = string([]byte(str)[:n])
	////fmt.Println(str)
	////strings.Repeat()
	//fmt.Println(utf8.DecodeLastRuneInString(str))
	//fmt.Println(strings.Replace(str, "l", "23", 3))
	//
	//fmt.Println("Hello, 世界")
	//var s string
	//s = "333,"
	//fmt.Println(strings.TrimRight(str, "/"))
	//fmt.Println(s)
	//s = strings.TrimRight(s, ",")
	//fmt.Println(s)
}

func LastIndex(str, sep string) int {
	return strings.LastIndex(str, sep)
}
func StringToInt(str string) (int, error) {
	if len(str) > 5 {
		data := StringToBytesHeader(str)
		str = ByteStringPointer(data[:5])
	}
	str = strings.Replace(str, ":", "", -1)
	if len(str) != 4 {

	}
	num, err := strconv.Atoi(str)
	if err != nil {
		return 0, err
	}
	return num, nil
}

func ByteStringPointer(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func StringToBytesHeader(str string)[]byte{
	sh := (*reflect.StringHeader)(unsafe.Pointer(&str))
	bh := reflect.SliceHeader{
		Data: sh.Data,
		Len: sh.Len,
		Cap: sh.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&bh))
}