package tinode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewMessageID(t *testing.T) {
	id := NewMessageID()

	assert.NotZero(t, id)
}

func Test_MessageID_Next(t *testing.T) {
	id := NewMessageID()
	bid := *id
	nid := id.Next()

	assert.Equal(t, *id, bid+1)
	assert.Equal(t, *nid, bid+1)
}

func Test_MessageID_String(t *testing.T) {
	id := NewMessageID()
	sid := fmt.Sprintf("%d", *id)

	assert.Equal(t, sid, id.String())
}
