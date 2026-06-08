package interfacedemo

import (
	"errors"
	"fmt"
)

type OpError struct {
	Op   string
	Code int
	Msg  string
}

func (e OpError) Error() string {
	return fmt.Sprintf("%s failed(%d): %s", e.Op, e.Code, e.Msg)
}

var ErrTemporary = errors.New("temporary error")

func WrapAsTemporary(op string) error {
	return fmt.Errorf("%s: %w", op, ErrTemporary)
}
