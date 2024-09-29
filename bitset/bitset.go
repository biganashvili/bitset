package bitset

const size = 32

type bits uint32

// BitSet is a set of bits that can be set, cleared and queried.
type BitSet []bits

// Set ensures that the given bit is set in the BitSet.
func (s *BitSet) Set(i uint32) {
	if len(*s) < int(i/size+1) {
		r := make([]bits, i/size+1)
		copy(r, *s)
		*s = r
	}
	(*s)[i/size] |= 1 << (i % size)
}

// Clear ensures that the given bit is cleared (not set) in the BitSet.
func (s *BitSet) Clear(i uint32) {
	if len(*s) >= int(i/size+1) {
		(*s)[i/size] &^= 1 << (i % size)
	}
}

// IsSet returns true if the given bit is set, false if it is cleared.
func (s *BitSet) IsSet(i uint32) bool {
	if len(*s) >= int(i/size+1) {
		return (*s)[i/size]&(1<<(i%size)) != 0
	}
	return false
}
