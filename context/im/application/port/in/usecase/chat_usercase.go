package usecase

import (
	"context"
)

type ChatUseCase interface {
	UserLogin(ctx context.Context, input UserLoginCommand) (UserLoginOutput, error)
}
