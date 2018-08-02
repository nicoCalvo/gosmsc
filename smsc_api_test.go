package api

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestStatus(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, `{"code": 200, "message": "hola", "data": {"linea_smsc": "false", "daemon": "false","estado": "true"}}`)
	}))
	targetUrl = ts.URL

	defer ts.Close()
	c, _ := NewCredential("alias", "14928_19675535414435534s4d0ba1424ff0fcdda32")
	u := NewUrl(c)
	s, err := u.Status()
	if err != nil {
		t.Fatalf("%s", err)
	}
	resp := new(Status)
	json.NewDecoder(s).Decode(resp)
	if resp.Data.LineaSmsc {
		log.Fatal("Invalid state parsing")
	}
}

func TestStatusInvalidServer(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusServiceUnavailable)
		io.WriteString(w, `Server Not Working`)
	}))
	targetUrl = ts.URL
	defer ts.Close()
	c, _ := NewCredential("alias", "14928_1963563sd54148534s4d0ba1424ff0fcdda32")
	u := NewUrl(c)
	_, err := u.Status()
	if err == nil {
		t.Fatal("No error was raised")
	}
	if err.Error() != "Server Not Working" {
		t.Fatal("Invalid Message")
	}

}

func TestParseResponse(t *testing.T) {
	m := []byte(`{"code": 200, "message": "hola", "data": {"linea_smsc": "false", "daemon": "false","estado": "true"}}`)
	resp := new(Status)
	reader := bytes.NewReader(m)
	json.NewDecoder(reader).Decode(resp)
	if resp.Data.LineaSmsc {
		t.Fatal("Invalid parsing")
	}
}

func TestSent(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, `Something wrong`)
	}))
	targetUrl = ts.URL
	defer ts.Close()
	c, _ := NewCredential("Alias", "14928_196755354148657530ba1424ff0fcd562")
	u := NewUrl(c)
	resp, err := u.Sent()
	if err != nil {
		t.Fatalf("%s", err)
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp)
	newStr := buf.String()
	if newStr != "Something wrong" {
		t.Fatal("Invalid parse")
	}
}

func TestUrlReceived(t *testing.T) {
	msg := `"code":200,"message":"","data":[{"id":"18943242","de":"0","linea":"0","fecha":"1532559120","fechahora":"2018-07-25T22:52:00Z","mensaje":"#Arrobatech Hola papu2"},{"id":"18943241","de":"0","linea":"0","fecha":"1532559119","fechahora":"2018-07-25T22:51:59Z","mensaje":"#Arrobatech Hola papu2"}`
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)

		io.WriteString(w, msg)
	}))
	targetUrl = ts.URL
	defer ts.Close()

	c, _ := NewCredential("Alias", "14928_1967553541488sdfre43f0fcd562")
	u := NewUrl(c)
	r, err := u.Received()
	if err != nil {
		t.Fatal(err)
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(r)
	newStr := buf.String()
	if newStr != msg {
		t.Fatal("Invalid return")
	}
}
