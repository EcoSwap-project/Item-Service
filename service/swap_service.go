package services

import (
	"context"

	eco "item_service/genproto/item_service"
	repo "item_service/storege/postgres"
)

type SwapService interface {
	CreateSwap(context.Context, *eco.CreateSwapRequest) (*eco.CreateSwapResponse, error)
	UpdateSwapStatus(context.Context, *eco.UpdateSwapStatusRequest) (*eco.UpdateSwapStatusResponse, error)
	GetSwaps(context.Context, *eco.GetSwapsRequest) (*eco.GetSwapsResponse, error)
}

type swapServiceImpl struct {
	repo repo.SwapRepository
}

func NewSwapService(repo repo.SwapRepository) SwapService {
	return &swapServiceImpl{repo: repo}
}

func (s *swapServiceImpl) CreateSwap(ctx context.Context, req *eco.CreateSwapRequest) (*eco.CreateSwapResponse, error) {
	swap, err := s.repo.CreateSwap(ctx, req)
	if err != nil {
		return nil, err
	}
	return &eco.CreateSwapResponse{Swap: swap.Swap}, nil
}

func (s *swapServiceImpl) UpdateSwapStatus(ctx context.Context, req *eco.UpdateSwapStatusRequest) (*eco.UpdateSwapStatusResponse, error) {
	swap, err := s.repo.UpdateSwapStatus(ctx, req)
	if err != nil {
		return nil, err
	}
	return &eco.UpdateSwapStatusResponse{Swap: swap.Swap}, nil
}

func (s *swapServiceImpl) GetSwaps(ctx context.Context, req *eco.GetSwapsRequest) (*eco.GetSwapsResponse, error) {
	swaps, err := s.repo.GetSwaps(ctx, req)
	if err != nil {
		return nil, err
	}
	return &eco.GetSwapsResponse{Swaps: swaps.Swaps, Page: req.Page, Limit: req.Limit}, nil
}
