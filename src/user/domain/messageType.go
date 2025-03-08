package domain

type MessageType string

const (
	MessageTypeRegister MessageType = "register"
	MessageTypeLogin    MessageType = "login"
)
