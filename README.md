
#Golang Data Structures


####Bloom Filter

Instantiated with an anticipated size and the number of hashes you wish to apply.
A Bob Jenkins hash is used, the result is fed as the data for the next application of the hash.

There are two methods, `Contains()` and `Add()`.
Both methods accept all data structures (`interface{}`) as elements of the Bloom filter.


####Set

Implements all methods from the [Python](https://docs.python.org/2/library/stdtypes.html#set) built-in set.

```go
func (S Set) Add(elem interface{})
func (S Set) AddList(elements []interface{})
func (S Set) Remove(elem interface{})
func (S Set) Contains(elem interface{}) bool
func (S Set) Cardinality() int 
func (S Set) Equal(other Set) bool
func (S Set) Array() []interface{}
func Union(sets ...Set) Set
func Intersection(sets ...Set) Set
func Difference(A Set, sets ...Set) Set
func (S Set) SubSet(other Set) bool
func (S Set) SuperSet(other Set) bool
func (S Set) IsDisjoint(other Set) bool
func SymmetricDifferences(sets ...Set) Set
func (S Set) Iterator() chan interface{}
```


