package api

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	// "gitlab.dev-shelter/libraries/smsc/api"
)

func TestNewSMSC(t *testing.T) {
	c := new(Credential)
	c.Alias = "Alias"
	c.APIKey = "14928_196755354456464a1424ff0fcd562*&07"
	s := NewSMSC(c)
	c2 := s.GetCredential()
	if c != c2 {
		t.Errorf("If you see this then Golang is broken by design")
	}
	s2 := NewSMSC(c)
	if s != s2 {
		t.Errorf("Invalid Singleton!")
	}

}

func TestSMSCStatus(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, `{"code": 200, "message": "hola", "data": {"linea_smsc": "false", "daemon": "false","estado": "true"}}`)

	}))
	defer ts.Close()
	targetUrl = ts.URL
	c := new(Credential)
	c.Alias = "Alias"
	c.APIKey = "14928_196755354456464a1424ff0fcd562"
	s := NewSMSC(c)
	r, err := s.Status()
	if err != nil {
		t.Fatal(err)
	}
	if r.Data.Daemon {
		t.Fatal("Invalid parsing")
	}

}

func TestSMSCSentMessages(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		msg := `{"code":200,"message":"","data":[{"id":"18998689","fecha":"1532721996","fechahora":"2018-07-27T20:06:36Z","mensaje":"#Arrobatech ANTONELA! Su producto CELULAR MOTOROLA G3 esta listo para retirar. Horario: L a V de 9 a 20hs. Sab 9 a 13 y 17 a 20hs.","destinatarios":[{"prefijo":"291","fijo":"4615151","enviado":{"intento":"20","estado":"ok","estado_desc":"Entregado"}}]},{"id":"18998080","fecha":"1532720489","fechahora":"2018-07-27T19:41:29Z","mensaje":"#Arrobatech KARINA! Su producto CELULAR SAMSUNG  J7 PRIME esta listo para retirar. Horario: L a V de 9 a 20hs. Sab 9 a 13 y 17 a 20hs.","destinatarios":[{"prefijo":"291","fijo":"4022761","enviado":{"intento":"20","estado":"ok","estado_desc":"Entregado"}}]}]}`
		io.WriteString(w, msg)
	}))
	targetUrl = ts.URL
	defer ts.Close()
	c := new(Credential)
	c.Alias = "Alias"
	c.APIKey = "14928_196755354456464a1424ff0fcd562"
	s := NewSMSC(c)
	r, err := s.Sent()
	if err != nil {
		t.Fatal(err)
	}
	if r.Code != 200 {
		t.Fatalf("Invalid parsing - Status code received: %d - Expected: %d", r.Code, 200)
	}
	if len(r.Data) != 2 {
		t.Fatal("Invalid parsing in sent messages")
	}

}

func TestSMSCSentMessagesSince(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		msg := `{"code":200,"message":"","data":[{"id":"18998689","fecha":"1532721996","fechahora":"2018-07-27T20:06:36Z","mensaje":"#Arrobatech ANTONELA! Su producto CELULAR MOTOROLA G3 esta listo para retirar. Horario: L a V de 9 a 20hs. Sab 9 a 13 y 17 a 20hs.","destinatarios":[{"prefijo":"291","fijo":"4615151","enviado":{"intento":"20","estado":"ok","estado_desc":"Entregado"}}]},{"id":"18998080","fecha":"1532720489","fechahora":"2018-07-27T19:41:29Z","mensaje":"#Arrobatech KARINA! Su producto CELULAR SAMSUNG  J7 PRIME esta listo para retirar. Horario: L a V de 9 a 20hs. Sab 9 a 13 y 17 a 20hs.","destinatarios":[{"prefijo":"291","fijo":"4022761","enviado":{"intento":"20","estado":"ok","estado_desc":"Entregado"}}]}]}`
		io.WriteString(w, msg)
	}))
	targetUrl = ts.URL
	defer ts.Close()
	c := new(Credential)
	c.Alias = "Alias"
	c.APIKey = "14928_196755354456464a1424ff0fcd562"
	s := NewSMSC(c)
	r, err := s.SentSince("18910095")
	if err != nil {
		t.Fatal(err)
	}
	if r.Code != 200 {
		t.Fatalf("Invalid parsing - Status code received: %d - Expected: %d", r.Code, 200)
	}
	if len(r.Data) != 2 {
		t.Fatal("Invalid parsing in sent messages")
	}
}

