package randomutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandomChars(t *testing.T) {
	assert.Equal(t, true, len(RandomChars(10)) == 10)
	assert.Equal(t, true, len(RandomChars(10, "abc")) == 10)
}

func TestRandomNumber(t *testing.T) {
	assert.Equal(t, true, len(RandomNumber(10)) == 10)
}

func TestRandomInt(t *testing.T) {
	assert.Equal(t, true, RandomInt(10, 20) != 0)
}

func TestRandomInt64(t *testing.T) {
	assert.Equal(t, true, RandomInt64(10, 20) != 0)
}

func TestRandomStr(t *testing.T) {
	assert.Equal(t, true, len(RandomStr(10)) == 10)
}
