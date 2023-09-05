package event

import "test/context/common/domain"

const UserCreateEventName = "UserCreatedEvent"

func NewUserCreatedEvent(userId, userName string) UserCreatedEvent {
	return UserCreatedEvent{
		EventBase: domain.NewEventBase(UserCreateEventName),
		UserId:    userId,
		UserName:  userName,
	}
}

type UserCreatedEvent struct {
	domain.EventBase
	UserId   string `json:"userId"`
	UserName string `json:"userName"`
}
