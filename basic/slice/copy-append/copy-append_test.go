package copy_append

import "testing"

var array = make([]int, 0)

func init() {
	for i := 0; i < 1000; i++ {
		array = append(array, i)
	}
}

func BenchmarkCopy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tmp := Copy(array)
		if len(tmp) != len(array) {
			b.Fatalf("err copy: %d", len(tmp))
		}
	}
}

func BenchmarkAppend(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tmp := Append(array)
		if len(tmp) != len(array) {
			b.Fatalf("err append")
		}
	}
}
