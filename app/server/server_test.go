package server

import (
	"os"
	"os/signal"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestServer(t *testing.T) {
	err := Run()
	require.NoError(t, err)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)

	osProcess, err := os.FindProcess(os.Getpid())
	require.NoError(t,err)

	err = osProcess.Signal(os.Interrupt)
	require.NoError(t, err)

}