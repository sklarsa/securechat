package message

import (
	"time"

	"github.com/sklarsa/securechat/pkg/user"
)

type Message struct {
	User                  user.User
	Text                  string `json:"text"`
	MultimediaCid         string `json:"multimediaCid"`
	MultimediaContentType string `json:"multimediaContentType"`
	Timestamp             time.Time
}

type InternalMessage struct {
	Message   `json:"message"`
	ParentCid string `json:"parentCid"`
}
