package aerror

import "errors"

var (
	ErrGameStarted = errors.New("game started")
	ErrFaliedMarshal = errors.New("failed marshal")
)
