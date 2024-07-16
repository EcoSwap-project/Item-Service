package services

import (
	"context"

	eco "item_service/genproto/item_service"
	repo "item_service/storege/postgres"
)

type StatisticsService interface {
	GetStatistics(context.Context, *eco.GetStatisticsRequest) (*eco.GetStatisticsResponse, error)
}

type statisticsServiceImpl struct {
	repo repo.StatisticsRepository
}

func NewStatisticsService(repo repo.StatisticsRepository) StatisticsService {
	return &statisticsServiceImpl{repo: repo}
}

func (s *statisticsServiceImpl) GetStatistics(ctx context.Context, req *eco.GetStatisticsRequest) (*eco.GetStatisticsResponse, error) {
	statistics, err := s.repo.GetStatistics(ctx, req)
	if err != nil {
		return nil, err
	}
	return statistics, nil
}
