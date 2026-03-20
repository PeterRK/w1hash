//go:build amd64

package w1hash

//go:noescape
func HashWithSeed(key []byte, seed uint64) uint64

func Hash(key []byte) uint64 {
	return HashWithSeed(key, 0)
}

//go:noescape
func Hash64(x uint64) uint64
