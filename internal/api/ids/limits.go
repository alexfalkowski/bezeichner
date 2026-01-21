package ids

import "github.com/alexfalkowski/go-service/v2/errors"

const (
	maxGenerateCount = 1000
	maxMapIDs        = 1000
)

// ErrInvalidArgument is an error for invalid argument.
var ErrInvalidArgument = errors.New("invalid argument")

// IsInvalidArgument checks if the error is an invalid argument error.
func IsInvalidArgument(err error) bool {
	return errors.Is(err, ErrInvalidArgument)
}
