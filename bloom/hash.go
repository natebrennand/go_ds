package bloom

type HashFn interface {
	ComputeHash([]byte) (uint64, error)
}
