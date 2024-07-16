package services

import (
	"context"
	"database/sql"
	"item_service/genproto/authentication_service"
	"item_service/genproto/item_service"

	repository "item_service/storege/postgres"
)

type MainService interface {
	ItemService() ItemService
	SwapService() SwapService
	RecyclingCenterService() RecyclingCenterService
	RatingService() RatingService
	EcoChallenge() ChallengesService
}

type mainServiceImpl struct {
	authClient             authentication_service.EcoServiceClient
	itemService            ItemService
	swapService            SwapService
	recyclingCenterService RecyclingCenterService
	ratingService          RatingService
	challenge          ChallengesService
}

func NewMainService(db *sql.DB, authClient authentication_service.EcoServiceClient) MainService {
	return &mainServiceImpl{
		authClient:             authClient,
		itemService:            NewItemService(repository.NewItemRepo(db)),
		swapService:            NewSwapService(repository.NewSwapRepo(db)),
		recyclingCenterService: NewRecyclingCenterService(repository.NewRecyclingCenterRepo(db)),
		ratingService:          NewRatingService(repository.NewRatingRepo(db)),
		challenge: NewChallengesService(repository.NewChallengeRepo(db)),
	}
}

func (ms *mainServiceImpl) ItemService() ItemService {
	return ms.itemService
}

func (ms *mainServiceImpl) SwapService() SwapService {
	return ms.swapService
}

func (ms *mainServiceImpl) RecyclingCenterService() RecyclingCenterService {
	return ms.recyclingCenterService
}

func (ms *mainServiceImpl) RatingService() RatingService {
	return ms.ratingService
}
