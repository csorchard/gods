package cms

import "encoding/binary"

// get the two basic hash function values for data.
func (s *CountMinSketch) baseHashes(key []byte) (a uint32, b uint32) {
	s.hasher.Reset()
	s.hasher.Write(key)
	sum := s.hasher.Sum(nil)
	upper := sum[0:4]
	lower := sum[4:8]
	a = binary.BigEndian.Uint32(lower)
	b = binary.BigEndian.Uint32(upper)
	return
}

// Get the _w_ locations to update/Estimate
// Based on https://github.com/willf/bloom/blob/master/bloom.go
func (s *CountMinSketch) locations(key []byte) (locs []uint) {
	locs = make([]uint, s.d)
	a, b := s.baseHashes(key)
	ua := uint(a)
	ub := uint(b)
	for r := uint(0); r < s.d; r++ {
		locs[r] = (ua + ub*r) % s.w
	}
	return
}
