package bloom

import (
	"fmt"
	"math/big"
)

type BitArray interface {
	Get(uint64) bool
	Set(uint64)
	Print() string
}

type BigIntBitArray []*big.Int

const (
	BITS_PER_INT         = 64
	BITS_PER_INT_POWER_2 = 6
)

func (b BigIntBitArray) getIndexes(index uint64) (int, int) {
	bucket := int(index>>BITS_PER_INT_POWER_2) % len(b)
	offset := int(index % BITS_PER_INT)
	return bucket, offset
}

// Size is the power of 2 for the size you would like, i.e. 10 means 1024
func NewBigIntBitArray(size int) BigIntBitArray {
	// guarantee at least 1 bucket
	pow2 := uint(size - BITS_PER_INT_POWER_2)
	if size-BITS_PER_INT_POWER_2 < 0 {
		pow2 = uint(0)
	}

	// shift for # of bits
	numBuckets := (1 << uint(pow2))
	arr := make(BigIntBitArray, numBuckets)
	for index := range arr {
		arr[index] = big.NewInt(0)
	}
	return arr
}

// returns the true if the bit is turned on
func (b BigIntBitArray) Get(index uint64) bool {
	bucket, offset := b.getIndexes(index)
	return 1 == b[bucket].Bit(offset)
}

// sets the bit to 1
func (b BigIntBitArray) Set(index uint64) {
	bucket, offset := b.getIndexes(index)
	b[bucket].SetBit(b[bucket], offset, 1)
}

// print representation of the bit array, 64bit lines
func (b BigIntBitArray) Print() string {
	printString := ""
	for i, block := range b {
		printString += fmt.Sprintf("%03d: %064b\n", i, block)
	}
	return printString
}
