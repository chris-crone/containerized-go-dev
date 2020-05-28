package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEcho(t *testing.T) {
	// Test happy path
	err := echo([]string{"bin-name", "hello", "world!"})
	require.NoError(t, err)
}

func TestEchoErrorNoArgs(t *testing.T) {
	// Test empty arguments
	err := echo([]string{})
	require.Error(t, err)
}
