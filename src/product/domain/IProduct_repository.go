package domain

type IProductRepository interface {
	PublishMessage(Message Message) error
}
