package bloom

import (
	"encoding/binary"
)

type BloomFilter struct {
	bits   BitArray
	hashes int
	hash   HashFn
}

// Create a new Bloom Filter
// size is the power of 2 for the size you would like, i.e. 10 means 1024
func NewJenkinsBloomFilter(size int, hashes int) BloomFilter {
	// + 3 makes it 8 bits per expected element
	x := BloomFilter{NewBigIntBitArray(size + 3), hashes, jenkinsHash{}}
	return x
}

// Adds a new element to the bloom filter
func (b BloomFilter) Add(elem []byte) {
	hash1, err := b.hash.ComputeHash(elem)
	if err != nil {
		panic("Hash fn failed")
	}
	b.bits.Set(hash1)

	hashN := hash1
	hashKey := make([]byte, 10)
	for i := 1; i < b.hashes; i++ {
		binary.PutUvarint(hashKey, hashN)
		hashN, err := b.hash.ComputeHash(hashKey)
		if err != nil {
			panic("Hash fn failed")
		}
		b.bits.Set(hashN)
	}
}

// Test if the Bloom Filter contains the element
func (b BloomFilter) Contains(elem []byte) bool {
	hash1, err := b.hash.ComputeHash(elem)
	if err != nil {
		panic("Hash fn failed")
	}
	if !b.bits.Get(hash1) {
		return false
	}

	hashN := hash1
	hashKey := make([]byte, 10)
	for i := 1; i < b.hashes; i++ {
		binary.PutUvarint(hashKey, hashN)
		hashN, err := b.hash.ComputeHash(hashKey)
		if err != nil {
			panic("Hash fn failed")
		}
		if !b.bits.Get(hashN) {
			return false
		}
	}
	return true
}

// Print representation of the current state
func (b BloomFilter) Print() {
	b.bits.Print()
}
