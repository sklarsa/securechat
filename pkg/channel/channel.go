package channel

import (
	"bytes"
	"encoding/json"

	shell "github.com/ipfs/go-ipfs-api"
	msg "github.com/sklarsa/securechat/pkg/message"
)

type Channel struct {
	sh      *shell.Shell
	Name    string
	TailCid string
}

func (c *Channel) Read(sh *shell.Shell, ch chan msg.Message) error {
	cid := c.TailCid
	for cid != "" {
		var m msg.Message
		var err error
		m, cid, err = msg.GetMessage(sh, cid)
		if err != nil {
			return err
		}
		ch <- m
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

func (c *Channel) Write(m msg.Message) error {
	iMsg := msg.InternalMessage{
		Message:   m,
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
