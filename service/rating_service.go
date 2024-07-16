package services

import (
	"context"

	eco "item_service/genproto/item_service"
	repo "item_service/storege/postgres"
)

type UserRatingService interface {
	RateUser(context.Context, *eco.AddUserRatingRequest) (*eco.AddUserRatingResponse, error)
	GetUserRatings(context.Context, *eco.GetUserRatingsRequest) (*eco.GetUserRatingsResponse, error)
}

type userRatingServiceImpl struct {
	repo repo.UserRatingRepository
}

func NewUserRatingService(repo repo.UserRatingRepository) UserRatingService {
	return &userRatingServiceImpl{repo: repo}
}

func (s *userRatingServiceImpl) RateUser(ctx context.Context, req *eco.AddUserRatingRequest) (*eco.AddUserRatingResponse, error) {
	rating, err := s.repo.AddUserRating(ctx, req)
	if err != nil {
		return nil, err
	}
	return &eco.AddUserRatingResponse{Rating: rating.Rating}, nil
}

func (s *userRatingServiceImpl) GetUserRatings(ctx context.Context, req *eco.GetUserRatingsRequest) (*eco.GetUserRatingsResponse, error) {
	ratings,  err := s.repo.GetUserRatings(ctx, req)
	if err != nil {
		return nil, err
	}
	return &eco.GetUserRatingsResponse{Ratings: ratings.Ratings, Page: req.Page, Limit: req.Limit}, nil
}
