package geohash

import (
	"testing"
)

type point struct {
	lat float64
	lon float64
}

func TestEncoding(t *testing.T) {
	expected := map[point]string{
		{1.0, 1.0}:               `s00twy01mtw0`,
		{17.123456, 45.987654}:   `t509qmppq581`,
		{-17.123456, 45.987654}:  `mhbww6z0wh2n`,
		{-17.123456, -45.987654}: `6uzq9dbb9ury`,
	}
	for coordinates, hash := range expected {
		if val := Encode(coordinates.lat, coordinates.lon); val != hash {
			t.Errorf("Got %s, but expected %s", val, hash)
		}
	}
}

type EncodingPrecisionTest struct {
	lat       float64
	lon       float64
	precision uint8
	hash      string
}

var EncodingPrecisionTests = []EncodingPrecisionTest{
	{1.0, 1.0, 4, `s00t`},
	{17.123456, 45.987654, 5, `t509q`},
	{-17.123456, 45.987654, 6, `mhbww6`},
	{-17.123456, -45.987654, 20, `6uzq9dbb9ury`},
	{-17.123456, -45.987654, 0, ``},
}

func TestEncodingPrecision(t *testing.T) {
	for _, test := range EncodingPrecisionTests {
		if hash := EncodePrecision(test.lat, test.lon, test.precision); hash != test.hash {
			t.Errorf("Expected %s, got %s", test.hash, hash)
		}
	}
}

type segmentIndex struct {
	lat uint32
	lon uint32
}

func TestMortonCode(t *testing.T) {
	expected := map[segmentIndex]uint64{
		{10, 9}:         198,
		{10000, 151253}: 34976146210,
	}
	for indexes, expectedValue := range expected {
		if val := mortonCode(indexes.lat, indexes.lon); val != expectedValue {
			t.Errorf("Got %d, expected %d", val, expectedValue)
		}
	}
}

func TestBase32Encode(t *testing.T) {
	expected := map[uint64]string{
		0b10111010110001001110010011100010: `0000005usmkf`,
	}
	for value, hash := range expected {
		if val := base32encode(value); val != hash {
			t.Errorf("Got %s, but expected %s", val, hash)
		}
	}
}
