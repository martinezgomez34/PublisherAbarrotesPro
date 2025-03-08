package domain

type IUserRepository interface {
	PublishUser(Message Message) error
	PublishLogin(Message Message) error
}
