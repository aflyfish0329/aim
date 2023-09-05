package model

import (
	"test/context/im/domain/event"
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

func Test_NewUser_Succeed(t *testing.T) {
	username := "allen"
	password := "password"
	user, err := NewUser(username, password)

	assert.NoError(t, err)
	assert.NotEmpty(t, user.Id)
	assert.Equal(t, username, user.Username)
	assert.Equal(t, password, user.Password)
	assert.Equal(t, 1, len(user.Events))
	assert.Equal(t, event.UserCreateEventName, user.Events[0].GetEventName())
}

func Test_NewUser_Validate_Failed(t *testing.T) {
	_, err := NewUser("", "")
	assert.ErrorContains(t, err, "username")

	_, err = NewUser("allen", "")
	assert.ErrorContains(t, err, "password")
}

func Test_User_ValidateSelf_Succeed(t *testing.T) {
	user := User{
		Id:       uuid.NewV4().String(),
		Username: "allen",
		Password: "password",
	}

	err := user.validateSelf()
	assert.NoError(t, err)
}

func Test_User_ModifyName_Succeed(t *testing.T) {
	name := "golang"

	user := User{
		Id:       uuid.NewV4().String(),
		Username: "allen",
		Password: "password",
	}

	err := user.ModifyName(name)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(user.Events))
	assert.Equal(t, event.UserNameModifiedEventName, user.Events[0].GetEventName())
}

func Test_User_ModifyName_NameEmpty_Failed(t *testing.T) {
	name := ""

	user := User{
		Id:       uuid.NewV4().String(),
		Username: "allen",
	}

	err := user.ModifyName(name)
	assert.ErrorContains(t, err, "name")
}

func Test_User_ValidateSelf_NameEmpty_Failed(t *testing.T) {
	user := User{
		Id:       "id1",
		Username: "",
	}

	err := user.validateSelf()
	assert.ErrorContains(t, err, "name")
}

func Test_User_ValidateSelf_IdEmpty_Failed(t *testing.T) {
	user := User{
		Id: "",
	}

	err := user.validateSelf()
	assert.ErrorContains(t, err, "id")
}
