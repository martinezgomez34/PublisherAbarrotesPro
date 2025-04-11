package application

import "publisher/src/product/domain"

type ListProductUseCase struct {
	repo domain.IProductRepository
}

func NewListProductUseCase(repo domain.IProductRepository) *ListProductUseCase {
	return &ListProductUseCase{repo: repo}
}

func (uc *ListProductUseCase) GetAll() ([]*domain.Product, error) {
	return uc.repo.GetAll()
}
