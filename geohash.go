package geohash

import (
	"math"

	"github.com/hoxovic/geohash/internal/encoder"
)

// EncodePrecision computes the geohash for lat, lon and returns the precision
// most significant code points of the hash.
// precision is bounded to [0, 12]
func EncodePrecision(lat, lon float64, precision uint8) string {
	fullHash := []rune(Encode(lat, lon))
	desiredPrecision := uint8(math.Min(12, float64(precision)))
	return string(fullHash[:desiredPrecision])
}

// Encode latitude and longitude into a 12-level geohash
func Encode(lat, lon float64) string {
	latSegmentIdx := computeSegmentIdx(lat, 90)
	lonSegmentIdx := computeSegmentIdx(lon, 180)
	mortonCode := mortonCode(latSegmentIdx, lonSegmentIdx)
	// Base32 -> 60bits is the greatest multiple of 5bits that is less than 64 bits.
	// Discard the 4 least significant bits.
	mortonCode >>= 4
	encoded := encoder.Base32encoder.Encode(mortonCode)
	return string(encoded[1:])
}

// Compute the index of the subsegment of the 2^32 equally sized discretized
// subsegments of [-angle, angle] that pos is contained by.
// The zero-based indexing starts at [-angle, -angle + 2^-32)
func computeSegmentIdx(pos, angle float64) uint32 {
	return uint32((pos + angle) / (2 * angle) * math.Exp2(32))
}

// Interleave the binary representation of a and b such that
// the MSB of b becomes the MSB of the final result
// https://graphics.stanford.edu/~seander/bithacks.html#InterleaveTableObvious
func mortonCode(a, b uint32) uint64 {
	return bitSpreadEven(a) | bitSpreadEven(b)<<1
}

// Spread the unsigned integer x into an unsigned integer of double the size
// such that the bits of x occupy the even bits of the new unsigned integer
// Effectively, this inserts a zero between all bits (and a leading zero)
// Based on the magic bits method
// https://graphics.stanford.edu/~seander/bithacks.html#InterleaveTableObvious
func bitSpreadEven(x uint32) uint64 {
	spread := uint64(x)
	spread = (spread | (spread << 16)) & 0x0000FFFF0000FFFF
	spread = (spread | (spread << 8)) & 0x00FF00FF00FF00FF
	spread = (spread | (spread << 4)) & 0x0F0F0F0F0F0F0F0F
	spread = (spread | (spread << 2)) & 0x3333333333333333
	spread = (spread | (spread << 1)) & 0x5555555555555555
	return spread
}
