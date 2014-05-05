package set

type Set map[interface{}]struct{}

// Returns a new empty Set
func NewSet() Set {
	return make(Set)
}

// Adds 'elem' to the Set
func (S Set) Add(elem interface{}) {
	S[elem] = struct{}{}
}

// Adds the list of elements to the set
func (S Set) AddList(elements []interface{}) {
	for _, elem := range elements {
		S.Add(elem)
	}
}

// Removes 'elem' from the Set
func (S Set) Remove(elem interface{}) {
	delete(S, elem)
}

// Returns True if elem is in the Set
func (S Set) Contains(elem interface{}) bool {
	_, found := S[elem]
	return found
}

// Returns the cardinality (size) of the set
func (S Set) Cardinality() int {
	return len(S)
}

// Returns true if the sets are found to be equivalent
func (S Set) Equal(other Set) bool {
	if S.Cardinality() != other.Cardinality() {
		return false
	}

	intersect := S.Intersection(other)
	if intersect.Cardinality() != other.Cardinality() {
		return false
	}

	return true
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
func (S Set) Union(sets ...Set) Set {
	A := NewSet()
	for _, s := range append(sets, S) {
		for elem := range s {
			A.Add(elem)
		}
	}
	return A
}

// Take the intersection of all added sets
func (S Set) Intersection(sets ...Set) Set {
	A := NewSet()
	for elem := range S {
		add := true
		for _, s := range sets {
			if !s.Contains(elem) {
				add = false
			}
		}
		if add {
			A.Add(elem)
		}
	}
	return A
}

// Find the difference between A and all other sets
func (S Set) Difference(sets ...Set) Set {
	A := NewSet()
	A.AddList(S.Array())
	for _, s := range sets {
		for elem := range s {
			if S.Contains(elem) {
				S.Remove(elem)
			}
		}
	}
	return S
}

// Returns True if every element in s is in 'other'
func (S Set) SubSet(other Set) bool {
	for elem := range S {
		if !other.Contains(elem) {
			return false
		}
	}
	return true
}

// Returns True if every element in 'other' is in S
func (S Set) SuperSet(other Set) bool {
	for elem := range other {
		if !S.Contains(elem) {
			return false
		}
	}
	return true
}

// Returns True if S has no elements in common with 'other'
func (S Set) Disjoint(other Set) bool {
	for elem := range S {
		if other.Contains(elem) {
			return false
		}
	}
	return true
}

// Returns a set with elements that are in one set, but not multiple
func (S Set) SymmetricDifferences(sets ...Set) Set {
	counter := make(map[interface{}]int)
	for _, s := range append(sets, S) {
		for elem := range s {
			val, ok := counter[elem]
			if ok {
				counter[elem] = val + 1
			} else {
				counter[elem] = 1
			}
		}
	}

	X := NewSet()
	for elem, count := range counter {
		if count == 1 {
			X.Add(elem)
		}
	}

	return X
}

// Returns a channel iterator for the set
func (S Set) Iterator() chan interface{} {
	c := make(chan interface{})
	go func() {
		for elem := range S {
			c <- elem
		}
		close(c)
	}()
	return c
}
