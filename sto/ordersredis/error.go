package ordersredis

import (
	"errors"

	"github.com/xh3b4sd/tracer"
)

var alreadyExistsError = &tracer.Error{
	Kind: "alreadyExistsError",
}

func IsAlreadyExists(err error) bool {
	return errors.Is(err, alreadyExistsError)
}

var invalidConfigError = &tracer.Error{
	Kind: "invalidConfigError",
}

func IsInvalidConfig(err error) bool {
	return errors.Is(err, invalidConfigError)
}

var notFoundError = &tracer.Error{
	Kind: "notFoundError",
}

func IsNotFound(err error) bool {
	return errors.Is(err, notFoundError)
}
