package errorsdemo

import (
	"errors"
	"os"
	"testing"
)

func TestHandleConfig(t *testing.T) {
	if msg := HandleConfig(""); msg != "please provide a config.yaml file" {
		t.Fatalf("unexpected message: %s", msg)
	}
	if msg := HandleConfig("config.txt"); msg != "config must be a .yaml file" {
		t.Fatalf("unexpected message: %s", msg)
	}
	if err := LoadConfig("app.yaml"); err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
}

func TestValidateAndDescribe(t *testing.T) {
	msg := ValidateAndDescribe("", 10)
	if msg != "field name invalid: cannot be empty" {
		t.Fatalf("unexpected message: %s", msg)
	}
	msg = ValidateAndDescribe("gopher", -1)
	if msg != "field age invalid: must be positive" {
		t.Fatalf("unexpected message: %s", msg)
	}
	msg = ValidateAndDescribe("gopher", 1)
	if msg != "user valid" {
		t.Fatalf("unexpected message: %s", msg)
	}
}

func TestLoadConfigErrorWrapping(t *testing.T) {
	err := LoadConfig("settings.txt")
	if !errors.Is(err, os.ErrInvalid) {
		t.Fatalf("expected os.ErrInvalid, got %v", err)
	}
}
