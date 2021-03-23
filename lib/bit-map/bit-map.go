package bit_map

type BitMap struct {
	bits []byte
	len  int
}

func NewBitMap(max int) *BitMap {
	return &BitMap{
		bits: make([]byte, (max>>3)+1),
		len:  max,
	}
}

func (b *BitMap) Add(num uint) {
	index := num >> 3
	if index < 0 {
		index = 0
	}
	pos := num & 0x07
	b.bits[index] |= 1 << pos
}

func (b *BitMap) Delete(num uint) {
	index := num >> 3
	if index < 0 {
		index = 0
	}
	pos := num & 0x07
	b.bits[index] = b.bits[index] &^ (1 << pos)
}

func (b *BitMap) IsExit(num uint) bool {
	index := num >> 3
	if index < 0 {
		index = 0
	}
	pos := num & 0x07
	return b.bits[index]&1<<pos != 0
}
