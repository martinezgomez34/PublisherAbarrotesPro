package domain

type Message struct {
	Type    MessageType `json:"type"`
	Product Product     `json:"product"`
}
