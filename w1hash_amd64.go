//go:build amd64 && !purego

package w1hash

//go:noescape
func HashWithSeed(key []byte, seed uint64) uint64

//go:noescape
func Hash(key []byte) uint64

//go:noescape
func Hash64(x uint64) uint64
