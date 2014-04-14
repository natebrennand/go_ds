package bloom

type BloomFilter struct {
	bits BitArray
	hashes int
}

func NewBloomFilter(size uint64, hashes int) BloomFilter {
	// x := BloomFilter{NewBigIntBitArray(size), hashes}
	x := BloomFilter{NewByteArray(size), hashes}
	return x
}

func (b BloomFilter) Add(elem interface{}) {
	hash1, err := ComputeHash(elem)
	if err != nil {
		panic("Hash fn failed")
	}
	b.bits.Set(hash1)

	hashN := hash1
	for i := 1; i < b.hashes; i++ {
		hashN, err := ComputeHash(hashN)
		if err != nil {
			panic("Hash fn failed")
		}
		b.bits.Set(hashN)
	}
}

func (b BloomFilter) Contains(elem interface{}) bool {
	hash1, err := ComputeHash(elem)
	if err != nil {
		panic("Hash fn failed")
	}
	if !b.bits.Get(hash1) {
		return false
	}

	for i := 1; i < b.hashes; i++ {
		hashN, err := ComputeHash(elem)
		if err != nil {
			panic("Hash fn failed")
		}
		if !b.bits.Get(hashN) {
			return false
		}
	}
	return true
}
