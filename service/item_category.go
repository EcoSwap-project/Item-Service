package services

import (
	"context"

	eco "item_service/genproto/item_service"
	repo "item_service/storege/postgres"
)

type ItemCategoryService interface {
	AddItemCategory(context.Context, *eco.AddItemCategoryRequest) (*eco.AddItemCategoryResponse, error)
}

type itemCategoryServiceImpl struct {
	repo repo.ItemCategoryRepository
}

func NewItemCategoryService(repo repo.ItemCategoryRepository) ItemCategoryService {
	return &itemCategoryServiceImpl{repo: repo}
}

func (s *itemCategoryServiceImpl) AddItemCategory(ctx context.Context, req *eco.AddItemCategoryRequest) (*eco.AddItemCategoryResponse, error) {
	category, err := s.repo.AddItemCategory(ctx, req)
	if err != nil {
		return nil, err
	}
	return category, nil
}

