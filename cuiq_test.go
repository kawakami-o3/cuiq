package cuiq

import (
	"bytes"
	"testing"
)

func TestEncodeStreamID(t *testing.T) {
}

func TestDencodeStreamID(t *testing.T) {
	b := []byte{0x25}
	buf := bytes.NewReader(b)
	id := DecodeStreamID(buf)
	if id != 37 {
		t.Fatalf("Failed: expected %v, but %v", id, 37)
	}

}
