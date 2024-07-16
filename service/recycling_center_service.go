package services

import (
	"context"

	eco "item_service/genproto/item_service"
	repo "item_service/storege/postgres"
)


type RecyclingCenterService interface {
	AddRecyclingCenter(context.Context, *eco.AddRecyclingCenterRequest) (*eco.AddRecyclingCenterResponse, error)
	GetRecyclingCenters(context.Context, *eco.GetRecyclingCentersRequest) (*eco.GetRecyclingCentersResponse, error)
	AddRecyclingSubmission(context.Context, *eco.AddRecyclingSubmissionRequest) (*eco.AddRecyclingSubmissionResponse, error)
}

type recyclingCenterServiceImpl struct {
	repo repo.RecyclingCenterRepository
}

func NewRecyclingCenterService(repo repo.RecyclingCenterRepository) RecyclingCenterService {
	return &recyclingCenterServiceImpl{repo: repo}
}

func (s *recyclingCenterServiceImpl) AddRecyclingCenter(ctx context.Context, req *eco.AddRecyclingCenterRequest) (*eco.AddRecyclingCenterResponse, error) {
	center, err := s.repo.AddRecyclingCenter(ctx, req)
	if err != nil {
		return nil, err
	}
	return center, nil
}

func (s *recyclingCenterServiceImpl) GetRecyclingCenters(ctx context.Context, req *eco.GetRecyclingCentersRequest) (*eco.GetRecyclingCentersResponse, error) {
	centers, err := s.repo.GetRecyclingCenters(ctx, req)
	if err != nil {
		return nil, err
	}
	return centers, nil
	// return &eco.GetRecyclingCentersResponse{RecyclingCenters: centers.Centers, Page: req.Page, Limit: req.Limit}, nil

}

func (s *recyclingCenterServiceImpl) AddRecyclingSubmission(ctx context.Context, req *eco.AddRecyclingSubmissionRequest) (*eco.AddRecyclingSubmissionResponse, error) {
	submission, err := s.repo.AddRecyclingSubmission(ctx, req)
	if err != nil {
		return nil, err
	}
	return submission, nil
}
