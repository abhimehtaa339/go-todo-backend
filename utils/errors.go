package utils

import "errors"

func BindError(err string) error {
	return errors.New(err)
}
