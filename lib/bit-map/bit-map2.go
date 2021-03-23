package bit_map

// 2-bitmap,使用两位表示一个数字，00表示不存在，01表示存在，10表示多个，11未定义

type BitMap2 struct {
	bits []byte
	len  int
}

func NewBitMap2(max int) *BitMap2 {
	return &BitMap2{
		bits: make([]byte, (max>>2)+1),
		len:  max,
	}
}

func (b *BitMap2) Add(x, num uint) {
	index := x >> 2
	if index < 0 {
		index = 0
	}
	pos := x & 0x03
	// 将所在的位清零
	b.bits[index] &= ^(0x03 << (2 * pos))
	b.bits[index] |= uint8((num & 0x03) << (2 * pos))
}

func (b *BitMap2) GetNum(x uint) uint {
	index := x >> 2
	if index < 0 {
		index = 0
	}
	pos := x & 0x03
	// 先取出所在的位，在移位操作
	return uint(b.bits[index]&(0x03<<(2*pos))) >> (2 * pos)
}
