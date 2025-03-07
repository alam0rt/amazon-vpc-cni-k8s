package utils

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	defaultPathEnv = "/host/opt/cni/bin"
	defaultBoolEnv = false

	envPath = "Path"
	envBool = "Bool"
)

// Validate that GetBoolAsStringEnvVar runs against acceptable format input without error
func TestGetBoolAsStringEnvVar(t *testing.T) {
	// Test environment flag variable not set
	tmp := GetBoolAsStringEnvVar(envBool, defaultBoolEnv)
	assert.Equal(t, tmp, defaultBoolEnv)

	// Test basic Boolean as string set with acceptable format
	os.Setenv(envBool, "True")
	tmp = GetBoolAsStringEnvVar(envBool, defaultBoolEnv)
	assert.Equal(t, tmp, true)

	// Test basic Boolean as string set with unacceptable format
	os.Setenv(envBool, "TrUe")
	defer os.Unsetenv(envBool)
	tmp = GetBoolAsStringEnvVar(envBool, defaultBoolEnv)
	assert.Equal(t, tmp, defaultBoolEnv)
}

// Validate that GetEnv runs without error against environment variable with type other than boolean as string
func TestGetEnv(t *testing.T) {
	// Test environment flag variable not set
	tmp := GetEnv(envPath, defaultPathEnv)
	assert.Equal(t, tmp, defaultPathEnv)

	// Test environment flag variable set
	os.Setenv(envPath, "/host/opt/cni/bin/test")
	defer os.Unsetenv(envPath)
	tmp = GetEnv(envPath, defaultPathEnv)
	assert.Equal(t, tmp, "/host/opt/cni/bin/test")
}
