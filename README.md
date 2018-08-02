# GoSMSC

## Intent

This library provides simple communcation with [SMSC](https://www.smsc.com.ar/usuario/iniciar/)

## Usage

1 Create a credential with your alias and api key
```go
    c := new(api.Credential)
    c.Alias = "MyAlias"
    c.APIKey = "14928_19675488er559c4d0ba1424ff0fcd562"
```
2  Create an instance of SMSC passing the credential as param

```go
    s := api.NewSMSC(c)
```

3 Create a message to be sent with the content, and a dialing number

```go
    m, err := api.NewMessage("Greetings from a gopher friend")
    if err != nil{
        fmt.Prinln(err)
    }	
    n, err := api.NewDialingNumber("011", "153256689")
    if err != nil {
        fmt.Println(err)
    }
```

4- Send the message

```go
	resp, _ := s.Send(m, n)
	fmt.Println(resp)
```

## Full example
```go
import (
	"fmt"

	smsc "github.com/nicoCalvo/gosmsc"
)

func main() {
	c, err := smsc.NewCredential("MyAlias", "14928_19675535410b24ff0fcd562")
	if err != nil {
		fmt.Println(err)
	}
	s := smsc.NewSMSC(c)

	n, err := smsc.NewDialingNumber("011", "153256689")
	m, err := smsc.NewMessage("Greetings from a gopher friend")
	resp, err := s.Send(m, n)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp)
}
```
## Other methods availables:

Check balance

```go
 s := api.NewSMSC(c)
 b, err := s.Balance()

```

Get sent messages

```go
 s := api.NewSMSC(c)
 b, err := s.Sent()
```

Get sent messages starting from a given id
```go
 s := api.NewSMSC(c)
 b, err := s.s.SentSince("17164713")
```


Get received messages:
```go
 s := api.NewSMSC(c)
 b, err := s.s.Received()
```

Get received messages starting from a given id
```go
 s := api.NewSMSC(c)
 b, err := s.s.ReceivedSince("17164713")
```

Cancel queue
```go
 s := api.NewSMSC(c)
 b, err := s.CancelQueue()
```

Check status
```go
 s := api.NewSMSC(c)
 b, err := s.Status()
```
