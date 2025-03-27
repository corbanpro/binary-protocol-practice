package proto

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"time"
)

type Chat struct {
	Timestamp time.Time
	UserId    string
	Message   string
}

func (c Chat) Print() {
	fmt.Println()
	fmt.Println("Timestamp:", c.Timestamp)
	fmt.Println("UserId:", c.UserId)
	fmt.Println("Message:", c.Message)
	fmt.Println()
}

func NewChat(userId, message string) Chat {
	return Chat{
		Timestamp: time.Now(),
		UserId:    userId,
		Message:   message,
	}
}

func Encode(c Chat) *bytes.Buffer {
	return encodeV2(c)
}

func Decode(b bytes.Buffer) (Chat, error) {
	if v := binary.BigEndian.Uint16(b.Next(2)); v != 2 {
		return Chat{}, fmt.Errorf("unknown version: %d", v)
	}

	return decodeV2(b)
}
