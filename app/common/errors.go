package common

import "errors"

var (
	ErrScanValueNil         = errors.New("scan value is nil")
	ErrScanValueIsNotString = errors.New("scan value is not a string")
)
