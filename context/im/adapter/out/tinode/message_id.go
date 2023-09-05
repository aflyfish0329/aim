package tinode

import (
	"math"
	"math/rand"
	"strconv"
)

func NewMessageID() *MessageID {
	id := int(math.Floor((rand.Float64() * 0xFFFF) + 0xFFFF))
	messageID := MessageID(id)

	return &messageID
}

type MessageID int

func (m *MessageID) Next() *MessageID {
	*m = *m + 1

	return m
}

func (m *MessageID) String() string {
	return strconv.FormatFloat(float64(*m), 'f', 0, 64)
}
