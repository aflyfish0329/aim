package service

import (
	"context"
	"test/context/im/application/port/in/usecase"
	"test/context/im/application/port/out/chat"

	"github.com/pkg/errors"
)

func (s ChatService) UserLogin(ctx context.Context, input usecase.UserLoginCommand) (usecase.UserLoginOutput, error) {
	user, err := s.userRepository.GetUserByUsername(ctx, input.Username)
	if err != nil {
		return usecase.UserLoginOutput{}, errors.Wrapf(err, "ChatService: Login: user %s not found", input.Username)
	}

	if user.Password != input.Password {
		return usecase.UserLoginOutput{}, errors.Wrapf(err, "ChatService: Login: user %s password incorrect", input.Username)
	}

	resp, err := s.chatService.UserLogin(
		chat.UserLoginReq{
			Username: input.Username,
			Password: input.Password,
		})
	if err != nil {
		return usecase.UserLoginOutput{}, err
	}

	output := usecase.UserLoginOutput{
		Token: resp.Token,
	}

	return output, nil
}
