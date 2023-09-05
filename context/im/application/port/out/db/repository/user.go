package repository

import (
	"context"
	"test/context/im/domain/model"
)

type UserRepository interface {
	GetUserByUsername(ctx context.Context, username string) (model.User, error)
}
