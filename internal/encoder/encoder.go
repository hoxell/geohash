package encoder

import (
	"math"
)

// Encoder type defining the encoding scheme to be used
type Encoder struct {
	nBitsPerRune uint
	alphabet     []rune
}

// NewEncoder creates a new encoder object using the provided UTF-8 alphabet.
// If len(alphabet) is not a power of 2, the encoder will use only
// the UTF-8 characters up until the largest power of 2
func NewEncoder(alphabet string) *Encoder {
	var encoder Encoder
	encoder.alphabet = []rune(alphabet)
	encoder.nBitsPerRune = uint(math.Log2(float64(len(encoder.alphabet))))
	return &encoder
}

func (encoder *Encoder) encode(value uint64) []rune {
	nRunes := 64 / encoder.nBitsPerRune
	encoded := make([]rune, nRunes)
	bitPattern := uint64(math.Exp2(float64(encoder.nBitsPerRune))) - 1
	for i := uint(0); i < nRunes; i++ {
		encoded[nRunes-1-i] = encoder.alphabet[value&bitPattern]
		value >>= encoder.nBitsPerRune
	}
	return encoded
}

var Base32encoder = NewEncoder("0123456789bcdefghjkmnpqrstuvwxyz")
