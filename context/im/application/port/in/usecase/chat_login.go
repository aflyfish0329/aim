package usecase

import "errors"

func NewUserLoginCommand(username, password string) (UserLoginCommand, error) {
	command := UserLoginCommand{
		Username: username,
		Password: password,
	}

	if err := command.validateSelf(); err != nil {
		return command, err
	}

	return command, nil
}

type UserLoginCommand struct {
	Username string
	Password string
}

func (c UserLoginCommand) validateSelf() error {
	if c.Username == "" {
		return errors.New("username cannot be empty")
	}

	if c.Password == "" {
		return errors.New("password cannot be empty")
	}

	return nil
}

type UserLoginOutput struct {
	Token string
}
