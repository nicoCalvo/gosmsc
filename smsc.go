// GoSMSC provides easy integration with [SMSC](https://www.smsc.com.ar/usuario/iniciar/)
// to send sms to any local number within Argentina.
// From sending a message to get account of current balance, every action available in
// smsc.com.ar is replicated for simple usage in GoSMSC
package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/matryer/resync"
)

type Credential struct {
	Alias  string
	APIKey string
	LineId string
}

func NewCredential(alias string, apikey string) (*Credential, error) {
	return &Credential{Alias: alias, APIKey: apikey}, nil
}

type SMSC struct {
	credential *Credential
	url        *Url
}

func (s *SMSC) Credential() *Credential {
	return s.credential
}

func (s *SMSC) GetCredential() *Credential { return s.credential }

var instance *SMSC
var once resync.Once

type Status struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		LineaSmsc bool `json:"linea_smsc,omitempty"`
		Daemon    bool `json:"daemon,omitempty"`
		Estado    bool `json:"estado,omitempty"`
	} `json:"data"`
}

type SentStatus struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		ID  int `json:"id,omitempty"`
		SMS int `json:"sms,omitempty"`
	} `json:"data"`
}

func (s *Status) String() string {
	return fmt.Sprintf("Code: %d  -  Message: %s  - Linea %t  - Estado: %t - Daemon: %t", s.Code, s.Message, s.Data.LineaSmsc, s.Data.Estado, s.Data.Daemon)
}

func (s *SentStatus) String() string {
	return fmt.Sprintf("Code: %d  -  Message: %s  - Id %d  - SMS: %d", s.Code, s.Message, s.Data.ID, s.Data.SMS)
}

type SentMessages struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    []struct {
		ID            string    `json:"id"`
		Fecha         string    `json:"fecha"`
		Fechahora     time.Time `json:"fechahora"`
		Mensaje       string    `json:"mensaje"`
		Destinatarios []struct {
			Prefijo string `json:"prefijo"`
			Fijo    string `json:"fijo"`
			Enviado struct {
				Intento    string `json:"intento"`
				Estado     string `json:"estado"`
				EstadoDesc string `json:"estado_desc"`
			} `json:"enviado"`
		} `json:"destinatarios"`
	} `json:"data"`
}

type ReceivedMessages struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    []struct {
		ID        string    `json:"id"`
		De        string    `json:"de"`
		Linea     string    `json:"linea"`
		Fecha     string    `json:"fecha"`
		Fechahora time.Time `json:"fechahora"`
		Mensaje   string    `json:"mensaje"`
	} `json:"data"`
}

type Balance struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Mensajes int `json:"mensajes"`
	} `json:"data"`
}

func NewSMSC(c *Credential) *SMSC {
	once.Do(func() {
		instance = new(SMSC)
		instance.credential = c
		instance.url = NewUrl(c)
	})
	return instance
}

func (s *SMSC) Status() (*Status, error) {
	body, err := s.url.Status()
	if err != nil {
		return nil, err
	}
	defer body.Close()
	resp := new(Status)
	json.NewDecoder(body).Decode(resp)
	return resp, nil
}

func (s *SMSC) Sent() (*SentMessages, error) {
	body, err := s.url.Sent()
	if err != nil {
		return nil, err
	}
	resp := new(SentMessages)
	json.NewDecoder(body).Decode(resp)
	return resp, nil
}

func (s *SMSC) SentSince(id string) (*SentMessages, error) {
	body, err := s.url.SentSince(id)
	if err != nil {
		return nil, err
	}
	resp := new(SentMessages)
	json.NewDecoder(body).Decode(resp)
	return resp, nil
}

func (s *SMSC) Received() (*ReceivedMessages, error) {
	body, err := s.url.Received()
	if err != nil {
		return nil, err
	}
	resp := new(ReceivedMessages)
	json.NewDecoder(body).Decode(resp)
	return resp, nil
}

func (s *SMSC) ReceivedSince(id string) (*ReceivedMessages, error) {
	body, err := s.url.ReceivedSince(id)
	if err != nil {
		return nil, err
	}
	resp := new(ReceivedMessages)
	json.NewDecoder(body).Decode(resp)
	return resp, nil
}

func (s *SMSC) Balance() (*Balance, error) {
	body, err := s.url.Balance()
	if err != nil {
		return nil, err
	}

	resp := new(Balance)
	json.NewDecoder(body).Decode(resp)
	return resp, nil

}

func (s *SMSC) CancelQueue() (bool, error) {
	if len(s.credential.LineId) == 0 {
		return false, errors.New("Invalid dedicated line id")
	}
	body, err := s.url.CancelQueue(s.credential.LineId)
	if err != nil {
		return false, err
	}
	resp := new(Status)
	json.NewDecoder(body).Decode(resp)
	return resp.Code == 200, nil

}

func (s *SMSC) Send(m *Message, n DialingNumber) (*SentStatus, error) {
	body, err := s.url.Send(m, n)
	if err != nil {
		return nil, err
	}
	resp := new(SentStatus)
	json.NewDecoder(body).Decode(resp)
	return resp, err
}
