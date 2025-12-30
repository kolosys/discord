package server

import "errors"

var (
	ErrAlreadyRunning = errors.New("server: manager already running")
	ErrNotRunning     = errors.New("server: manager not running")
	ErrNoCallbacks    = errors.New("server: connect and stop callbacks required")
)
