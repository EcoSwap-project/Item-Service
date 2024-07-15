package services

import (
	"context"
	"item_service/genproto/item_service"
	repository"item_service/storege/postgres"
)

type RecyclingCenterService interface {
	AddRecyclingCenter(context.Context, *item_service.AddRecyclingCenterRequest) (*item_service.AddRecyclingCenterResponse, error)
	GetRecyclingCenters(context.Context, *item_service.GetRecyclingCentersRequest) (*item_service.GetRecyclingCentersResponse, error)
}

type recyclingCenterServiceImpl struct {
	repo *repository.RecyclingCenterRepo
}

func NewRecyclingCenterService(repo *repository.RecyclingCenterRepo) RecyclingCenterService {
	return &recyclingCenterServiceImpl{repo: repo}
}

func (rcs *recyclingCenterServiceImpl) AddRecyclingCenter(ctx context.Context, req *item_service.AddRecyclingCenterRequest) (*item_service.AddRecyclingCenterResponse, error) {
	return rcs.repo.AddRecyclingCenter(ctx, req)
}

func (rcs *recyclingCenterServiceImpl) GetRecyclingCenters(ctx context.Context, req *item_service.GetRecyclingCentersRequest) (*item_service.GetRecyclingCentersResponse, error) {
	return rcs.repo.GetRecyclingCenters(ctx, req)
}
