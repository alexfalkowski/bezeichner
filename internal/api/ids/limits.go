package ids

import "github.com/alexfalkowski/go-service/v2/errors"

const (
	maxGenerateCount = 1000
	maxMapIDs        = 1000
)

// ErrInvalidArgument indicates the caller supplied an invalid request.
//
// In this package it is returned when a request exceeds the configured limits,
// for example an excessive generate count or too many IDs to map.
//
// Use IsInvalidArgument to classify errors without relying on error strings.
var ErrInvalidArgument = errors.New("invalid argument")

// IsInvalidArgument reports whether err is (or wraps) ErrInvalidArgument.
func IsInvalidArgument(err error) bool {
	return errors.Is(err, ErrInvalidArgument)
}
