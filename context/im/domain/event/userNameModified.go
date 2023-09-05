package event

import (
	"test/context/common/domain"
)

const UserNameModifiedEventName = "userNameModified"

func NewUserNameModifiedEvent(userId, userName string) UserNameModified {
	return UserNameModified{
		EventBase: domain.NewEventBase(UserNameModifiedEventName),
		UserId:    userId,
		UserName:  userName,
	}
}

type UserNameModified struct {
	domain.EventBase
	UserId   string `json:"userId"`
	UserName string `json:"userName"`
}
