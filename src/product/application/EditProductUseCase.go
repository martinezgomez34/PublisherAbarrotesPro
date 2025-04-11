package application

import (
	"publisher/src/product/domain"
)
type EditProductUseCase struct {
	Repo domain.IProductPublisher
}

func NewEditProductUseCase(repo domain.IProductPublisher) *EditProductUseCase {
	return &EditProductUseCase{Repo: repo}
}

func (uc *EditProductUseCase) UpdateProduct(message domain.Message) error {
	return uc.Repo.PublishMessage(message)
}




