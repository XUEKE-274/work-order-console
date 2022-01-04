package exception

import (
	"errors"
	"work-order-console/domain/enum"
)

func NewException(code enum.ErrorCodeEnum) error {
	return errors.New(string(code))
}

func Code(e error) enum.ErrorCodeEnum {
	text := e.Error()
	return enum.ErrorCodeEnum(text)
}
