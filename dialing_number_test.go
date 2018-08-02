package api_test

import (
	"strings"
	"testing"

	"gitlab.dev-shelter/libraries/smsc/api"
)

func TestMinCode(t *testing.T) {
	c := "1"
	n := "155123456"
	_, err := api.NewDialingNumber(c, n)
	if _, ok := err.(*api.AreaCodeError); !ok {
		t.Fatal("Invalid Min area code not caught")
	}
}

func TestInvalidTypeCode(t *testing.T) {
	c := "12k3"
	n := "155123456"
	_, err := api.NewDialingNumber(c, n)
	if _, ok := err.(*api.AreaCodeError); !ok {
		t.Fatal("Invalid type area code not caught")

	}
}

func TestMaxCode(t *testing.T) {
	c := strings.Repeat("1", 6)
	n := "155123456"
	_, err := api.NewDialingNumber(c, n)
	if _, ok := err.(*api.AreaCodeError); !ok {
		t.Fatal("Invalid Max area code not caught")
	}
}

func TestMinNumber(t *testing.T) {
	c := "123"
	n := "15512345"
	_, err := api.NewDialingNumber(c, n)
	if _, ok := err.(*api.LocalNumberError); !ok {
		t.Fatal("Invalid Min number code not caught")
	}
}

func TestInvalidTypeNumber(t *testing.T) {
	c := "123"
	n := "155123a56"
	_, err := api.NewDialingNumber(c, n)
	if _, ok := err.(*api.LocalNumberError); !ok {
		t.Fatal("Invalid type number code not caught")

	}
}

func TestMaxNumber(t *testing.T) {
	c := strings.Repeat("1", 4)
	n := "155123456123"
	_, err := api.NewDialingNumber(c, n)
	if _, ok := err.(*api.LocalNumberError); !ok {
		t.Fatal("Invalid Max number code not caught")
	}
}

func TestValidDialingNumber(t *testing.T) {
	_, err := api.NewDialingNumber("0291", "155123456")
	if err != nil {
		t.Fatalf("Error creating dialing number1: %d", err)
	}
	_, err = api.NewDialingNumber("011", "1551234567")
	if err != nil {
		t.Fatalf("Error creating dialing number2: %d", err)
	}
	_, err = api.NewDialingNumber("02932", "155123456")
	if err != nil {
		t.Fatalf("Error creating dialing number3: %d", err)
	}

}

func TestGetDialingNumber(t *testing.T) {
	_, err := api.NewDialingNumber("0291", "155123456")
	if err != nil {
		t.Fatalf("Error creating dialing number1: %d", err)
	}

}
