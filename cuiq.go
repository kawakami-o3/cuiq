package cuiq

import (
	"encoding/binary"
	"fmt"
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

func DecodeStreamID(r io.Reader) StreamID {
	b := make([]byte, 1)

	r.Read(b)

	fmt.Println(b)
	return 0
}
