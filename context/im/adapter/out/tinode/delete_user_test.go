//go:build integration

package tinode

import (
	"test/context/im/application/port/out/chat"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_DeleteUser(t *testing.T) {
	s := NewService("localhost:16060")
	_, err := s.DeleteUser(
		chat.DeleteUserReq{
			Token:        "EqdbcdTxa9dvFANlHgABAAEAT7Xs9cg2BuE9vrIigXuEC4J0RPMuhHjwFL46iE8jG60=",
			User:         "usr94Y0w1QvDVM",
			IsHardDelete: true,
		},
	)
	assert.NoError(t, err)
}
