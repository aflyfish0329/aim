//go:build integration

package tinode

import (
	"fmt"
	"test/context/im/application/port/out/chat"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CreateUser(t *testing.T) {
	s := NewService("localhost:16060")
	resp, err := s.CreateUser(
		chat.CreateUserReq{
			Username: "aaaaaaaaaaaaa",
			Password: "aaaaaaaaaaaaa",
			Email:    "aaaaaaaaaaaaa@example.com",
		},
	)
	assert.NoError(t, err)
	fmt.Println(resp)
}
