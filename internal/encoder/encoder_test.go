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
	for _, test := range alphabetLengthTests {
		encoder := NewEncoder(test.alphabet)
		if encoder.nBitsPerRune != test.numBits {
			t.Errorf("Expected nBitsPerRune %d, actual nBitsPerRune %d", test.numBits, encoder.nBitsPerRune)
		}
	}
}

type encodeDefaultAlphabetTest struct {
	val     uint64
	encoded []rune
}

var encodeDefaultAlphabetTests = []encodeDefaultAlphabetTest{
	{4, []rune("0000000000004")},
	{128, []rune("0000000000040")},
	{^uint64(0), []rune("gzzzzzzzzzzzz")},
}

func TestEncodeDefaultAlphabet(t *testing.T) {
	for _, test := range encodeDefaultAlphabetTests {
		encoded := Base32encoder.Encode(test.val)
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
		encoded := encoder.Encode(test.value)
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

type DecodeTest struct {
	alphabet string
	hash     []rune
	val      uint64
}

var DecodeTests = []DecodeTest{
	{"abcd", []rune("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaba"), 4},
	{"0123456789bcdefghjkmnpqrstuvwxyz", []rune("012345678910"), 1199710202504224},
}

func TestDecode(t *testing.T) {
	for _, test := range DecodeTests {
		encoder := NewEncoder(test.alphabet)
		if decoded := encoder.Decode(test.hash); decoded != test.val {
			t.Errorf("Expected %d, got %d", test.val, decoded)
		}
	}
}

type EncodeDecodeReversibleTest struct {
	alphabet string
	val      uint64
}

var ReversibleEncodingDecodingTests = [...]EncodeDecodeReversibleTest{
	{"asdfg", 1234},
	{"0123456789bcdefghjkmnpqrstuvwxyz", 1199710202504224},
}

func TestEncodeDecodeReversible(t *testing.T) {
	for _, test := range ReversibleEncodingDecodingTests {
		encoder := NewEncoder(test.alphabet)
		if decoded := encoder.Decode(encoder.Encode(test.val)); decoded != test.val {
			t.Errorf("Expected %d, got %d", test.val, decoded)
		}
	}
}
