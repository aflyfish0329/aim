package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewEventBase_AssignFields(t *testing.T) {
	eventName := "helloworld"
	base := NewEventBase(eventName)

	assert.NotEmpty(t, base.EventId)
	assert.NotEmpty(t, base.CreateTime)
	assert.Equal(t, eventName, base.EventName)
}
