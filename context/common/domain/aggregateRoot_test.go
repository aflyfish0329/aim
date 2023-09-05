package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type FakeEvent struct{}

func (e FakeEvent) GetEventName() string {
	return "helloworld"
}

func Test_AggregateRoot_AddEvent(t *testing.T) {
	root := AggregateRoot{}
	root.AddEvent(FakeEvent{})

	assert.Equal(t, 1, len(root.Events))
}
