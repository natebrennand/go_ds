package set

import (
	"testing"
)

// help fn
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
	elements3 := []int{6,7,8,9,10,11,12,13}
	s := newTestSet(elements)
	a := newTestSet(elements)
	b := newTestSet(elements2)
	c := newTestSet(elements3)

	if !s.Equal(a) {
		t.Error("The sets should be equal")
	}

	if s.Equal(b) {
		t.Error("The sets should not be equal")
	}

	if s.Equal(c) {
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

func TestUnion(t *testing.T) {
	elements := []int{1,2,3,4,5}
	elements2 := []int{6,7,8,9,10}
	a := newTestSet(elements)
	b := newTestSet(elements2)
	c := newTestSet(append(elements, elements2...))


	s := a.Union(b)
	if !s.Equal(c) {
		t.Errorf("Error using Union()\nExpected: %v\nRecieved: %v",
			c.Array(),
			s.Array(),
		)
	}
}

func TestIntersection(t *testing.T) {
	elements := []int{4,5,6,7}
	elements2 := []int{1,2,3,4,5,6,7}
	elements3 := []int{4,5,6,7,8,9,10}
	a := newTestSet(elements2)
	b := newTestSet(elements3)
	c := newTestSet(elements)


	s := a.Intersection(b)
	if !s.Equal(c) {
		t.Errorf("Error using Intersection()\nExpected: %v\nRecieved: %v",
			c.Array(),
			s.Array(),
		)
	}
}

func TestDifference(t *testing.T) {
	elements := []int{1,2,3}
	elements2 := []int{1,2,3,4,5,6,7}
	elements3 := []int{4,5,6,7,8,9,10}
	c := newTestSet(elements)
	a := newTestSet(elements2)
	b := newTestSet(elements3)


	s := a.Difference(b)
	if !s.Equal(c) {
		t.Errorf("Error using Intersection()\nExpected: %v\nRecieved: %v",
			c.Array(),
			s.Array(),
		)
	}
}

func TestSubSet(t *testing.T) {
	elements := []int{1,2,3}
	elements2 := []int{1,2,3,4,5,6,7}
	elements3 := []int{4,5,6,7,8,9,10}
	a := newTestSet(elements)
	b := newTestSet(elements2)
	c := newTestSet(elements3)

	if a.SubSet(c) {
		t.Errorf("Should not be a subset")
	}
	if b.SubSet(c) {
		t.Errorf("Should not be a subset")
	}
	if !a.SubSet(b) {
		t.Errorf("Should be a subset")
	}
	if !b.SubSet(b) {
		t.Errorf("A set should be a subset of itself")
	}
	if !NewSet().SubSet(a) {
		t.Errorf("A empty set should be a subset of all other sets")
	}
}

func TestSuperSet(t *testing.T) {
	elements := []int{1,2,3}
	elements2 := []int{1,2,3,4,5,6,7}
	elements3 := []int{4,5,6,7,8,9,10}
	a := newTestSet(elements)
	b := newTestSet(elements2)
	c := newTestSet(elements3)

	if !b.SuperSet(a) {
		t.Errorf("Should be a superset")
	}
	if b.SuperSet(c) {
		t.Errorf("Should not be a superset")
	}
	if a.SuperSet(b) {
		t.Errorf("Should not be a superset")
	}
	if !b.SuperSet(b) {
		t.Errorf("A set should be a superset of itself")
	}
	if !a.SuperSet(NewSet()) {
		t.Errorf("Any set should be the superset of an empty set")
	}
}

func TestDisjoint(t *testing.T) {
	elements := []int{1,2,3}
	elements2 := []int{1,2,3,4,5,6,7}
	elements3 := []int{4,5,6,7,8,9,10}
	a := newTestSet(elements)
	b := newTestSet(elements2)
	c := newTestSet(elements3)

	if !a.Disjoint(c) {
		t.Errorf("Should be disjoint")
	}
	if b.Disjoint(c) {
		t.Errorf("Should not be a superset")
	}
	if a.Disjoint(b) {
		t.Errorf("Should not be a superset")
	}
}
