package bloom_filter

type BloomFilter struct {
	seed     []int
	bitMap   *BitMap
	hashFunc *HashFunc
}

type HashFunc struct {
	size int
}

func NewBloomFilter() *BloomFilter {
	return &BloomFilter{
		seed:     []int{3, 5, 7 /*, 11, 13, 31, 37, 61*/},
		bitMap:   NewBitMap(256 << 22),
		hashFunc: NewHashFunc(256 << 22),
	}
}
func NewHashFunc(size int) *HashFunc {
	return &HashFunc{
		size: size,
	}
}

func (h *HashFunc) Hash(data string, seed int) uint {
	rue := 0
	for _, d := range data {
		rue = rue*seed + int(d)
	}
	return uint((h.size - 1) & rue)
}

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

func (bf *BloomFilter) Add(value string) {
	if len(value) > 0 {
		for _, s := range bf.seed {
			bf.bitMap.Add(bf.hashFunc.Hash(value, s))
		}
	}
}

func (bf *BloomFilter) IsExit(value string) bool {
	if len(value) <= 0 {
		return false
	}
	for _, s := range bf.seed {
		if !bf.bitMap.IsExit(bf.hashFunc.Hash(value, s)) {
			return false
		}
	}
	return true
}
