package hippo

import "github.com/pkg/errors"

var (
	ErrorRequiredOption    = errors.New("re	quired option")
	ErrorInvalidTimeFormat = errors.New("invalid time format")
	ErrorInvalidDataFormat = errors.New("invalid data format")
	ErrorDirectoryNotFound = errors.New("directory not found")
)
