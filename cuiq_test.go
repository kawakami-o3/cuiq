package cuiq

import (
	"bytes"
	"testing"
)

type StreamIDData struct {
	bytes []byte
	id    StreamID
}

var streamIDs = []StreamIDData{
	StreamIDData{
		[]byte{0x25},
		37,
	},
	StreamIDData{
		[]byte{0x7b, 0xbd},
		15293,
	},
	StreamIDData{
		[]byte{0x9d, 0x7f, 0x3e, 0x7d},
		494878333,
	},
	StreamIDData{
		[]byte{0xc2, 0x19, 0x7c, 0x5e, 0xff, 0x14, 0xe8, 0x8c},
		151288809941952652,
	},
}

func TestEncodeStreamID(t *testing.T) {
	for _, tc := range streamIDs {
		buf := bytes.NewBuffer([]byte{})
		err := EncodeStreamID(buf, tc.id)
		if err != nil {
			t.Fatalf("Failed: error %v", err)
		} else if !bytes.Equal(tc.bytes, buf.Bytes()[:len(tc.bytes)]) {
			t.Fatalf("Failed: [%v] expected %v, but %v", tc.id, tc.bytes, buf.Bytes()[:len(tc.bytes)])
		}
	}
}

func TestDencodeStreamID(t *testing.T) {
	for _, tc := range streamIDs {
		b := tc.bytes
		buf := bytes.NewReader(b)
		id, err := DecodeStreamID(buf)
		if err != nil {
			t.Fatalf("Failed: error %v", err)
		} else if id != tc.id {
			t.Fatalf("Failed: expected %v, but %v", tc.id, id)
		}
	}
}
