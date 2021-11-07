package business

import "errors"

var (
	ErrDataNotSpec = errors.New("Data Not Spec")

	ErrDataConflict = errors.New("Data Conflict")

	ErrUnauthorized = errors.New("Unauthorized")
)
