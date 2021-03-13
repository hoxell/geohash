package encoder

import (
	"fmt"
	"testing"
)

func TestNewEncoder(t *testing.T) {
	expected := map[string]uint{"testÖ": 2,
		"0123456789bcdefghjkmnpqrstuvwxyz": 5}
	for alphabet, nBits := range expected {
		encoder := NewEncoder(alphabet)
		if encoder.nBitsPerRune != nBits {
			t.Errorf("Expected nBitsPerRune %d, actual nBitsPerRune %d", nBits, encoder.nBitsPerRune)
		}
	}
}

func TestEncodeDefaultAlphabet(t *testing.T) {
	expected := map[uint64][]rune{
		4:          []rune("000000000004"),
		128:        []rune("000000000040"),
		^uint64(0): []rune("zzzzzzzzzzzz"),
	}
	for val, expected := range expected {
		encoded := Base32encoder.encode(val)
		for i, codePoint := range encoded {
			if codePoint != expected[i] {
				fmt.Println(val)
				t.Errorf("Expected %s, got %s", string(expected), string(encoded))
			}
		}
	}
}

func TestEncodeCustomAlphabet(t *testing.T) {
	alphabet := `0abc`
	encoder := NewEncoder(alphabet)
	expected := map[uint64][]rune{
		4:          []rune("000000000000000000000000000000a0"),
		^uint64(0): []rune("cccccccccccccccccccccccccccccccc"),
	}
	for val, expected := range expected {
		encoded := encoder.encode(val)
		for i, codePoint := range encoded {
			if codePoint != expected[i] {
				fmt.Println(val)
				t.Errorf("Expected %s, got %s", string(expected), string(encoded))
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
