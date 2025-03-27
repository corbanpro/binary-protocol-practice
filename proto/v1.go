package proto

import (
	"bytes"
	"encoding/binary"
	"time"
)

func encodeV1(c Chat) *bytes.Buffer {
	version := 1
	b := new(bytes.Buffer)

	// version
	binary.Write(b, binary.BigEndian, uint16(version))

	// timestamp
	binary.Write(b, binary.BigEndian, uint64(c.Timestamp.Unix()))

	// userId length
	binary.Write(b, binary.BigEndian, uint32(len(c.UserId)))

	// userId
	b.Write([]byte(c.UserId))

	// message length
	binary.Write(b, binary.BigEndian, uint32(len(c.Message)))

	// message
	b.Write([]byte(c.Message))
	return b
}

func decodeV1(b bytes.Buffer) (chat Chat, err error) {
	defer func() {
		if r := recover(); r != nil {
			chat = Chat{}
			err = r.(error)
		}
	}()

	chat.Timestamp = time.Unix(int64(binary.BigEndian.Uint64(b.Next(8))), 0)

	// userId length
	userIdLen := int(binary.BigEndian.Uint32(b.Next(4)))

	// userId
	chat.UserId = string(b.Next(userIdLen))

	// message length
	messageLen := int(binary.BigEndian.Uint32(b.Next(4)))

	// message
	chat.Message = string(b.Next(messageLen))

	return chat, nil
}
