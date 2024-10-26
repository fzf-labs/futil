package uuidutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenUUID(t *testing.T) {
	uuid := GenUUID()
	assert.Equal(t, len(uuid), 36)
}

func TestGenShortID(t *testing.T) {
	shortID, err := GenShortID()
	assert.Equal(t, shortID != "", true)
	assert.Equal(t, err, nil)
}

func TestKSUID(t *testing.T) {
	ksuid := KSUID()
	assert.Equal(t, len(ksuid), 27)
}

func TestKSUIDByTime(t *testing.T) {
	ksuidByTime := KSUIDByTime()
	assert.Equal(t, len(ksuidByTime), 27)
}
