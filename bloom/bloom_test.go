package bloom

import (
	"math"
	"testing"
)

func TestJenkinsBloomFilter(t *testing.T) {
	j := NewJenkinsBloomFilter(uint64(math.Pow(2, 10)), 2)
	testSize := 100

	// fill with first 'n' digits
	for i := 0; i < testSize; i++ {
		j.Add(i)
	}

	// check that first 'n' digits read as contained
	for i := 0; i < testSize; i++ {
		if !j.Contains(i) {
			t.Errorf("Bloom Filter does not contain %d\n", i)
		}
	}

	// check if any of the n - n*10 digits show up as contained
	for i := testSize; i < 10*testSize; i++ {
		if j.Contains(i) {
			t.Errorf("Bloom Filter does not contain %d\n", i)
		}
	}
}
