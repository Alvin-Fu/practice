package bytestring

import (
	"testing"
	"fmt"
)

var data = []byte{1, 2, 3, 4, 4, 4, 4,4 ,4, 1, 2, 3, 4, 4, 4, 4,4 ,4, 1,
2, 3, 4, 4, 4, 4,4 ,4, 1, 2, 3, 4, 4, 4, 4,4 ,4,1, 2, 3, 4, 4, 4, 4,4 ,4,
1, 2, 3, 4, 4, 4, 4,4 ,4,1, 2, 3, 4, 4, 4, 4,4 ,4,1, 2, 3, 4, 4, 4, 4,4 ,
4,1, 2, 3, 4, 4, 4, 4,4 ,4,1, 2, 3, 4, 4, 4, 4,4 ,4,1, 2, 3, 4, 4, 4, 4,4 ,4}

var str = "jsdklsjldksfjlsajklsjdklsjklsdajdkas"
func BenchmarkByteStringPointer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ByteStringPointer(data)
	}
}

func BenchmarkByteString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ByteString(data)
	}
}

func BenchmarkStringToBytes(b *testing.B) {
	for i := 0; i < b.N; i++{
		StringToBytes(str)
	}
}

func BenchmarkStringToBytesHeader(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StringToBytesHeader(str)
	}
}

func TestByteStringPointer(t *testing.T) {
	d := StringToBytesHeader(str)
	d = append(d, 0)
	fmt.Println(d)
}