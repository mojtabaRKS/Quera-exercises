package main

import (
	"bytes"
	"make/os_util"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRunSimple(t *testing.T) {
	command := `echo "salam"`
	var b bytes.Buffer
	os_util.Run(command, &b)
	assert.Equal(t, "salam\n", b.String())
}

func TestRunMultipleCommand(t *testing.T) {
	var b bytes.Buffer
	command := `sleep 0.1 && echo "salam"`

	err := os_util.Run(command, &b)

	assert.NoError(t, err)
	assert.Equal(t, "salam\n", b.String())
}

func TestRunNonExistingCommand(t *testing.T) {
	var b bytes.Buffer
	command := `hopefully_non_existing_command_or_this_test_would_fail`

	err := os_util.Run(command, &b)

	assert.Error(t, err)
	assert.Equal(t, "", b.String())
}

func TestRunSyntaxError(t *testing.T) {
	var b bytes.Buffer
	command := `echo "hi" &;!@#$`

	err := os_util.Run(command, &b)

	assert.Error(t, err)
}
