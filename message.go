package api

import (
	"fmt"
)

const LEN = 70

type LenError struct {
	Len int
}

func (l *LenError) Error() string {
	return fmt.Sprintf("Invalid Len %d - Max allowed: %d", l.Len, LEN)
}

type EmptyError struct{}

func (e *EmptyError) Error() string {
	return fmt.Sprintf("Empty message not allowed")
}

//Message to be sent to SMSC API
type Message struct {
	content string
	len     int
	time    int
}

func (m *Message) Content() string { return m.content }

func NewMessage(body string) (*Message, error) {
	_len := len(body)
	if _len > LEN {
		return nil, &LenError{_len}
	}
	if _len == 0 {
		return nil, &EmptyError{}
	}
	m := new(Message)
	m.content = body
	m.len = _len
	return m, nil
}
