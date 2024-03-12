package osutil

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGoVersion(t *testing.T) {
	assert.Equal(t, true, GoVersion() != "")
}
