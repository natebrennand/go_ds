package bloom

import (
	"math/big"
)

type BitArray interface {
	Get(uint64) bool
	Set(uint64)
}

type BigIntBitArray []*big.Int

func NewBigIntBitArray(size uint64) BigIntBitArray {
	numBits := int(1 + 10*size/64) // 10 bits/element
	arr := make(BigIntBitArray, numBits)
	for index := range arr {
		arr[index] = big.NewInt(0)
	}
	return arr
}

func (b BigIntBitArray) Get(index uint64) bool {
	intNum := int(index / uint64(len(b))) % len(b)
	offset := int(index % 64)
	if 1 == b[intNum].Bit(offset) {
		return true
	}
	return false
}

func (b BigIntBitArray) Set(index uint64) {
	intNum := int(index / uint64(len(b))) % len(b)
	offset := int(index % 64)
	b[intNum].SetBit(b[intNum], offset, 1)
}

type ByteArray []byte

func NewByteArray(size uint64) ByteArray {
	numBits := int(1 + 10*size/8) // 10 bits/element
	arr := make(ByteArray, numBits)
	return arr
}

func (b ByteArray) Get(index uint64) bool{
	byteNum := int(index / uint64(len(b))) % len(b)
	offset := int(index % 8)
	if b[byteNum] == b[byteNum] | byte(2^offset) {
		return true
	}
	return false
}

func (b ByteArray) Set(index uint64) {
	byteNum := int(index / uint64(len(b))) % len(b)
	offset := int(index % 8)
	b[byteNum] |= byte(2^offset)
}
