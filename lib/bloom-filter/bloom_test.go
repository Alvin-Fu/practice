package bloom_filter

import (
	"strconv"
	"testing"
)

func TestBitMap_Add(t *testing.T) {
	bf := NewBloomFilter()
	for i := 0; i < 100000000; i++ {
		bf.Add(strconv.Itoa(i))
	}
	t.Log(bf.IsExit("1"))
}