func TestSMSCSReceivedMessages(t *testing.T) {
	msg := `{"code":200,"message":"","data":[{"id":"18943242","de":"0","linea":"0","fecha":"1532559120","fechahora":"2018-07-25T22:52:00Z","mensaje":"#Arrobatech Hola papu2"},{"id":"18943241","de":"0","linea":"0","fecha":"1532559119","fechahora":"2018-07-25T22:51:59Z","mensaje":"#Arrobatech Hola papu2"}]}`
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)

		io.WriteString(w, msg)
	}))
	targetUrl = ts.URL
	defer ts.Close()
	c := new(Credential)
	c.Alias = "Alias"
	c.APIKey = "14928_196755354456464a1424ff0fcd562"
	s := NewSMSC(c)
	resp, err := s.Received()
	if err != nil {
		t.Fatal(err)
	}
	if len(resp.Data) != 2 {
		t.Fatal("Invalid parsed response")
	}

}
func TestSMSCReceivedMessagesSince(t *testing.T) {
	msg := `{"code":200,"message":"","data":[{"id":"18943242","de":"0","linea":"0","fecha":"1532559120","fechahora":"2018-07-25T22:52:00Z","mensaje":"#Arrobatech Hola papu2"},{"id":"18943241","de":"0","linea":"0","fecha":"1532559119","fechahora":"2018-07-25T22:51:59Z","mensaje":"#Arrobatech Hola papu2"}]}`
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, msg)
	}))
	targetUrl = ts.URL
	defer ts.Close()
	c := new(Credential)
	c.Alias = "Alias"
	c.APIKey = "14928_196755354456464a1424ff0fcd562"
	s := NewSMSC(c)
	resp, err := s.ReceivedSince("17164713")
	if err != nil {
		t.Fatal(err)
	}
	if len(resp.Data) != 2 {
		t.Fatal("Invalid parsed response")
	}

}

func TestSMSCBalanceInvalidCredentials(t *testing.T) {
	msg := `{"code":401,"message":"Acceso No Autorizado","data":[]}`
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, msg)
	}))
	targetUrl = ts.URL
	defer ts.Close()
	c := new(Credential)
	c.Alias = "Alias"
	c.APIKey = "14928_196755354456464a1424ff0fcd562"
	s := NewSMSC(c)
	bal, err := s.Balance()
	if err != nil {
		t.Fatal(err)
	}
	if bal.Code != 401 {
		t.Fatal("Invalid access not parsed well")
	}
	if bal.Message != "Acceso No Autorizado" {
		t.Fatal("Invalid message parsing")
	}

}

func TestSMSCBalance(t *testing.T) {
	msg := `{"code":200,"message":"","data":{"mensajes":953}}`
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, msg)
	}))
	targetUrl = ts.URL
	defer ts.Close()
	c := new(Credential)
	c.Alias = "Alias"
	c.APIKey = "14928_196755354456464a1424ff0fcd562"
	s := NewSMSC(c)
	bal, err := s.Balance()
	if err != nil {
		t.Fatal(err)
	}
	if bal.Data.Mensajes != 953 {
		t.Fatal("Invalid access not parsed well")
	}

}

func TestSMSCCancelQueue(t *testing.T) {
	msg := `sarasa`
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, msg)
	}))
	targetUrl = ts.URL
	defer ts.Close()
	c := new(Credential)
	c.Alias = "ASDASDASD"
	c.APIKey = "SOY UN AHUACATE"
	c.LineId = "SomeLineId"
	once.Reset()
	s := NewSMSC(c)
	_, err := s.CancelQueue()
	fmt.Println(err)
	if err != nil {
		t.Fatal(err)
	}

}

func TestSMSCCancelQueueInvalidLineId(t *testing.T) {
	once.Reset()
	c := new(Credential)
	c.Alias = "Arrobatech"
	c.APIKey = "14928_1967553541488e4d0ba1424ff0fcd562"
	s := NewSMSC(c)
	_, err := s.CancelQueue()
	if err == nil {
		t.Fatal(err)
	}

}

func TestSMSCSendMessage(t *testing.T) {
	msg := `{"code":200,"message":"Mensaje enviado.","data":{"id":19013234,"sms":1}}`
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, msg)
	}))
	targetUrl = ts.URL
	defer ts.Close()
	c := new(Credential)
	c.Alias = "Arrobatech"
	c.APIKey = "14928_1967553541488e4d0ba1424ff0fcd562"
	s := NewSMSC(c)
	m, _ := NewMessage("mekemkemkekmekme")
	n, _ := NewDialingNumber("011", "1531255999")
	resp, _ := s.Send(m, n)
	if resp.Code != 200 {
		t.Fatal("Invalid code")
	}
}
