// Code generated from canonical protocol manifest v2. DO NOT EDIT.

package protocol

// Bitset131 stores the 131-bit value used by the wire bitset encoding.
type Bitset131 [3]uint64

const Bitset131Length = 131

// Set marks bit index i. It panics when i is outside [0, 131).
func (b *Bitset131) Set(i int) {
	if i < 0 || i >= Bitset131Length {
		panic("index out of bounds")
	}
	b[i/64] |= uint64(1) << uint(i%64)
}

// Unset clears bit index i. It panics when i is outside [0, 131).
func (b *Bitset131) Unset(i int) {
	if i < 0 || i >= Bitset131Length {
		panic("index out of bounds")
	}
	b[i/64] &^= uint64(1) << uint(i%64)
}

// Load reports whether bit index i is set. It panics when i is outside [0, 131).
func (b Bitset131) Load(i int) bool {
	if i < 0 || i >= Bitset131Length {
		panic("index out of bounds")
	}
	return b[i/64]&(uint64(1)<<uint(i%64)) != 0
}

// Len returns the number of bits in the bitset.
func (b Bitset131) Len() int { return Bitset131Length }
