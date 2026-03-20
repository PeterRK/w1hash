//go:build cgo

package w1hash

import (
	"math/rand"
	"testing"

	"github.com/peterrk/w1hash/internal/cgotest"
)

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
