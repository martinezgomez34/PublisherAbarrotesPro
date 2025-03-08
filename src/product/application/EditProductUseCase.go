package application

import (
	"publisher/src/product/domain"
)
type EditProductUseCase struct {
	Repo domain.IProductRepository
}

func NewEditProductUseCase(repo domain.IProductRepository) *EditProductUseCase {
	return &EditProductUseCase{Repo: repo}
}

func (uc *EditProductUseCase) UpdateProduct(message domain.Message) error {
	return uc.Repo.PublishMessage(message)
}




