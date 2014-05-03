package bloom

import (
	"testing"
)

func TestJenkinsBloomFilter(t *testing.T) {
	bf := NewJenkinsBloomFilter(10, 2)
	testSize := 100

	// fill with first 'n' digits
	for i := 0; i < testSize; i++ {
		bf.Add([]byte{byte(i)})
	}

	// check that first 'n' digits read as contained
	for i := 0; i < testSize; i++ {
		if !bf.Contains([]byte{byte(i)}) {
			t.Errorf("Bloom Filter does not contain %d\n", i)
		}
	}
}
