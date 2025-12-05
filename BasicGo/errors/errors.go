package errorsdemo

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

var ErrConfigNotFound = errors.New("config file not found")

type ValidationError struct {
	Field string
	Msg   string
}

func (e ValidationError) Error() string {
	if e.Field == "" {
		return e.Msg
	}
	return fmt.Sprintf("%s: %s", e.Field, e.Msg)
}

func LoadConfig(path string) error {
	if path == "" {
		return fmt.Errorf("open config: %w", ErrConfigNotFound)
	}
	if !strings.HasSuffix(path, ".yaml") {
		return fmt.Errorf("open config %s: %w", path, os.ErrInvalid)
	}
	return nil
}

func ValidateUser(name string, age int) error {
	if strings.TrimSpace(name) == "" {
		return ValidationError{Field: "name", Msg: "cannot be empty"}
	}
	if age < 0 {
		return ValidationError{Field: "age", Msg: "must be positive"}
	}
	return nil
}

func HandleConfig(path string) string {
	err := LoadConfig(path)
	if err == nil {
		return "config loaded"
	}
	switch {
	case errors.Is(err, ErrConfigNotFound):
		return "please provide a config.yaml file"
	case errors.Is(err, os.ErrInvalid):
		return "config must be a .yaml file"
	default:
		return err.Error()
	}
}

func ValidateAndDescribe(name string, age int) string {
	err := ValidateUser(name, age)
	if err == nil {
		return "user valid"
	}
	var v ValidationError
	if errors.As(err, &v) {
		return fmt.Sprintf("field %s invalid: %s", v.Field, v.Msg)
	}
	return "unknown validation error"
}
