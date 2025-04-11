package application

import (
	"publisher/src/product/domain"
)

type DeleteProductUseCase struct {
	Repo domain.IProductPublisher
}

func NewDeleteProductUseCase(repo domain.IProductPublisher) *DeleteProductUseCase {
	return &DeleteProductUseCase{Repo: repo}
}

func (uc *DeleteProductUseCase) DeleteProduct(message domain.Message) error {
	return uc.Repo.PublishMessage(message)
}
