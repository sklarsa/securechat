package message

type Message struct {
	User                  string `json:"user"`
	Text                  string `json:"text"`
	MultimediaCid         string `json:"multimediaCid"`
	MultimediaContentType string `json:"multimediaContentType"`
}

type InternalMessage struct {
	Message   `json:"message"`
	ParentCid string `json:"parentCid"`
}
