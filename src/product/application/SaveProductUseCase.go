package application

import (
	"publisher/src/product/domain"
)

type ProductUseCase struct {
	Repo domain.IProductRepository
}

func NewProductUseCase(repo domain.IProductRepository) *ProductUseCase {
	return &ProductUseCase{Repo: repo}
}

func (uc *ProductUseCase) CreateProduct(message domain.Message) error {
	return uc.Repo.PublishMessage(message)
}
