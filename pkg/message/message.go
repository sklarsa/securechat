package message

import (
	"bytes"
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
