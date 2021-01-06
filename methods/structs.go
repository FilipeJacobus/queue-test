package methods

import "errors"

type StringServiceInt interface {
	Persist(string) (string, error)
}

type StringService struct{}

var ErrEmpty = errors.New("empty string")

type PersistRequest struct {
	CliID     string  `json:"id"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type PersistResponse struct {
	Message string `json:"message"`
	Err     string `json:"err,omitempty"` // errors don't define JSON marshaling
}
