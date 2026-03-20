package w1hash

import "math/bits"

const (
	seed0 = 0x2d358dccaa6c78a5
	seed1 = 0x8bb84b93962eacc9
	seed2 = 0x4b33a62ed433d4a3
	seed3 = 0x4d5a2da51de1aa47
)

type u128 struct {
	lo uint64
	hi uint64
}

func mul128(a, b uint64) u128 {
	hi, lo := bits.Mul64(a, b)
	return u128{lo: lo, hi: hi}
}

func mix(a, b uint64) uint64 {
	t := mul128(a, b)
	return t.lo ^ t.hi
}

func read1(p []byte) uint64 { return uint64(p[0]) }

func read2(p []byte) uint64 {
	return uint64(p[0]) |
		uint64(p[1])<<8
}

func read3(p []byte) uint64 {
	return read2(p) |
		uint64(p[2])<<16
}

func read4(p []byte) uint64 {
	return read2(p) |
		read2(p[2:])<<16
}

func read5(p []byte) uint64 {
	return read4(p) |
		uint64(p[4])<<32
}

func read6(p []byte) uint64 {
	return read4(p) |
		read2(p[4:])<<32
}

func read7(p []byte) uint64 {
	return read4(p) |
		read2(p[4:])<<32 |
		uint64(p[6])<<48
}

func read8(p []byte) uint64 {
	return read4(p) |
		read4(p[4:])<<32
}

func HashWithSeed(key []byte, seed uint64) uint64 {
	length := len(key)
	seed ^= mix(seed^seed0, uint64(length)^seed1)

	p := key
	l := length

	for {
		var t u128
		switch l {
		case 0:
		case 1:
			t.lo = read1(p)
		case 2:
			t.lo = read2(p)
		case 3:
			t.lo = read3(p)
		case 4:
			t.lo = read4(p)
		case 5:
			t.lo = read5(p)
		case 6:
			t.lo = read6(p)
		case 7:
			t.lo = read7(p)
		case 8:
			t.lo = read8(p)
		case 9:
			t.lo = read8(p)
			t.hi = read1(p[8:])
		case 10:
			t.lo = read8(p)
			t.hi = read2(p[8:])
		case 11:
			t.lo = read8(p)
			t.hi = read3(p[8:])
		case 12:
			t.lo = read8(p)
			t.hi = read4(p[8:])
		case 13:
			t.lo = read8(p)
			t.hi = read5(p[8:])
		case 14:
			t.lo = read8(p)
			t.hi = read6(p[8:])
		case 15:
			t.lo = read8(p)
			t.hi = read7(p[8:])
		case 16:
			t.lo = read8(p)
			t.hi = read8(p[8:])
		default:
			if l > 64 {
				x := seed
				y := seed
				z := seed
				for l > 64 {
					seed = mix(read8(p)^seed0, read8(p[8:])^seed)
					x = mix(read8(p[16:])^seed1, read8(p[24:])^x)
					y = mix(read8(p[32:])^seed2, read8(p[40:])^y)
					z = mix(read8(p[48:])^seed3, read8(p[56:])^z)
					p = p[64:]
					l -= 64
				}
				seed ^= x ^ y ^ z
			}
			if l > 32 {
				x := seed
				seed = mix(read8(p)^seed0, read8(p[8:])^seed)
				x = mix(read8(p[16:])^seed1, read8(p[24:])^x)
				seed ^= x
				p = p[32:]
				l -= 32
			}
			if l > 16 {
				seed = mix(read8(p)^seed0, read8(p[8:])^seed)
				p = p[16:]
				l -= 16
			}
			continue
		}
		t = mul128(t.lo^seed1, t.hi^seed)
		return mix(t.lo^(seed0^uint64(length)), t.hi^seed1)
	}
}

func Hash(key []byte) uint64 {
	return HashWithSeed(key, 0)
}

func Hash64(x uint64) uint64 {
	const (
		s0 = 0x2d358dccaa6c78ad
		s1 = seed1
		k  = 0x702daa6e740fb546
	)
	t := mul128(x^s1, k)
	return mix(t.lo^s0, t.hi^s1)
}
