package sharedtxclient

import (
	"errors"
)

var (
	ErrClientAlreadyConnected = errors.New("Already connected to SharedTx server")
	ErrClientNotConnected     = errors.New("Client is not connected to SharedTx server")
	ErrCannotConnect = errors.New("Cannot connect to server")
)
