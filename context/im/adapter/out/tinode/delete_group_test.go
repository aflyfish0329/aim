//go:build integration

package tinode

import (
	"fmt"
	"test/context/im/application/port/out/chat"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_DeleteGroup(t *testing.T) {
	s := NewService("localhost:16060")
	resp, err := s.DeleteGroup(
		chat.DeleteGroupReq{
			Token: "id6cwUjXSHbNQwVlFAABAAEAuKKmIsCmWT0Lf1C1fpIwVGXHbK6zmMVfQ4rRwqPtYj0=",
			Group: "grplApGTyf3_VU",
		},
	)
	assert.NoError(t, err)
	fmt.Println(resp)
}
