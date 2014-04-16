package set

import (
	"testing"
)

func newTestSet(nums []int) Set{
	s := NewSet()
	for _, i := range nums {
		s.Add(i)
	}
	return s
}

func TestNewSet(t *testing.T) {
	s := NewSet()

	if s.Cardinality() != 0 {
		t.Error("NewSet() should return an empty set");
	}
}
