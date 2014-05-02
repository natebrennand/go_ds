package bloom

type HashFn interface {
	ComputeHash(key interface{}) (uint64, error)
}
