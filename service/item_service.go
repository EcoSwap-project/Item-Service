package services

import (
	"context"
	"item_service/genproto/item_service"
	repository"item_service/storege/postgres"
)

type ItemService interface {
	AddItem(context.Context, *item_service.AddItemRequest) (*item_service.AddItemResponse, error)
	UpdateItem(context.Context, *item_service.UpdateItemRequest) (*item_service.UpdateItemResponse, error)
	DeleteItem(context.Context, *item_service.DeleteItemRequest) (*item_service.DeleteItemResponse, error)
	GetItems(context.Context, *item_service.GetItemsRequest) (*item_service.GetItemsResponse, error)
	GetItem(context.Context, *item_service.GetItemRequest) (*item_service.GetItemResponse, error)
	SearchItems(context.Context, *item_service.SearchItemsRequest) (*item_service.SearchItemsResponse, error)
}

type itemServiceImpl struct {
	repo *repository.ItemRepo
}

func NewItemService(repo *repository.ItemRepo) ItemService {
	return &itemServiceImpl{repo: repo}
}

func (is *itemServiceImpl) AddItem(ctx context.Context, req *item_service.AddItemRequest) (*item_service.AddItemResponse, error) {
	return is.repo.AddItem(ctx, req)
}

func (is *itemServiceImpl) UpdateItem(ctx context.Context, req *item_service.UpdateItemRequest) (*item_service.UpdateItemResponse, error) {
	return is.repo.UpdateItem(ctx, req)
}

func (is *itemServiceImpl) DeleteItem(ctx context.Context, req *item_service.DeleteItemRequest) (*item_service.DeleteItemResponse, error) {
	return is.repo.DeleteItem(ctx, req)
}

func (is *itemServiceImpl) GetItems(ctx context.Context, req *item_service.GetItemsRequest) (*item_service.GetItemsResponse, error) {
	return is.repo.GetItems(ctx, req)
}

func (is *itemServiceImpl) GetItem(ctx context.Context, req *item_service.GetItemRequest) (*item_service.GetItemResponse, error) {
	return is.repo.GetItem(ctx, req)
}

func (is *itemServiceImpl) SearchItems(ctx context.Context, req *item_service.SearchItemsRequest) (*item_service.SearchItemsResponse, error) {
	return is.repo.SearchItems(ctx, req)
}
