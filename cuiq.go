package cuiq

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"
)

type StreamID int64

func EncodeStreamID(w io.Writer, streamID StreamID) error {
	err := binary.Write(w, binary.BigEndian, streamID)
	if err != nil {
		return err
	}

	return nil
}

func DecodeStreamID(r io.Reader) (StreamID, error) {
	b := make([]byte, 1)

	_, err := r.Read(b)
	if err != nil {
		return 0, err
	}

	flagBits := b[0] & 0xc0
	b[0] &= 0x3f
	if flagBits == 0x00 {
		return StreamID(b[0]), nil
	} else if flagBits == 0x40 {
		bs := make([]byte, 1)
		_, err := r.Read(bs)
		if err != nil {
			return 0, err
		}

		buf := bytes.NewReader(append(b, bs...))
		var id int16
		err = binary.Read(buf, binary.BigEndian, &id)
		if err != nil {
			return 0, err
		}
		return StreamID(id), nil
	} else if flagBits == 0x80 {
		bs := make([]byte, 3)
		_, err := r.Read(bs)
		if err != nil {
			return 0, err
		}

		buf := bytes.NewReader(append(b, bs...))
		var id int32
		err = binary.Read(buf, binary.BigEndian, &id)
		if err != nil {
			return 0, err
		}
		return StreamID(id), nil
	} else if flagBits == 0xc0 {
		bs := make([]byte, 7)
		_, err := r.Read(bs)
		if err != nil {
			return 0, err
		}

		buf := bytes.NewReader(append(b, bs...))
		var id int64
		err = binary.Read(buf, binary.BigEndian, &id)
		if err != nil {
			return 0, err
		}
		return StreamID(id), nil
	}

	return 0, errors.New("invalid stream ID")
}
