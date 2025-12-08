package interfacedemo

import (
	"errors"
	"fmt"
)

// OpError 演示自定义错误类型实现 error 接口。
type OpError struct {
	Op   string
	Code int
	Msg  string
}

func (e OpError) Error() string {
	return fmt.Sprintf("%s failed(%d): %s", e.Op, e.Code, e.Msg)
}

var ErrTemporary = errors.New("temporary error")

// WrapAsTemporary 用 %w 包装 ErrTemporary，保留错误链。
func WrapAsTemporary(op string) error {
	return fmt.Errorf("%s: %w", op, ErrTemporary)
}

