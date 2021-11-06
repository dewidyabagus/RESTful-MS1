package validator

import (
	"sync"

	validator "github.com/go-playground/validator/v10"
)

var mtx = &sync.Mutex{}
var validate *validator.Validate

func GetValidator() *validator.Validate {
	mtx.Lock()
	defer mtx.Unlock()

	if validate == nil {
		validate = validator.New()
	}

	return validate
}
