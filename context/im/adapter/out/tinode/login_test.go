//go:build integration

package tinode

import (
	"fmt"
	"test/context/im/application/port/out/chat"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_UserLogin(t *testing.T) {
	s := NewService("localhost:16060")
	resp, err := s.UserLogin(chat.UserLoginReq{
		Username: "zzzzzzzzzz",
		Password: "zzzzzzzzzz",
	})
	assert.NoError(t, err)
	fmt.Println(resp)
}
