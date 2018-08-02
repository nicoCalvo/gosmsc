package api_test

import (
	"strings"
	"testing"

	"gitlab.dev-shelter/libraries/smsc/api"
)

func TestNewInvalidLenMessage(t *testing.T) {
	msg := strings.Repeat("#", (api.LEN + 1))
	_, err := api.NewMessage(msg)
	if _, ok := err.(*api.LenError); !ok {
		t.Errorf("Validation error - Limit length not caught")
	}

}

func TestNewValidLenMessage(t *testing.T) {
	msg := strings.Repeat("#", api.LEN)
	_len := len(msg)
	_, err := api.NewMessage(msg)
	if _, ok := err.(*api.LenError); ok {
		t.Errorf("Validation error, Limit len: %d, message len: %d.", api.LEN, _len)
	}

}
func TestNewEmptyMessage(t *testing.T) {
	msg := ""
	_, err := api.NewMessage(msg)
	if _, ok := err.(*api.EmptyError); !ok {
		t.Errorf("Validation error, Empty message not caugth")
	}

}
