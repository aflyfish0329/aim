package tinode

import (
	"fmt"
	"test/context/im/application/port/out/chat"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_AddGroupUser(t *testing.T) {
	s := NewService("localhost:16060")
	resp, err := s.AddGroupUser(
		chat.AddGroupUserReq{
			Token: "id6cwUjXSHbNQwVlFAABAAEAuKKmIsCmWT0Lf1C1fpIwVGXHbK6zmMVfQ4rRwqPtYj0=",
			Group: "grpG1S5PNJbDwU",
			User:  "usrtqRQpJxO7Jk",
		},
	)
	assert.NoError(t, err)
	fmt.Println(resp)
}
