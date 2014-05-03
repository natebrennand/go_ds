
[![Build Status](https://travis-ci.org/natebrennand/go_ds.svg?branch=master)](https://travis-ci.org/natebrennand/go_ds)

#Golang Data Structures


####Bloom Filter

Instantiated with an anticipated size and the number of hashes you wish to apply.
A [Bob Jenkins hash](bloom/jenkins.go) is used, the result is fed as the data for the next application of the hash.

There are two methods, `Contains()` and `Add()`.
Both methods accept all data structures (`interface{}`) as elements of the Bloom filter.

```go
(b BloomFilter) Add(elem interface{}) {
(b BloomFilter) Contains(elem interface{}) bool {
```

####Set

Implements all methods from the [Python](https://docs.python.org/2/library/stdtypes.html#set) built-in set.

```go
(S Set) Add(elem interface{})
(S Set) AddList(elements []interface{})
(S Set) Remove(elem interface{})
(S Set) Contains(elem interface{}) bool
(S Set) Cardinality() int 
(S Set) Equal(other Set) bool
(S Set) Array() []interface{}
(S Set) Union(sets ...Set) Set
(S Set) Intersection(sets ...Set) Set
(S Set) Difference(sets ...Set) Set
(S Set) SubSet(other Set) bool
(S Set) SuperSet(other Set) bool
(S Set) Disjoint(other Set) bool
SymmetricDifferences(sets ...Set) Set
(S Set) Iterator() chan interface{}
```


