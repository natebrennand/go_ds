package bloom

import (
	"hash/fnv"
)

type BloomFilter struct {
	bits BigIntBitArray
}

func NewBloomFilter(size uint64) BloomFilter {
	return BloomFilter{NewBigIntBitArray(int(1 + 10*size/2^32))}
}

func (b BloomFilter) Add (elem interface{}) {
	b.bits.Set(fnv.New64().Sum64())
	b.bits.Set(fnv.New64a().Sum64())
}

func (b BloomFilter) Contains (elem interface{}) bool {
	return b.bits.Get(fnv.New64().Sum64()) && b.bits.Get(fnv.New64a().Sum64())
}
