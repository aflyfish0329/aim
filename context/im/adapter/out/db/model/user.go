package model

import "gorm.io/plugin/optimisticlock"

type User struct {
	Id       string `gorm:"column:id;type:uuid;primaryKey"`
	Username string `gorm:"column:username;type:varchar(64)"`
	Password string `gorm:"column:password;type:varchar(64)"`
	Version  optimisticlock.Version
}
