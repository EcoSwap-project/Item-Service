package services

import (
	"context"
	"item_service/genproto/item_service"
	repository"item_service/storege/postgres"
)


type RatingService interface {
	AddUserRating(context.Context, *item_service.AddUserRatingRequest) (*item_service.AddUserRatingResponse, error)
	GetUserRatings(context.Context, *item_service.GetUserRatingsRequest) (*item_service.GetUserRatingsResponse, error)
}

type ratingServiceImpl struct {
	repo *repository.RatingRepo
}

func NewRatingService(repo *repository.RatingRepo) RatingService {
	return &ratingServiceImpl{repo: repo}
}

func (rs *ratingServiceImpl) AddUserRating(ctx context.Context, req *item_service.AddUserRatingRequest) (*item_service.AddUserRatingResponse, error) {
	return rs.repo.AddUserRating(ctx, req)
}

func (rs *ratingServiceImpl) GetUserRatings(ctx context.Context, req *item_service.GetUserRatingsRequest) (*item_service.GetUserRatingsResponse, error) {
	return rs.repo.GetUserRatings(ctx, req)
}
