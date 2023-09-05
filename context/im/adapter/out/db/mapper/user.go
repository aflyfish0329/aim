package mapper

import (
	"test/context/im/adapter/out/db/model"
	model1 "test/context/im/domain/model"

	"gorm.io/plugin/optimisticlock"
)

func ModelUserToDomainUser(modelUser model.User) model1.User {
	return model1.User{
		Id:       modelUser.Id,
		Username: modelUser.Username,
		Password: modelUser.Password,
		Version:  modelUser.Version.Int64,
	}
}

func DomainUserToModelUser(domainUser model1.User) model.User {
	return model.User{
		Id:       domainUser.Id,
		Username: domainUser.Username,
		Password: domainUser.Password,
		Version: optimisticlock.Version{
			Int64: domainUser.Version,
			Valid: true,
		},
	}
}
