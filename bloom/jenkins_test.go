package bloom

import (
	"encoding/binary"
	"testing"

	"crypto/sha1"
	"crypto/sha256"
)

// Test that the first 10k integers have unique hashes
func TestJenkins(t *testing.T) {
	numTestKeys := 10000
	used := make(map[uint64]bool)
	hashKey := make([]byte, 10)
	j := jenkinsHash{}

	for i := 0; i < numTestKeys; i++ {
		binary.PutUvarint(hashKey, uint64(i))
		key, err := j.ComputeHash(hashKey)
		if err != nil {
			t.Error(err.Error())
		}
		used[key] = true
	}

	if len(used) != numTestKeys {
		t.Errorf("Expected set of %d keys, found %s", numTestKeys, len(used))
	}
}

func BenchmarkJenkins(b *testing.B) {
	j := jenkinsHash{}
	for i := 0; i < b.N; i++ {
		j.ComputeHash([]byte{byte(i)})
	}
}

func BenchmarkSHA1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sha1.Sum([]byte{byte(i)})
	}
}

func BenchmarkSHA256(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sha256.Sum256([]byte{byte(i)})
	}
}
