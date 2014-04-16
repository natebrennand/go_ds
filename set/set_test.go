package set

import (
	"testing"
)

func newTestSet(nums []int) Set {
	s := NewSet()
	for _, i := range nums {
		s.Add(i)
	}
	return s
}

func TestNewSet(t *testing.T) {
	s := NewSet()

	if s.Cardinality() != 0 {
		t.Error("NewSet() should return an empty set")
	}
}

func TestAdd(t *testing.T) {
	s := NewSet()
	elements := []int{1,2,3,4,5}

	for _, i := range elements {
		s.Add(i)
	}
	if s.Cardinality() != len(elements) {
		t.Errorf("Cardinality should be %d after Add() %d elements", len(elements), len(elements))
	}

	for _, i := range elements {
		s.Add(i)
	}
	if s.Cardinality() != len(elements) {
		t.Errorf("Cardinality should still be %d after readding elements", len(elements), len(elements))
	}
}

func TestAddList(t *testing.T) {
	s := NewSet()
	elements := []interface{}{1, 2, 3, 4, 5}

	s.AddList(elements)

	if s.Cardinality() != len(elements) {
		t.Errorf("Cardinality should be %d after AddList()", len(elements))
	}
}

func TestRemove(t *testing.T) {
	elements := []int{1,2,3,4,5}
	s := newTestSet(elements)

	for _, i := range elements {
		s.Remove(i)
		if s.Cardinality() != len(elements)-i {
			t.Errorf("Cardinality should be %d after %d Remove()", len(elements)-i, len(elements))
		}
	}
}

func TestContains(t *testing.T) {
	elements := []int{1,2,3,4,5}
	s := newTestSet(elements)

	for _, i := range elements {
		if !s.Contains(i) {
			t.Errorf("The set should contain %d", i)
		}
	}

	if s.Contains(0) {
		t.Error("The set should not contain 0")
	}
}

func TestCardinality(t *testing.T) {
	elements := []int{1,2,3,4,5}
	s := NewSet()

	for _, i := range elements {
		s.Add(i)
	}
	if s.Cardinality() != len(elements) {
		t.Errorf("Cardinality should be %d after Add() %d elements", len(elements), len(elements))
	}
}

func TestEqual(t *testing.T) {
	elements := []int{1,2,3,4,5}
	elements2 := []int{6,7,8,9,10}
	s := newTestSet(elements)
	a := newTestSet(elements)
	b := newTestSet(elements2)

	if !s.Equal(a) {
		t.Error("The sets should be equal")
	}

	if s.Equal(b) {
		t.Error("The sets should not be equal")
	}
}

func TestArray(t *testing.T) {
	elements := []int{1,2,3,4,5}
	s := newTestSet(elements)

	a := s.Array()

	for _, i := range elements {
		fail := true
		for _, j := range a {
			if i == j {
				fail = false
			}
		}
		if fail {
			t.Error("Element %d was not found in the produced array", i)
		}
	}
}
