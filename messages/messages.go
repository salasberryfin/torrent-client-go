package messages

import (
	"encoding/binary"
	"fmt"
	"io"
)

func Read(r io.Reader) (*Message, error) {
	lengthBuf := make([]byte, 4)
	_, err := io.ReadFull(r, lengthBuf)
	if err != nil {
		return nil, err
	}
	length := binary.BigEndian.Uint32(lengthBuf)
	//buf := make([]byte, 4)
	//_, err := io.ReadFull(r, buf)
	//if err != nil {
	//	return &Message{}, fmt.Errorf("Unable to parse received message: %v\n", err)
	//}
	//length := binary.BigEndian.Uint32(buf)
	fmt.Printf("Length of the message %d\n", length)

	msgBuf := make([]byte, length)
	_, err = io.ReadFull(r, msgBuf)
	if err != nil {
		return &Message{}, fmt.Errorf("Unable to parse received message: %v\n", err)
	}
	msg := Message{
		Length:  length,
		ID:      int(msgBuf[0]),
		Payload: msgBuf[1:],
	}

	fmt.Printf("Message %v\n", msg)

	return &Message{}, nil
}
