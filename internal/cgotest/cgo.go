//go:build cgo

package cgotest

/*
#include "../../w1hash.h"
*/
import "C"

import "unsafe"

func HashWithSeed(b []byte, seed uint64) uint64 {
	if len(b) == 0 {
		return uint64(C.w1hash_with_seed(nil, 0, C.uint64_t(seed)))
	}
	return uint64(C.w1hash_with_seed(unsafe.Pointer(&b[0]), C.size_t(len(b)), C.uint64_t(seed)))
}

func Hash64(x uint64) uint64 {
	return uint64(C.w1hash64(C.uint64_t(x)))
}
