package domain

type Message struct {
	Type  MessageType      `json:"type"`
	User  User             `json:"user,omitempty"`
	Login LoginCredentials `json:"login,omitempty"`
}
