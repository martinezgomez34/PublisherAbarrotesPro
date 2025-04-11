package application

import "publisher/src/product/domain"

type GetByIdProductUseCase struct {
	repo domain.IProductRepository
}

func NewGetByIDProductUseCase(repo domain.IProductRepository) *GetByIdProductUseCase {
	return &GetByIdProductUseCase{repo: repo}
}

func (uc *GetByIdProductUseCase) GetByID(id string) (*domain.Product, error) {
	return uc.repo.GetByID(id)
}
