package osutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCPUInfo(t *testing.T) {
	info, err := GetCPUInfo()
	assert.Equal(t, true, err == nil)
	assert.Equal(t, true, info.CPUModel != "")
}

func TestGetMemInfo(t *testing.T) {
	info, err := GetMemInfo()
	assert.Equal(t, true, err == nil)
	assert.Equal(t, true, info.Total != "")
}

func TestGetDiskInfo(t *testing.T) {
	info, err := GetDiskInfo()
	assert.Equal(t, true, err == nil)
	assert.Equal(t, true, len(info) > 0)
}

func TestGetSysInfo(t *testing.T) {
	info, err := GetSysInfo()
	assert.Equal(t, true, err == nil)
	assert.Equal(t, true, info.Arch != "")
}
