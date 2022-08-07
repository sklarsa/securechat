package message

type Message struct {
	User string `json:"user"`
	Text string `json:"text"`
}

type InternalMessage struct {
	Message   `json:"message"`
	ParentCid string `json:"parentCid"`
}
