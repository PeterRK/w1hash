//go:build cgo

package cgotest

/*
#include "../../w1hash.h"

static uint64_t w1hash_with_seed_n(const void* key, size_t len, uint64_t seed, size_t n) {
	uint64_t acc = 0;
	for (size_t i = 0; i < n; i++) {
		acc ^= w1hash_with_seed(key, len, seed);
	}
	return acc;
}
*/
import "C"

import "unsafe"

func HashWithSeed(b []byte, seed uint64) uint64 {
	if len(b) == 0 {
		return uint64(C.w1hash_with_seed(nil, 0, C.uint64_t(seed)))
	}
	return uint64(C.w1hash_with_seed(unsafe.Pointer(&b[0]), C.size_t(len(b)), C.uint64_t(seed)))
}

func HashWithSeedN(b []byte, seed uint64, n int) uint64 {
	if len(b) == 0 {
		return uint64(C.w1hash_with_seed_n(nil, 0, C.uint64_t(seed), C.size_t(n)))
	}
	return uint64(C.w1hash_with_seed_n(unsafe.Pointer(&b[0]), C.size_t(len(b)), C.uint64_t(seed), C.size_t(n)))
}

func Hash64(x uint64) uint64 {
	return uint64(C.w1hash64(C.uint64_t(x)))
}
