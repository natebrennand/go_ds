package bloom

type BloomFilter struct {
	bits   BitArray
	hashes int
	hash   HashFn
}

func NewBloomFilter(size uint64, hashes int) BloomFilter {
	x := BloomFilter{NewBigIntBitArray(size), hashes, jenkinsHash{}}
	return x
}

func NewJenkinsBloomFilter(size uint64, hashes int) BloomFilter {
	x := BloomFilter{NewBigIntBitArray(size), hashes, jenkinsHash{}}
	return x
}

func (b BloomFilter) Add(elem interface{}) {
	hash1, err := b.hash.ComputeHash(elem)
	if err != nil {
		panic("Hash fn failed")
	}
	b.bits.Set(hash1)

	hashN := hash1
	for i := 1; i < b.hashes; i++ {
		hashN, err := b.hash.ComputeHash(hashN)
		if err != nil {
			panic("Hash fn failed")
		}
		b.bits.Set(hashN)
	}
}

func (b BloomFilter) Contains(elem interface{}) bool {
	hash1, err := b.hash.ComputeHash(elem)
	if err != nil {
		panic("Hash fn failed")
	}
	if !b.bits.Get(hash1) {
		return false
	}

	hashN := hash1
	for i := 1; i < b.hashes; i++ {
		hashN, err := b.hash.ComputeHash(hashN)
		if err != nil {
			panic("Hash fn failed")
		}
		if !b.bits.Get(hashN) {
			return false
		}
	}
	return true
}
