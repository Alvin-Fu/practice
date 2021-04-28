package bytestring

import "unsafe"

func ByteStringPointer(b []byte) string{
	return *(*string)(unsafe.Pointer(&b))
}

func ByteString(b []byte)string{
	return string(b)
}