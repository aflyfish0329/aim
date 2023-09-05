package model

import (
	"errors"
	"test/context/common/domain"
	"test/context/im/domain/event"

	uuid "github.com/satori/go.uuid"
)

func NewUser(username, password string) (User, error) {
	user := User{
		Id:       uuid.NewV4().String(),
		Username: username,
		Password: password,
	}

	user.AggregateRoot.AddEvent(event.NewUserCreatedEvent(user.Id, user.Username))

	if err := user.validateSelf(); err != nil {
		return user, err
	}

	return user, nil
}

type User struct {
	domain.AggregateRoot
	Id       string
	Username string
	Password string

	Version int64
}

func (u User) validateSelf() error {
	if u.Id == "" {
		return errors.New("user id cannot be empty")
	}

	if u.Username == "" {
		return errors.New("username cannot be empty")
	}

	if u.Password == "" {
		return errors.New("password cannot be empty")
	}

	return nil
}

func (u *User) ModifyName(name string) error {
	u.Username = name
	u.AggregateRoot.AddEvent(event.NewUserNameModifiedEvent(u.Id, u.Username))

	return u.validateSelf()
}
