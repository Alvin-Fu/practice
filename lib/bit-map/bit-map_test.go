package bit_map

import "testing"

func TestBitMap_Add(t *testing.T) {
	bm := NewBitMap2(255 << 22)
	bm.Add(6, 1)
	bm.Add(6, bm.GetNum(6)+1)
	t.Log(bm.GetNum(6))
}
