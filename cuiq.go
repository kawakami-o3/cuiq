package cuiq

import (
	"bytes"
	"encoding/binary"
	"io"
	"math"
)

type StreamID int64

func EncodeStreamID(w io.Writer, streamID StreamID) error {
	var err error
	if streamID < 64 {
		err = binary.Write(w, binary.BigEndian, byte(streamID))
	} else if streamID < 16384 {
		buf := bytes.NewBuffer([]byte{})
		err = binary.Write(buf, binary.BigEndian, int16(streamID))
		if err != nil {
			return err
		}

		bs := buf.Bytes()
		bs[0] += 0x40
		_, err = w.Write(buf.Bytes())
	} else if streamID < 1073741824 {
		buf := bytes.NewBuffer([]byte{})
		err = binary.Write(buf, binary.BigEndian, int32(streamID))
		if err != nil {
			return err
		}

		bs := buf.Bytes()
		bs[0] += 0x80
		_, err = w.Write(buf.Bytes())
	} else {
		buf := bytes.NewBuffer([]byte{})
		err = binary.Write(buf, binary.BigEndian, streamID)
		if err != nil {
			return err
		}

		bs := buf.Bytes()
		bs[0] += 0xc0
		_, err = w.Write(buf.Bytes())
	}

	return err
}

func DecodeStreamID(r io.Reader) (StreamID, error) {
	b := make([]byte, 1)
	_, err := r.Read(b)
	if err != nil {
		return 0, err
	}

	flagBits := b[0] >> 6
	buf := []byte{b[0] & 0x3f}

	length := 1
	if flagBits > 0 {
		length = int(math.Pow(2, float64(flagBits)))
		bs := make([]byte, length-1)
		_, err := r.Read(bs)
		if err != nil {
			return 0, err
		}
		buf = append(buf, bs...)
	}
	buf = append(make([]byte, 8-length), buf...)

	var id int64
	err = binary.Read(bytes.NewReader(buf), binary.BigEndian, &id)
	return StreamID(id), err
}
