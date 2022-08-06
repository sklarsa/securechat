package main

import (
	"bytes"
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"

	shell "github.com/ipfs/go-ipfs-api"
)

func GetMessage(sh *shell.Shell, cid string) (Message, string, error) {
	msg := &InternalMessage{}
	rdr, err := sh.Cat(cid)
	if err != nil {
		return Message{}, "", err
	}
	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(rdr)
	if err != nil {
		return Message{}, "", err
	}
	err = json.Unmarshal(buf.Bytes(), msg)
	if err != nil {
		return Message{}, "", err
	}
	return msg.Message, msg.ParentCid, nil
}

type User struct {
	Id             string
	PublicKeyBytes []byte
}

func (u *User) PublicKey() (*rsa.PublicKey, error) {
	return x509.ParsePKCS1PublicKey(u.PublicKeyBytes)
}

type Message struct {
	User string `json:"user"`
	Text string `json:"text"`
}

type InternalMessage struct {
	Message   `json:"message"`
	ParentCid string `json:"parentCid"`
}

type Channel struct {
	sh      *shell.Shell
	Name    string
	TailCid string
}

func (c *Channel) Read(sh *shell.Shell, ch chan Message) error {
	cid := c.TailCid
	for cid != "" {
		var msg Message
		var err error
		msg, cid, err = GetMessage(sh, cid)
		if err != nil {
			return err
		}
		ch <- msg
	}
	return nil
}

func NewChannel(sh *shell.Shell, name string) Channel {
	return Channel{
		sh:      sh,
		Name:    name,
		TailCid: "",
	}
}

func (c *Channel) Write(msg Message) error {
	iMsg := InternalMessage{
		Message:   msg,
		ParentCid: c.TailCid,
	}

	buf, err := json.Marshal(iMsg)
	if err != nil {
		return err
	}
	rdr := bytes.NewReader(buf)
	cid, err := c.sh.Add(rdr)
	if err != nil {
		return err
	}
	c.TailCid = cid
	return nil
}

func main() {
	sh := shell.NewShell("http://localhost:5001")
	c := NewChannel(sh, "my channel")

	c.Write(Message{User: "me", Text: "hello!"})
	c.Write(Message{User: "me", Text: "World!"})

	ch := make(chan Message)
	go func() { c.Read(sh, ch) }()

	for elem := range ch {
		println(elem.User, elem.Text)
	}
}
