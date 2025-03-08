package domain

type MessageType string

const (
	MessageTypeCreateProduct MessageType = "create_product"
	MessageTypeUpdateProduct MessageType = "update_product"
	MessageTypeDeleteProduct MessageType = "delete_product"
)
