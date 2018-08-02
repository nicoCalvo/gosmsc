package api

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	Version = "0.3"
	Timeout = 5
)

var targetUrl = "https://www.smsc.com.ar/api/0.3"

type Url struct {
	*Credential
	BaseUrl string
}

func NewUrl(c *Credential) *Url {
	u := new(Url)
	u.Credential = c
	return u
}

type param struct {
	Name  string
	Value string
}

func (u *Url) getQuery(params ...*param) *http.Request {
	req, _ := http.NewRequest("GET", targetUrl, nil)
	q := req.URL.Query()
	q.Add("alias", u.Credential.Alias)
	q.Add("apikey", u.Credential.APIKey)
	for _, param := range params {
		q.Add(param.Name, param.Value)
	}
	req.URL.RawQuery = q.Encode()
	return req
}
func (u *Url) sendToAPI(req *http.Request) (io.ReadCloser, error) {
	timeout := time.Duration(5 * time.Second)
	client := &http.Client{Timeout: timeout}
	r, err := client.Get(req.URL.String())
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		b, _ := ioutil.ReadAll(r.Body)
		return nil, errors.New(string(b))
	}
	return r.Body, nil
}

func (u *Url) Status() (io.ReadCloser, error) {
	p := &param{Name: "cmd", Value: "estado"}
	req := u.getQuery(p)
	return u.sendToAPI(req)

}

func (u *Url) Sent() (io.ReadCloser, error) {
	p := &param{Name: "cmd", Value: "enviados"}
	req := u.getQuery(p)
	return u.sendToAPI(req)
}
func (u *Url) SentSince(id string) (io.ReadCloser, error) {
	p1 := &param{Name: "cmd", Value: "enviados"}
	p2 := &param{Name: "ultimoid", Value: id}
	req := u.getQuery(p1, p2)
	return u.sendToAPI(req)
}

func (u *Url) Received() (io.ReadCloser, error) {
	p := &param{Name: "cmd", Value: "recibidos"}
	req := u.getQuery(p)
	return u.sendToAPI(req)
}
func (u *Url) ReceivedSince(id string) (io.ReadCloser, error) {
	p1 := &param{Name: "cmd", Value: "recibidos"}
	p2 := &param{Name: "ultimoid", Value: id}
	req := u.getQuery(p1, p2)
	return u.sendToAPI(req)
}

func (u *Url) Balance() (io.ReadCloser, error) {
	p1 := &param{Name: "cmd", Value: "saldo"}
	req := u.getQuery(p1)
	return u.sendToAPI(req)
}

func (u *Url) CancelQueue(lineid string) (io.ReadCloser, error) {
	fmt.Println(lineid)
	p1 := &param{Name: "cmd", Value: "cancelqueue"}
	p2 := &param{Name: "lineid", Value: lineid}
	req := u.getQuery(p1, p2)
	return u.sendToAPI(req)
}

func (u *Url) Send(m *Message, n DialingNumber) (io.ReadCloser, error) {
	p1 := &param{Name: "cmd", Value: "enviar"}
	p2 := &param{Name: "num", Value: n.DialingNumber()}
	p3 := &param{Name: "msj", Value: m.Content()}
	req := u.getQuery(p1, p2, p3)
	return u.sendToAPI(req)
}
