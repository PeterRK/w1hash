//go:build cgo

package w1hash

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/peterrk/w1hash/internal/cgotest"
)

func hashWithSeedN(key []byte, seed uint64, n int) uint64 {
	var acc uint64
	for i := 0; i < n; i++ {
		acc ^= HashWithSeed(key, seed)
	}
	return acc
}

func hash64N(x uint64, n int) uint64 {
	var acc uint64
	for i := 0; i < n; i++ {
		acc ^= Hash64(x)
	}
	return acc
}

func TestHashMatchesC(t *testing.T) {
	seeds := []uint64{
		0,
		1,
		0x123456789abcdef0,
		0xffffffffffffffff,
	}

	for n := 0; n <= 256; n++ {
		buf := make([]byte, n)
		for i := range buf {
			buf[i] = byte(i*131 + n)
		}
		for _, seed := range seeds {
			got := HashWithSeed(buf, seed)
			want := cgotest.HashWithSeed(buf, seed)
			if got != want {
				t.Fatalf("len=%d seed=%#x got=%#x want=%#x", n, seed, got, want)
			}
		}
	}
}

func TestHashRandomMatchesC(t *testing.T) {
	rng := rand.New(rand.NewSource(1))
	for i := 0; i < 2000; i++ {
		n := rng.Intn(4096)
		buf := make([]byte, n)
		if _, err := rng.Read(buf); err != nil {
			t.Fatalf("rand.Read: %v", err)
		}
		seed := rng.Uint64()
		got := HashWithSeed(buf, seed)
		want := cgotest.HashWithSeed(buf, seed)
		if got != want {
			t.Fatalf("case=%d len=%d seed=%#x got=%#x want=%#x", i, n, seed, got, want)
		}
	}
}

func TestHash64MatchesC(t *testing.T) {
	rng := rand.New(rand.NewSource(2))
	values := []uint64{
		0,
		1,
		0xffffffffffffffff,
		0x123456789abcdef0,
	}
	for i := 0; i < 2000; i++ {
		values = append(values, rng.Uint64())
	}
	for _, v := range values {
		got := Hash64(v)
		want := cgotest.Hash64(v)
		if got != want {
			t.Fatalf("value=%#x got=%#x want=%#x", v, got, want)
		}
	}
}

func BenchmarkHash(b *testing.B) {
	seed := uint64(0x123456789abcdef0)
	rng := rand.New(rand.NewSource(3))
	for size := 0; size <= 32; size++ {
		buf := make([]byte, size)
		if _, err := rng.Read(buf); err != nil {
			b.Fatalf("rand.Read: %v", err)
		}

		b.Run(fmt.Sprintf("go/%dB", size), func(b *testing.B) {
			b.SetBytes(int64(size))
			b.ResetTimer()
			_ = hashWithSeedN(buf, seed, b.N)
		})

		b.Run(fmt.Sprintf("c/%dB", size), func(b *testing.B) {
			b.SetBytes(int64(size))
			b.ResetTimer()
			_ = cgotest.HashWithSeedN(buf, seed, b.N)
		})
	}
}
