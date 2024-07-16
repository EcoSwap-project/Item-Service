package services

import (
	"context"

	eco "item_service/genproto/item_service"
	repo "item_service/storege/postgres"
)

type ItemService interface {
	AddItem(context.Context, *eco.AddItemRequest) (*eco.AddItemResponse, error)
	UpdateItem(context.Context, *eco.UpdateItemRequest) (*eco.UpdateItemResponse, error)
	DeleteItem(context.Context, *eco.DeleteItemRequest) (*eco.DeleteItemResponse, error)
	GetItems(context.Context, *eco.GetItemsRequest) (*eco.GetItemsResponse, error)
	GetItem(context.Context, *eco.GetItemRequest) (*eco.GetItemResponse, error)
	SearchItems(context.Context, *eco.SearchItemsRequest) (*eco.SearchItemsResponse, error)
}

type itemServiceImpl struct {
	repo repo.ItemRepository
}

func NewItemService(repo repo.ItemRepository) ItemService {
	return &itemServiceImpl{repo: repo}
}

func (s *itemServiceImpl) AddItem(ctx context.Context, req *eco.AddItemRequest) (*eco.AddItemResponse, error) {
	item, err := s.repo.AddItem(ctx, req)
	if err != nil {
		return nil, err
	}
	return &eco.AddItemResponse{Item: item.Item}, nil
}

func (s *itemServiceImpl) UpdateItem(ctx context.Context, req *eco.UpdateItemRequest) (*eco.UpdateItemResponse, error) {
	item, err := s.repo.UpdateItem(ctx, req)
	if err != nil {
		return nil, err
	}
	return &eco.UpdateItemResponse{Item: item.Item}, nil
}

func (s *itemServiceImpl) DeleteItem(ctx context.Context, req *eco.DeleteItemRequest) (*eco.DeleteItemResponse, error) {
	_,err := s.repo.DeleteItem(ctx, req)
	if err != nil {
		return nil, err
	}
	return &eco.DeleteItemResponse{Message: "Item deleted successfully"}, nil
}

func (s *itemServiceImpl) GetItems(ctx context.Context, req *eco.GetItemsRequest) (*eco.GetItemsResponse, error) {
	items, err := s.repo.GetItems(ctx, req)
	if err != nil {
		return nil, err
	}
	return &eco.GetItemsResponse{Items: items.Items,  Page: req.Page, Limit: req.Limit}, nil
}

func (s *itemServiceImpl) GetItem(ctx context.Context, req *eco.GetItemRequest) (*eco.GetItemResponse, error) {
	item, err := s.repo.GetItem(ctx, req)
	if err != nil {
		return nil, err
	}
	return &eco.GetItemResponse{Item: item.Item}, nil
}

func (s *itemServiceImpl) SearchItems(ctx context.Context, req *eco.SearchItemsRequest) (*eco.SearchItemsResponse, error) {
	items, err := s.repo.SearchItems(ctx, req)
	if err != nil {
		return nil, err
	}
	return &eco.SearchItemsResponse{Items: items.Items,  Page: req.Page, Limit: req.Limit}, nil
}
