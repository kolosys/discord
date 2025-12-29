package state

import "errors"

var (
	ErrNotFound         = errors.New("state: entity not found")
	ErrDisabled         = errors.New("state: cache disabled")
	ErrRESTNotAvailable = errors.New("state: REST client not available")
	ErrNotReady         = errors.New("state: not ready")
)
