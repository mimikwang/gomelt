package gomelt

import "errors"

var (
	ErrInvalidBase     = errors.New("invalid DNA base")
	ErrInvalidDivalent = errors.New("divalent should be greater or equal to 0")
	ErrInvalidDntp     = errors.New("dntp should be greater or equal to 0")
)
