package database

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSaveMsg_NormalCase(t *testing.T) {
	db := getDBInstance(nil)
	msg := []byte(`Hello World`)

	id, err := SaveMsg(db, string(msg), "user1", "user2")
	assert.NoError(t, err)
	print(id)
	UpdateStatus(nil, id, 1)
}
