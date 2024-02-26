package strutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandSeed(t *testing.T) {
	assert.Equal(t, true, RandSeed().Int() != 0)
}

func TestRandom(t *testing.T) {
	assert.Equal(t, true, len(Random(10)) == 10)
}

func TestRandomChars(t *testing.T) {
	assert.Equal(t, true, len(RandomChars(10)) == 10)
	assert.Equal(t, true, len(RandomChars(10, "abc")) == 10)
}

func TestRandomNumber(t *testing.T) {
	assert.Equal(t, true, len(RandomNumber(10)) == 10)
}
