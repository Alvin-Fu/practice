package bytestring

import (
	"reflect"
	"unsafe"
)

func StringToBytesHeader(str string)[]byte{
	sh := (*reflect.StringHeader)(unsafe.Pointer(&str))
	bh := reflect.SliceHeader{
		Data: sh.Data,
		Len: sh.Len,
		Cap: sh.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&bh))
}

func StringToBytes(str string)[]byte{
	return []byte(str)
}