package bloom

import (
	"testing"
)

func TestJenkins(t *testing.T) {
	numTestKeys := 200
	used := make(map[uint64]bool)
	j := jenkinsHash{}

	for i := 0; i < numTestKeys; i++ {
		key, err := j.ComputeHash(i)
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
		j.ComputeHash(i)
	}
}

