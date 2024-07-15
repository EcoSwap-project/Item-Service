package services

import (
	"context"
	"item_service/genproto/item_service"
	repository"item_service/storege/postgres"
)

type SwapService interface {
	CreateSwap(context.Context, *item_service.CreateSwapRequest) (*item_service.CreateSwapResponse, error)
	UpdateSwapStatus(context.Context, *item_service.UpdateSwapStatusRequest) (*item_service.UpdateSwapStatusResponse, error)
	GetSwaps(context.Context, *item_service.GetSwapsRequest) (*item_service.GetSwapsResponse, error)
}

type swapServiceImpl struct {
	repo *repository.SwapRepo
}

func NewSwapService(repo *repository.SwapRepo) SwapService {
	return &swapServiceImpl{repo: repo}
}

func (ss *swapServiceImpl) CreateSwap(ctx context.Context, req *item_service.CreateSwapRequest) (*item_service.CreateSwapResponse, error) {
	return ss.repo.CreateSwap(ctx, req)
}

func (ss *swapServiceImpl) UpdateSwapStatus(ctx context.Context, req *item_service.UpdateSwapStatusRequest) (*item_service.UpdateSwapStatusResponse, error) {
	return ss.repo.UpdateSwapStatus(ctx, req)
}

func (ss *swapServiceImpl) GetSwaps(ctx context.Context, req *item_service.GetSwapsRequest) (*item_service.GetSwapsResponse, error) {
	return ss.repo.GetSwaps(ctx, req)
}
