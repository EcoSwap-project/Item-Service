package services

import (
	"context"

	eco "item_service/genproto/item_service"
	repo "item_service/storege/postgres"
)

type UserActivityService interface {
	GetUserActivity(context.Context, *eco.GetUserActivityRequest) (*eco.GetUserActivityResponse, error)
}

type userActivityServiceImpl struct {
	repo repo.UserActivityRepository
}

func NewUserActivityService(repo repo.UserActivityRepository) UserActivityService {
	return &userActivityServiceImpl{repo: repo}
}

func (s *userActivityServiceImpl) GetUserActivity(ctx context.Context, req *eco.GetUserActivityRequest) (*eco.GetUserActivityResponse, error) {
	activities, err := s.repo.GetUserActivity(ctx, req)
	if err != nil {
		return nil, err
	}
	return activities, nil
}
