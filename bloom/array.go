package bloom

import (
	"math/big"
)

type BitArray interface {
	Get() bool
	Set()
}

type BigIntBitArray []*big.Int

func NewBigIntBitArray (size int) BigIntBitArray{
	arr := make(BigIntBitArray, size)
	for index := range arr {
		arr[index] = big.NewInt(0)
	}
	return arr
}

func (b BigIntBitArray) Get(index uint64) bool {
	intNum := int(index / uint64(len(b))) % len(b)
	offset := int(index % uint64(len(b)))
	if 1 == b[intNum].Bit(offset) {
		return true
	}
	return false
}

func (b BigIntBitArray) Set(index uint64) {
	intNum := int(index / uint64(len(b))) % len(b)
	offset := int(index % uint64(len(b))) % 2^64
	b[intNum].SetBit(b[intNum], offset, 1)
}
