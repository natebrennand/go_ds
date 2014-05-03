package bloom

import (
	"math"
	"math/big"
	"testing"
)

func getPowerOf2(x int) int {
	return int(math.Pow(float64(2), float64(x)))
}

func TestBigIntNewArray(t *testing.T) {
	size := 10
	arr := NewBigIntBitArray(size)
	if len(arr)*64 != getPowerOf2(size) {
		t.Errorf("Expected %d bits, found %d",
			getPowerOf2(size),
			len(arr)*64,
		)
	}

	size = 6
	arr = NewBigIntBitArray(size)
	if len(arr)*64 != getPowerOf2(size) {
		t.Errorf("Expected %d bits, found %d",
			getPowerOf2(size),
			len(arr)*64,
		)
	}
}

func TestBigIntGetIndexes(t *testing.T) {
	size := 6
	arr := NewBigIntBitArray(size)
	bucket, offset := arr.getIndexes(uint64(5))
	if bucket != 0 {
		t.Errorf("1st bucket should've been returned, instead %d was", bucket)
	}
	if offset != 5 {
		t.Error("Offset=%d should've been returned, instead %d was", 5, offset)
	}

	size = 9
	arr = NewBigIntBitArray(size)
	bucket, offset = arr.getIndexes(uint64(250))
	if bucket != 3 {
		t.Errorf("4th bucket should've been returned, instead %d was", bucket)
	}
	if offset != 58 { // 64 - (256 - 250)
		t.Error("Offset=%d should've been returned, instead %d was", 58, offset)
	}
}

func TestBigIntBitArrayGet(t *testing.T) {
	size := 6
	arr := NewBigIntBitArray(size)

	for indexPower2 := 0; indexPower2 < 64; indexPower2++ {
		arr[0] = big.NewInt(int64(getPowerOf2(indexPower2)))
		if !arr.Get(uint64(indexPower2)) {
			t.Errorf("Expected bit %d to be turned on", indexPower2)
		}
	}
}

func TestBigIntBitArraySet(t *testing.T) {
	size := 6
	arr := NewBigIntBitArray(size)

	for indexPower2 := 0; indexPower2 < 63; indexPower2++ {
		arr.Set(uint64(indexPower2))
		if arr[0].Int64() != int64(getPowerOf2(indexPower2+1)-1) {
			t.Errorf("Expected all bits up to %d to be set.", indexPower2)
		}
	}
}
