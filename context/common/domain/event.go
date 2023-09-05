package domain

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type Event interface {
	GetEventName() string
}

func NewEventBase(eventName string) EventBase {
	return EventBase{
		EventId:    uuid.NewV4().String(),
		CreateTime: time.Now(),

		EventName: eventName,
	}
}

type EventBase struct {
	EventId    string    `json:"eventId"`
	CreateTime time.Time `json:"createTime"`

	EventName string `json:"eventName"`
}

func (e EventBase) GetEventName() string {
	return e.EventName
}
