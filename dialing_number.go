package api

import (
	"fmt"
	"strconv"
)

const (
	minlenCode          = 2
	maxlenCode          = 5
	minlenNumber        = 9
	maxlenNumber        = 10
	minlenDialingNumber = 13
	maxlenDialingNumber = 14
)

type Diable interface {
	isValidCode(string) bool
	isValidNumber(string) bool
}

type DialingNumber interface {
	AreaCode() string
	Number() string
	DialingNumber() string
}

type number struct {
	AreaCode string
	Number   string
}

type LocalNumber struct {
	*number
}

func (l *LocalNumber) Number() string {
	return l.number.Number
}

func (l *LocalNumber) DialingNumber() string {
	return fmt.Sprintf("%s-%s", l.AreaCode(), l.Number())
}

func (l *LocalNumber) AreaCode() string {
	return l.number.AreaCode
}

type AreaCodeError struct {
	code string
}

func (cerr *AreaCodeError) Error() string {
	return fmt.Sprintf(" Invalid area code length: %s - Expected between %d and %d", cerr.code, minlenCode, maxlenCode)

}

type LocalNumberError struct {
	number string
}

func (derr *LocalNumberError) Error() string {
	return fmt.Sprintf(" Invalid number length: %s - Expected between %d and %d", derr.number, minlenNumber, maxlenNumber)

}

func isValidCode(code string) bool {
	_len := len(code)
	if _, ok := strconv.Atoi(code); ok != nil {
		return false
	}
	return _len >= minlenCode && _len <= maxlenCode
}

func isValidNumber(number string) bool {
	_len := len(number)
	if _, ok := strconv.Atoi(number); ok != nil {
		return false
	}
	return _len >= minlenNumber && _len <= maxlenNumber
}

func isvalidTotalLenght(dialnumber string) bool {
	_len := len(dialnumber)
	return _len >= minlenDialingNumber && _len <= maxlenDialingNumber
}

func NewDialingNumber(code string, num string) (DialingNumber, error) {
	if !isValidCode(code) {
		return nil, &AreaCodeError{code}
	}
	if !isValidNumber(num) {
		return nil, &LocalNumberError{num}
	}
	l := new(LocalNumber)
	l.number = &number{Number: num, AreaCode: code}
	var d DialingNumber = l
	return d, nil
}
