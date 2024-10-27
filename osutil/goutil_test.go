package osutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGoVersion(t *testing.T) {
	assert.Equal(t, true, GoVersion() != "")
}
