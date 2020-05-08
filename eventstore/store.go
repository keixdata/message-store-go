package eventstore

import (
	"strings"
	"time"

	"github.com/google/uuid"
)

// EventStore represents an event store that follows the message db specification.
type EventStore interface {
	Write(msg *Message) (uint64, error)
}

type MessageMetadata interface{}
type MessageData interface{}

type Message struct {
	ID         string
	StreamName string
	Type       string
	Data       MessageData
	Metadata   MessageMetadata
	Time       time.Time
}

func NewMessage(StreamName string, Type string, Data MessageData, Metadata MessageMetadata) *Message {
	return &Message{
		ID:         uuid.New().String(),
		StreamName: StreamName,
		Type:       Type,
		Data:       Data,
		Metadata:   Metadata,
		Time:       time.Now(),
	}
}

func (msg *Message) isCategory() bool {
	return strings.Contains(msg.StreamName, "-")
}
