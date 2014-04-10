package set

type Set map[interface{}]struct{}

// Returns a new empty Set
func NewSet() Set {
	return make(Set)
}

// Adds 'elem' to the Set
func (s Set) Add(elem interface{}) {
	s[elem] = struct{}{}
}

// Removes 'elem' from the Set
func (s Set) Remove(elem interface{}) {
	delete(s, elem)
}

// Returns True if elem is in the Set
func (s Set) Contains(elem interface{}) bool {
	_, found := s[elem]
	return found
}

// Return an array of the set contents
func (S Set) Array() []interface{} {
	keys := make([]interface{}, len(S))
	for k := range S {
		keys = append(keys, k)
	}
	return keys
}

// Find the union of all the elements in every set
func Union(sets ...Set) Set {
	S := NewSet()
	for _, s := range(sets) {
		for elem := range(s) {
			S.Add(elem)
		}
	}
	return S
}

// Take the intersection of all added sets
func Intersection(sets ...Set) Set {
	S := NewSet()
	for elem := range(sets[0]) {
		add := true
		for _, s := range(sets[1:]) {
			if !s.Contains(elem) {
				add = false
			}
		}
		if add {
			S.Add(elem)
		}
	}
	return S
}

// Find the difference between A and all other sets
func Difference (A Set, sets ...Set) Set {
	for _, s := range(sets) {
		for elem := range(s) {
			if A.Contains(elem) {
				A.Remove(elem)
			}
		}
	}
	return A
}

// Returns True if every element in s is in 'other'
func (S Set) SubSet (other Set) bool {
	for elem := range(S) {
		if !other.Contains(elem) {
			return false
		}
	}
	return true
}

// Returns True if every element in 'other' is in S
func (S Set) SuperSet (other Set) bool {
	for elem := range(other) {
		if !S.Contains(elem) {
			return false
		}
	}
	return true
}

// Returns True if S has no elements in common with 'other'
func (S Set) IsDisjoint (other Set) bool {
	for elem := range(S) {
		if other.Contains(elem) {
			return false
		}
	}
	for elem := range(other) {
		if S.Contains(elem) {
			return false
		}
	}

	return true
}

// Returns a set with elements that are in one set, but not multiple
func SymmetricDifferences(sets ...Set) Set {
	counter := make(map[interface{}]int)
	for _, s := range(sets) {
		for elem := range(s) {
			val, ok := counter[elem]
			if ok {
				counter[elem] = val + 1
			} else {
				counter[elem] = 1
			}
		}
	}

	S := NewSet()
	for elem, count := range(counter) {
		if count == 1 {
			S.Add(elem)
		}
	}

	return S
}
