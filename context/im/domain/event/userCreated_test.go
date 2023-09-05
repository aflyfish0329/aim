package event

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewNewUserCreatedEvent_AssignFields(t *testing.T) {
	userId := "userId"
	userName := "userName"
	event := NewUserCreatedEvent(userId, userName)

	assert.Equal(t, userId, event.UserId)
	assert.Equal(t, userName, event.UserName)
}
