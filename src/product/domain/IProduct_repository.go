package domain

type IProductPublisher interface {
    PublishMessage(message Message) error
}

type IProductRepository interface {
    SaveProduct(product Product) error
    EditProduct(product Product) error
    DeleteProduct(id string) error
    GetAll() ([]*Product, error)
    GetByID(id string) (*Product, error)
}