package bloom

import (
	"bytes"
	"encoding/gob"
)

type jenkinsHash struct {
}

// adapted from http://bretmulvey.com/hash/7.html
func (_ jenkinsHash) ComputeHash(key interface{}) (uint64, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(key)
	if err != nil {
		return 0, err
	}
	data := buf.Bytes()

	var a, b, c uint64
	a, b = 0x9e3779b9, 0x9e3779b9
	c = 0
	i := 0

	for i = 0; i < len(data)-12; {
		a += uint64(data[i]) | uint64(data[i+1]<<8) | uint64(data[i+2]<<16) | uint64(data[i+3]<<24)
		i += 4
		b += uint64(data[i]) | uint64(data[i+1]<<8) | uint64(data[i+2]<<16) | uint64(data[i+3]<<24)
		i += 4
		c += uint64(data[i]) | uint64(data[i+1]<<8) | uint64(data[i+2]<<16) | uint64(data[i+3]<<24)

		a, b, c = mix(a, b, c)
	}

	c += uint64(len(data))

	if i < len(data) {
		a += uint64(data[i])
		i++
	}
	if i < len(data) {
		a += uint64(data[i]) << 8
		i++
	}
	if i < len(data) {
		a += uint64(data[i]) << 16
		i++
	}
	if i < len(data) {
		a += uint64(data[i]) << 24
		i++
	}

	if i < len(data) {
		b += uint64(data[i])
		i++
	}
	if i < len(data) {
		b += uint64(data[i]) << 8
		i++
	}
	if i < len(data) {
		b += uint64(data[i]) << 16
		i++
	}
	if i < len(data) {
		b += uint64(data[i]) << 24
		i++
	}

	if i < len(data) {
		c += uint64(data[i]) << 8
		i++
	}
	if i < len(data) {
		c += uint64(data[i]) << 16
		i++
	}
	if i < len(data) {
		c += uint64(data[i]) << 24
		i++
	}

	a, b, c = mix(a, b, c)
	return c, nil
}

func mix(a, b, c uint64) (uint64, uint64, uint64) {
	a -= b
	a -= c
	a ^= (c >> 13)
	b -= c
	b -= a
	b ^= (a << 8)
	c -= a
	c -= b
	c ^= (b >> 13)
	a -= b
	a -= c
	a ^= (c >> 12)
	b -= c
	b -= a
	b ^= (a << 16)
	c -= a
	c -= b
	c ^= (b >> 5)
	a -= b
	a -= c
	a ^= (c >> 3)
	b -= c
	b -= a
	b ^= (a << 10)
	c -= a
	c -= b
	c ^= (b >> 15)
	return a, b, c
}
