package shared

import (
	"errors"
)

var ErrInvalidURL = errors.New("URL provided is not valid")

type Entry struct {
	URL string
}
