package encoder

import (
	"testing"
)

type alphabetLengthTest struct {
	alphabet string
	numBits  uint
}

var alphabetLengthTests = []alphabetLengthTest{
	{"testÖ", 2},
	{"0123456789bcdefghjkmnpqrstuvwxyz", 5},
}

func TestNewEncoder(t *testing.T) {
	for _, expected := range alphabetLengthTests {
		encoder := NewEncoder(expected.alphabet)
		if encoder.nBitsPerRune != expected.numBits {
			t.Errorf("Expected nBitsPerRune %d, actual nBitsPerRune %d", expected.numBits, encoder.nBitsPerRune)
		}
	}
}

type encodeDefaultAlphabetTest struct {
	val     uint64
	encoded []rune
}

var encodeDefaultAlphabetTests = []encodeDefaultAlphabetTest{
	{4, []rune("000000000004")},
	{128, []rune("000000000040")},
	{^uint64(0), []rune("zzzzzzzzzzzz")},
}

func TestEncodeDefaultAlphabet(t *testing.T) {
	for _, test := range encodeDefaultAlphabetTests {
		encoded := Base32encoder.encode(test.val)
		for i, codePoint := range encoded {
			if codePoint != test.encoded[i] {
				t.Errorf("Expected %s, got %s", string(test.encoded), string(encoded))
			}
		}
	}
}

type customAlphabetEncodingTest struct {
	alphabet string
	value    uint64
	encoded  []rune
}

var customAlphabetEncodingTests = []customAlphabetEncodingTest{
	{"0abc", 4, []rune("000000000000000000000000000000a0")},
	{"0abc", ^uint64(0), []rune("cccccccccccccccccccccccccccccccc")},
}

func TestEncodeCustomAlphabet(t *testing.T) {
	for _, test := range customAlphabetEncodingTests {
		encoder := NewEncoder(test.alphabet)
		encoded := encoder.encode(test.value)
		for i, codePoint := range encoded {
			if codePoint != test.encoded[i] {
				t.Errorf("Expected %s, got %s", string(test.encoded), string(encoded))
			}
		}
	}
}

func TestEncodeCustomAlphabetNotPowerOfTwoExpectTruncate(t *testing.T) {
	alphabet := "abcde"
	encoder := NewEncoder(alphabet)
	if encoder.nBitsPerRune != 2 {
		t.Errorf("Expected number of bits per code point: 2. Actual value: %d", encoder.nBitsPerRune)
	}
	if string(encoder.alphabet) != "abcd" {
		t.Errorf("Expected alphabet 'abcd', got '%s'", string(encoder.alphabet))
	}

}
