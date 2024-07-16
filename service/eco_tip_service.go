package services

import (
	"context"

	eco "item_service/genproto/item_service"
	repo "item_service/storege/postgres"
)

type EcoTipService interface {
	AddEcoTip(context.Context, *eco.AddEcoTipRequest) (*eco.AddEcoTipResponse, error)
	GetEcoTips(context.Context, *eco.GetEcoTipsRequest) (*eco.GetEcoTipsResponse, error)
}

type ecoTipServiceImpl struct {
	repo repo.EcoTipRepository
}

func NewEcoTipService(repo repo.EcoTipRepository) EcoTipService {
	return &ecoTipServiceImpl{repo: repo}
}

func (s *ecoTipServiceImpl) AddEcoTip(ctx context.Context, req *eco.AddEcoTipRequest) (*eco.AddEcoTipResponse, error) {
	tip, err := s.repo.AddEcoTip(ctx, req)
	if err != nil {
		return nil, err
	}
	return tip, nil
}

func (s *ecoTipServiceImpl) GetEcoTips(ctx context.Context, req *eco.GetEcoTipsRequest) (*eco.GetEcoTipsResponse, error) {
	tips, err := s.repo.GetEcoTips(ctx, req)
	if err != nil {
		return nil, err
	}
	return tips, nil
}
