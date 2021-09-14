package bitmap

// BitMap, 存储int,提供判断是否存在的方法
type BitMap struct {
	words  []uint64
	length int
}

func NewBitMap() *BitMap {
	return &BitMap{}
}

func (b *BitMap) Add(num int) {
	index, bit := num/64, num%64
	for index >= len(b.words) {
		b.words = append(b.words, 0)
	}
	if b.words[index]&(1<<bit) == 0 { //不存在
		b.words[index] |= 1 << bit
		b.length++
	}
}

func (b *BitMap) Has(num int) bool {
	index, bit := num/64, num%64

	return index < len(b.words) && b.words[index]&(1<<bit) != 0
}

func (b *BitMap) Remove(num int) {
	index, bit := num/64, num%64
	if index >= len(b.words) {
		return
	}
	b.words[index] &= ^(1 << bit)
}
